package apimanagement

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// GroupUsersClient is the apiManagement Client
type GroupUsersClient struct {
	BaseClient
}

// NewGroupUsersClient creates an instance of the GroupUsersClient client.
func NewGroupUsersClient(subscriptionID string) GroupUsersClient {
	return NewGroupUsersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupUsersClientWithBaseURI creates an instance of the GroupUsersClient client.
func NewGroupUsersClientWithBaseURI(baseURI string, subscriptionID string) GroupUsersClient {
	return GroupUsersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Add adds a user to the specified group.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// UID - user identifier. Must be unique in the current API Management service instance.
func (client GroupUsersClient) Add(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (result ErrorBodyContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUsersClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "groupID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUsersClient", "Add", err.Error())
	}

	req, err := client.AddPreparer(ctx, resourceGroupName, serviceName, groupID, UID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client GroupUsersClient) AddPreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users/{uid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUsersClient) AddSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client GroupUsersClient) AddResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByGroup lists a collection of the members of the group, specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// filter - | Field            | Supported operators    | Supported functions               |
// |------------------|------------------------|-----------------------------------|
// | id               | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | firstName        | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | lastName         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | email            | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | state            | eq                     | N/A                               |
// | registrationDate | ge, le, eq, ne, gt, lt | N/A                               |
// | note             | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client GroupUsersClient) ListByGroup(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result UserCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUsersClient.ListByGroup")
		defer func() {
			sc := -1
			if result.uc.Response.Response != nil {
				sc = result.uc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "groupID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUsersClient", "ListByGroup", err.Error())
	}

	result.fn = client.listByGroupNextResults
	req, err := client.ListByGroupPreparer(ctx, resourceGroupName, serviceName, groupID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "ListByGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByGroupSender(req)
	if err != nil {
		result.uc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "ListByGroup", resp, "Failure sending request")
		return
	}

	result.uc, err = client.ListByGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "ListByGroup", resp, "Failure responding to request")
	}

	return
}

// ListByGroupPreparer prepares the ListByGroup request.
func (client GroupUsersClient) ListByGroupPreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByGroupSender sends the ListByGroup request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUsersClient) ListByGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByGroupResponder handles the response to the ListByGroup request. The method always
// closes the http.Response Body.
func (client GroupUsersClient) ListByGroupResponder(resp *http.Response) (result UserCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByGroupNextResults retrieves the next set of results, if any.
func (client GroupUsersClient) listByGroupNextResults(ctx context.Context, lastResults UserCollection) (result UserCollection, err error) {
	req, err := lastResults.userCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "listByGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "listByGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "listByGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client GroupUsersClient) ListByGroupComplete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result UserCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUsersClient.ListByGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByGroup(ctx, resourceGroupName, serviceName, groupID, filter, top, skip)
	return
}

// Remove remove existing user from existing group.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// UID - user identifier. Must be unique in the current API Management service instance.
func (client GroupUsersClient) Remove(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (result ErrorBodyContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUsersClient.Remove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "groupID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUsersClient", "Remove", err.Error())
	}

	req, err := client.RemovePreparer(ctx, resourceGroupName, serviceName, groupID, UID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUsersClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client GroupUsersClient) RemovePreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users/{uid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUsersClient) RemoveSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client GroupUsersClient) RemoveResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
