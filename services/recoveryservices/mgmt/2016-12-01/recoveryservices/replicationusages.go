package recoveryservices

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
	"net/http"
)

// ReplicationUsagesClient is the recovery Services Client
type ReplicationUsagesClient struct {
	BaseClient
}

// NewReplicationUsagesClient creates an instance of the ReplicationUsagesClient client.
func NewReplicationUsagesClient(subscriptionID string) ReplicationUsagesClient {
	return NewReplicationUsagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReplicationUsagesClientWithBaseURI creates an instance of the ReplicationUsagesClient client.
func NewReplicationUsagesClientWithBaseURI(baseURI string, subscriptionID string) ReplicationUsagesClient {
	return ReplicationUsagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List fetches the replication usages of the vault.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is the
// name of the recovery services vault.
func (client ReplicationUsagesClient) List(ctx context.Context, resourceGroupName string, vaultName string) (result ReplicationUsageList, err error) {
	req, err := client.ListPreparer(ctx, resourceGroupName, vaultName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.ReplicationUsagesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.ReplicationUsagesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.ReplicationUsagesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationUsagesClient) ListPreparer(ctx context.Context, resourceGroupName string, vaultName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/replicationUsages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationUsagesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationUsagesClient) ListResponder(resp *http.Response) (result ReplicationUsageList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
