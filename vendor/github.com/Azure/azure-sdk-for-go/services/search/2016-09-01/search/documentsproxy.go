package search

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/satori/go.uuid"
	"net/http"
)

// DocumentsProxyClient is the search Client
type DocumentsProxyClient struct {
	BaseClient
}

// NewDocumentsProxyClient creates an instance of the DocumentsProxyClient client.
func NewDocumentsProxyClient() DocumentsProxyClient {
	return NewDocumentsProxyClientWithBaseURI(DefaultBaseURI)
}

// NewDocumentsProxyClientWithBaseURI creates an instance of the DocumentsProxyClient client.
func NewDocumentsProxyClientWithBaseURI(baseURI string) DocumentsProxyClient {
	return DocumentsProxyClient{NewWithBaseURI(baseURI)}
}

// Count queries the number of documents in the Azure Search index.
//
// clientRequestID is the tracking ID sent with the request to help with debugging.
func (client DocumentsProxyClient) Count(ctx context.Context, clientRequestID *uuid.UUID) (result Int64, err error) {
	req, err := client.CountPreparer(ctx, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.DocumentsProxyClient", "Count", nil, "Failure preparing request")
		return
	}

	resp, err := client.CountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.DocumentsProxyClient", "Count", resp, "Failure sending request")
		return
	}

	result, err = client.CountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.DocumentsProxyClient", "Count", resp, "Failure responding to request")
	}

	return
}

// CountPreparer prepares the Count request.
func (client DocumentsProxyClient) CountPreparer(ctx context.Context, clientRequestID *uuid.UUID) (*http.Request, error) {
	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/docs/$count"),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CountSender sends the Count request. The method will close the
// http.Response Body if it receives an error.
func (client DocumentsProxyClient) CountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CountResponder handles the response to the Count request. The method always
// closes the http.Response Body.
func (client DocumentsProxyClient) CountResponder(resp *http.Response) (result Int64, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
