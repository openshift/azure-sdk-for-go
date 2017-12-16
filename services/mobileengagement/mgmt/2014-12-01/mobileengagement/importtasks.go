package mobileengagement

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ImportTasksClient is the microsoft Azure Mobile Engagement REST APIs.
type ImportTasksClient struct {
	BaseClient
}

// NewImportTasksClient creates an instance of the ImportTasksClient client.
func NewImportTasksClient(subscriptionID string) ImportTasksClient {
	return NewImportTasksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewImportTasksClientWithBaseURI creates an instance of the ImportTasksClient client.
func NewImportTasksClientWithBaseURI(baseURI string, subscriptionID string) ImportTasksClient {
	return ImportTasksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a job to import the specified data to a storageUrl.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name.
func (client ImportTasksClient) Create(ctx context.Context, resourceGroupName string, appCollection string, appName string, parameters ImportTask) (result ImportTaskResult, err error) {
	req, err := client.CreatePreparer(ctx, resourceGroupName, appCollection, appName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ImportTasksClient) CreatePreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, parameters ImportTask) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/importTasks", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ImportTasksClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ImportTasksClient) CreateResponder(resp *http.Response) (result ImportTaskResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get the Get import job operation retrieves information about a previously created import job.
//
// ID is import job identifier. resourceGroupName is the name of the resource group. appCollection is application
// collection. appName is application resource name.
func (client ImportTasksClient) Get(ctx context.Context, ID string, resourceGroupName string, appCollection string, appName string) (result ImportTaskResult, err error) {
	req, err := client.GetPreparer(ctx, ID, resourceGroupName, appCollection, appName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ImportTasksClient) GetPreparer(ctx context.Context, ID string, resourceGroupName string, appCollection string, appName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"id":                autorest.Encode("path", ID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/importTasks/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ImportTasksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ImportTasksClient) GetResponder(resp *http.Response) (result ImportTaskResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get the list of import jobs.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name. skip is control paging of import jobs, start results at the given offset, defaults to 0 (1st page of
// data). top is control paging of import jobs, number of import jobs to return with each call. By default, it returns
// all import jobs with a default paging of 20.
// The response contains a `nextLink` property describing the path to get the next page if there are more results.
// The maximum paging limit for $top is 40. orderby is sort results by an expression which looks like `$orderby=jobId
// asc` (default when not specified).
// The syntax is orderby={property} {direction} or just orderby={property}.
// Properties that can be specified for sorting: jobId, errorDetails, dateCreated, jobStatus, and dateCreated.
// The available directions are asc (for ascending order) and desc (for descending order).
// When not specified the asc direction is used.
// Only one orderby property can be specified.
func (client ImportTasksClient) List(ctx context.Context, resourceGroupName string, appCollection string, appName string, skip *int32, top *int32, orderby string) (result ImportTaskListResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 40, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "mobileengagement.ImportTasksClient", "List")
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, appCollection, appName, skip, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.itlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", resp, "Failure sending request")
		return
	}

	result.itlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ImportTasksClient) ListPreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, skip *int32, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/importTasks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ImportTasksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ImportTasksClient) ListResponder(resp *http.Response) (result ImportTaskListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ImportTasksClient) listNextResults(lastResults ImportTaskListResult) (result ImportTaskListResult, err error) {
	req, err := lastResults.importTaskListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.ImportTasksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ImportTasksClient) ListComplete(ctx context.Context, resourceGroupName string, appCollection string, appName string, skip *int32, top *int32, orderby string) (result ImportTaskListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, appCollection, appName, skip, top, orderby)
	return
}
