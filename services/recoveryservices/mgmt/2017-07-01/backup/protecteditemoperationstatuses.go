package backup

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

// ProtectedItemOperationStatusesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectedItemOperationStatusesClient struct {
	BaseClient
}

// NewProtectedItemOperationStatusesClient creates an instance of the ProtectedItemOperationStatusesClient client.
func NewProtectedItemOperationStatusesClient(subscriptionID string) ProtectedItemOperationStatusesClient {
	return NewProtectedItemOperationStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectedItemOperationStatusesClientWithBaseURI creates an instance of the ProtectedItemOperationStatusesClient
// client.
func NewProtectedItemOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) ProtectedItemOperationStatusesClient {
	return ProtectedItemOperationStatusesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get fetches the status of an operation such as triggering a backup, restore. The status can be in progress,
// completed or failed. You can refer to the OperationStatus enum for all the possible states of the operation. Some
// operations create jobs. This method returns the list of jobs associated with the operation.
//
// vaultName is the name of the recovery services vault. resourceGroupName is the name of the resource group where the
// recovery services vault is present. fabricName is fabric name associated with the backup item. containerName is
// container name associated with the backup item. protectedItemName is backup item name whose details are to be
// fetched. operationID is operationID represents the operation whose status needs to be fetched.
func (client ProtectedItemOperationStatusesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result OperationStatus, err error) {
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemOperationStatusesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemOperationStatusesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectedItemOperationStatusesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectedItemOperationStatusesClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"operationId":       autorest.Encode("path", operationID),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/operationsStatus/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemOperationStatusesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectedItemOperationStatusesClient) GetResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
