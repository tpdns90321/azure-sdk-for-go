//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ArmTemplatesClient contains the methods for the ArmTemplates group.
// Don't use this type directly, use NewArmTemplatesClient() instead.
type ArmTemplatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewArmTemplatesClient creates a new instance of ArmTemplatesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArmTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArmTemplatesClient, error) {
	cl, err := arm.NewClient(moduleName+".ArmTemplatesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ArmTemplatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get azure resource manager template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - artifactSourceName - The name of the artifact source.
//   - name - The name of the azure resource manager template.
//   - options - ArmTemplatesClientGetOptions contains the optional parameters for the ArmTemplatesClient.Get method.
func (client *ArmTemplatesClient) Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArmTemplatesClientGetOptions) (ArmTemplatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, name, options)
	if err != nil {
		return ArmTemplatesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArmTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArmTemplatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArmTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, options *ArmTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArmTemplatesClient) getHandleResponse(resp *http.Response) (ArmTemplatesClientGetResponse, error) {
	result := ArmTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmTemplate); err != nil {
		return ArmTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List azure resource manager templates in a given artifact source.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - artifactSourceName - The name of the artifact source.
//   - options - ArmTemplatesClientListOptions contains the optional parameters for the ArmTemplatesClient.NewListPager method.
func (client *ArmTemplatesClient) NewListPager(resourceGroupName string, labName string, artifactSourceName string, options *ArmTemplatesClientListOptions) *runtime.Pager[ArmTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ArmTemplatesClientListResponse]{
		More: func(page ArmTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ArmTemplatesClientListResponse) (ArmTemplatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, labName, artifactSourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ArmTemplatesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ArmTemplatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ArmTemplatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ArmTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, options *ArmTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArmTemplatesClient) listHandleResponse(resp *http.Response) (ArmTemplatesClientListResponse, error) {
	result := ArmTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArmTemplateList); err != nil {
		return ArmTemplatesClientListResponse{}, err
	}
	return result, nil
}