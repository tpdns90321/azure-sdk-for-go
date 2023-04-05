//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple8000series

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// AccessControlRecordsClient contains the methods for the AccessControlRecords group.
// Don't use this type directly, use NewAccessControlRecordsClient() instead.
type AccessControlRecordsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessControlRecordsClient creates a new instance of AccessControlRecordsClient with the specified values.
//   - subscriptionID - The subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessControlRecordsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessControlRecordsClient, error) {
	cl, err := arm.NewClient(moduleName+".AccessControlRecordsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessControlRecordsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or Updates an access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - accessControlRecordName - The name of the access control record.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - parameters - The access control record to be added or updated.
//   - options - AccessControlRecordsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccessControlRecordsClient.BeginCreateOrUpdate
//     method.
func (client *AccessControlRecordsClient) BeginCreateOrUpdate(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, parameters AccessControlRecord, options *AccessControlRecordsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AccessControlRecordsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, accessControlRecordName, resourceGroupName, managerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AccessControlRecordsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AccessControlRecordsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or Updates an access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *AccessControlRecordsClient) createOrUpdate(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, parameters AccessControlRecord, options *AccessControlRecordsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, accessControlRecordName, resourceGroupName, managerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccessControlRecordsClient) createOrUpdateCreateRequest(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, parameters AccessControlRecord, options *AccessControlRecordsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}"
	urlPath = strings.ReplaceAll(urlPath, "{accessControlRecordName}", accessControlRecordName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - accessControlRecordName - The name of the access control record to delete.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - AccessControlRecordsClientBeginDeleteOptions contains the optional parameters for the AccessControlRecordsClient.BeginDelete
//     method.
func (client *AccessControlRecordsClient) BeginDelete(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientBeginDeleteOptions) (*runtime.Poller[AccessControlRecordsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, accessControlRecordName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AccessControlRecordsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AccessControlRecordsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the access control record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
func (client *AccessControlRecordsClient) deleteOperation(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, accessControlRecordName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessControlRecordsClient) deleteCreateRequest(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}"
	urlPath = strings.ReplaceAll(urlPath, "{accessControlRecordName}", accessControlRecordName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Returns the properties of the specified access control record name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-06-01
//   - accessControlRecordName - Name of access control record to be fetched.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - AccessControlRecordsClientGetOptions contains the optional parameters for the AccessControlRecordsClient.Get
//     method.
func (client *AccessControlRecordsClient) Get(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientGetOptions) (AccessControlRecordsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, accessControlRecordName, resourceGroupName, managerName, options)
	if err != nil {
		return AccessControlRecordsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessControlRecordsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessControlRecordsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccessControlRecordsClient) getCreateRequest(ctx context.Context, accessControlRecordName string, resourceGroupName string, managerName string, options *AccessControlRecordsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}"
	urlPath = strings.ReplaceAll(urlPath, "{accessControlRecordName}", accessControlRecordName)
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessControlRecordsClient) getHandleResponse(resp *http.Response) (AccessControlRecordsClientGetResponse, error) {
	result := AccessControlRecordsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessControlRecord); err != nil {
		return AccessControlRecordsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagerPager - Retrieves all the access control records in a manager.
//
// Generated from API version 2017-06-01
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - AccessControlRecordsClientListByManagerOptions contains the optional parameters for the AccessControlRecordsClient.NewListByManagerPager
//     method.
func (client *AccessControlRecordsClient) NewListByManagerPager(resourceGroupName string, managerName string, options *AccessControlRecordsClientListByManagerOptions) *runtime.Pager[AccessControlRecordsClientListByManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessControlRecordsClientListByManagerResponse]{
		More: func(page AccessControlRecordsClientListByManagerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AccessControlRecordsClientListByManagerResponse) (AccessControlRecordsClientListByManagerResponse, error) {
			req, err := client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
			if err != nil {
				return AccessControlRecordsClientListByManagerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AccessControlRecordsClientListByManagerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessControlRecordsClientListByManagerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByManagerHandleResponse(resp)
		},
	})
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *AccessControlRecordsClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *AccessControlRecordsClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", client.subscriptionID)
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", resourceGroupName)
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", managerName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagerHandleResponse handles the ListByManager response.
func (client *AccessControlRecordsClient) listByManagerHandleResponse(resp *http.Response) (AccessControlRecordsClientListByManagerResponse, error) {
	result := AccessControlRecordsClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessControlRecordList); err != nil {
		return AccessControlRecordsClientListByManagerResponse{}, err
	}
	return result, nil
}