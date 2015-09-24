package network

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// LocalNetworkGatewaysClient is the client for the LocalNetworkGateways
// methods of the Network service.
type LocalNetworkGatewaysClient struct {
	ManagementClient
}

// NewLocalNetworkGatewaysClient creates an instance of the
// LocalNetworkGatewaysClient client.
func NewLocalNetworkGatewaysClient(subscriptionID string) LocalNetworkGatewaysClient {
	return NewLocalNetworkGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocalNetworkGatewaysClientWithBaseURI creates an instance of the
// LocalNetworkGatewaysClient client.
func NewLocalNetworkGatewaysClientWithBaseURI(baseURI string, subscriptionID string) LocalNetworkGatewaysClient {
	return LocalNetworkGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put LocalNetworkGateway operation creates/updates a
// local network gateway in the specified resource group through Network
// resource provider.
//
// resourceGroupName is the name of the resource group.
// localNetworkGatewayName is the name of the local network gateway.
// parameters is parameters supplied to the Begin Create or update Local
// Network Gateway operation through Network resource provider.
func (client LocalNetworkGatewaysClient) CreateOrUpdate(resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (result LocalNetworkGateway, ae error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "CreateOrUpdate", "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LocalNetworkGatewaysClient) CreateOrUpdatePreparer(resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": url.QueryEscape(localNetworkGatewayName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusCreated, http.StatusOK)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete LocalNetworkGateway operation deletes the specifed local
// network Gateway through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// localNetworkGatewayName is the name of the local network gateway.
func (client LocalNetworkGatewaysClient) Delete(resourceGroupName string, localNetworkGatewayName string) (result autorest.Response, ae error) {
	req, err := client.DeletePreparer(resourceGroupName, localNetworkGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "Delete", "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "Delete", "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "Delete", "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LocalNetworkGatewaysClient) DeletePreparer(resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": url.QueryEscape(localNetworkGatewayName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusAccepted, http.StatusNoContent, http.StatusOK)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get LocalNetworkGateway operation retrieves information about the
// specified local network gateway through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// localNetworkGatewayName is the name of the local network gateway.
func (client LocalNetworkGatewaysClient) Get(resourceGroupName string, localNetworkGatewayName string) (result LocalNetworkGateway, ae error) {
	req, err := client.GetPreparer(resourceGroupName, localNetworkGatewayName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "Get", "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "Get", "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "Get", "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LocalNetworkGatewaysClient) GetPreparer(resourceGroupName string, localNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"localNetworkGatewayName": url.QueryEscape(localNetworkGatewayName),
		"resourceGroupName":       url.QueryEscape(resourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) GetResponder(resp *http.Response) (result LocalNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List LocalNetworkGateways opertion retrieves all the local network
// gateways stored.
//
// resourceGroupName is the name of the resource group.
func (client LocalNetworkGatewaysClient) List(resourceGroupName string) (result LocalNetworkGatewayListResult, ae error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "List", "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "List", "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "List", "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LocalNetworkGatewaysClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocalNetworkGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, http.StatusOK)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocalNetworkGatewaysClient) ListResponder(resp *http.Response) (result LocalNetworkGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		autorest.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client LocalNetworkGatewaysClient) ListNextResults(lastResults LocalNetworkGatewayListResult) (result LocalNetworkGatewayListResult, ae error) {
	req, err := lastResults.LocalNetworkGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "List", "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "List", "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		ae = autorest.NewErrorWithError(err, "network/LocalNetworkGatewaysClient", "List", "Failure responding to next results request request")
	}

	return
}
