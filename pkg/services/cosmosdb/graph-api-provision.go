package cosmosdb

import (
	"context"
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (g *graphAccountManager) GetProvisioner(
	service.Plan,
) (service.Provisioner, error) {
	return service.NewProvisioner(
		service.NewProvisioningStep(
			"preProvision", g.preProvision),
		service.NewProvisioningStep("deployARMTemplate", g.deployARMTemplate),
	)
}

func (g *graphAccountManager) deployARMTemplate(
	ctx context.Context,
	instance service.Instance,
) (service.InstanceDetails, service.SecureInstanceDetails, error) {

	pp := &provisioningParameters{}
	if err :=
		service.GetStructFromMap(instance.ProvisioningParameters, pp); err != nil {
		return nil, nil, err
	}

	dt := &cosmosdbInstanceDetails{}
	if err := service.GetStructFromMap(instance.Details, &dt); err != nil {
		return nil, nil, err
	}

	p := g.buildGoTemplateParams(pp, dt, "GlobalDocumentDB")
	p["capability"] = "EnableGremlin"
	if instance.Tags == nil {
		instance.Tags = make(map[string]string)
	}
	instance.Tags["defaultExperience"] = "Graph"

	dt, sdt, err := g.cosmosAccountManager.deployARMTemplate(ctx, instance, p)

	if err != nil {
		return nil, nil, fmt.Errorf("error deploying ARM template: %s", err)
	}

	sdt.ConnectionString = fmt.Sprintf("AccountEndpoint=%s;AccountKey=%s;",
		dt.FullyQualifiedDomainName,
		sdt.PrimaryKey,
	)

	dtMap, err := service.GetMapFromStruct(dt)
	if err != nil {
		return nil, nil, err
	}
	sdtMap, err := service.GetMapFromStruct(sdt)
	return dtMap, sdtMap, err
}
