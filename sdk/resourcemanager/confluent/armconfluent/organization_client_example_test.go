//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_ListBySubscription.json
func ExampleOrganizationClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OrganizationResourceListResult = armconfluent.OrganizationResourceListResult{
		// 	Value: []*armconfluent.OrganizationResource{
		// 		{
		// 			Name: to.Ptr("myOrganizations"),
		// 			Type: to.Ptr("Microsoft.Confluent/organizations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganizations"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armconfluent.OrganizationResourceProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
		// 				OfferDetail: &armconfluent.OfferDetail{
		// 					ID: to.Ptr("string"),
		// 					PlanID: to.Ptr("string"),
		// 					PlanName: to.Ptr("string"),
		// 					PublisherID: to.Ptr("string"),
		// 					Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
		// 					TermUnit: to.Ptr("string"),
		// 				},
		// 				OrganizationID: to.Ptr("string"),
		// 				ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
		// 				SsoURL: to.Ptr("string"),
		// 				UserDetail: &armconfluent.UserDetail{
		// 					EmailAddress: to.Ptr("contoso@microsoft.com"),
		// 					FirstName: to.Ptr("string"),
		// 					LastName: to.Ptr("string"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_ListByResourceGroup.json
func ExampleOrganizationClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationClient().NewListByResourceGroupPager("myResourceGroup", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OrganizationResourceListResult = armconfluent.OrganizationResourceListResult{
		// 	Value: []*armconfluent.OrganizationResource{
		// 		{
		// 			Name: to.Ptr("myOrganizations"),
		// 			Type: to.Ptr("Microsoft.Confluent/organizations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganizations"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armconfluent.OrganizationResourceProperties{
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
		// 				OfferDetail: &armconfluent.OfferDetail{
		// 					ID: to.Ptr("string"),
		// 					PlanID: to.Ptr("string"),
		// 					PlanName: to.Ptr("string"),
		// 					PublisherID: to.Ptr("string"),
		// 					Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
		// 					TermUnit: to.Ptr("string"),
		// 				},
		// 				OrganizationID: to.Ptr("string"),
		// 				ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
		// 				SsoURL: to.Ptr("string"),
		// 				UserDetail: &armconfluent.UserDetail{
		// 					EmailAddress: to.Ptr("contoso@microsoft.com"),
		// 					FirstName: to.Ptr("string"),
		// 					LastName: to.Ptr("string"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_Get.json
func ExampleOrganizationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().Get(ctx, "myResourceGroup", "myOrganization", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OrganizationResource = armconfluent.OrganizationResource{
	// 	Name: to.Ptr("myOrganization"),
	// 	Type: to.Ptr("Microsoft.Confluent/organizations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armconfluent.OrganizationResourceProperties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		OfferDetail: &armconfluent.OfferDetail{
	// 			ID: to.Ptr("string"),
	// 			PlanID: to.Ptr("string"),
	// 			PlanName: to.Ptr("string"),
	// 			PublisherID: to.Ptr("string"),
	// 			Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
	// 			TermUnit: to.Ptr("string"),
	// 		},
	// 		OrganizationID: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
	// 		SsoURL: to.Ptr("string"),
	// 		UserDetail: &armconfluent.UserDetail{
	// 			EmailAddress: to.Ptr("contoso@microsoft.com"),
	// 			FirstName: to.Ptr("string"),
	// 			LastName: to.Ptr("string"),
	// 		},
	// 	},
	// 	SystemData: &armconfluent.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_Create.json
func ExampleOrganizationClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationClient().BeginCreate(ctx, "myResourceGroup", "myOrganization", &armconfluent.OrganizationClientBeginCreateOptions{Body: &armconfluent.OrganizationResource{
		Location: to.Ptr("West US"),
		Properties: &armconfluent.OrganizationResourceProperties{
			OfferDetail: &armconfluent.OfferDetail{
				ID:          to.Ptr("string"),
				PlanID:      to.Ptr("string"),
				PlanName:    to.Ptr("string"),
				PublisherID: to.Ptr("string"),
				TermUnit:    to.Ptr("string"),
			},
			UserDetail: &armconfluent.UserDetail{
				EmailAddress: to.Ptr("contoso@microsoft.com"),
				FirstName:    to.Ptr("string"),
				LastName:     to.Ptr("string"),
			},
		},
		Tags: map[string]*string{
			"Environment": to.Ptr("Dev"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OrganizationResource = armconfluent.OrganizationResource{
	// 	Name: to.Ptr("myOrganization"),
	// 	Type: to.Ptr("Microsoft.Confluent/organizations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armconfluent.OrganizationResourceProperties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		OfferDetail: &armconfluent.OfferDetail{
	// 			ID: to.Ptr("string"),
	// 			PlanID: to.Ptr("string"),
	// 			PlanName: to.Ptr("string"),
	// 			PublisherID: to.Ptr("string"),
	// 			Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
	// 			TermUnit: to.Ptr("string"),
	// 		},
	// 		OrganizationID: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
	// 		SsoURL: to.Ptr("string"),
	// 		UserDetail: &armconfluent.UserDetail{
	// 			EmailAddress: to.Ptr("contoso@microsoft.com"),
	// 			FirstName: to.Ptr("string"),
	// 			LastName: to.Ptr("string"),
	// 		},
	// 	},
	// 	SystemData: &armconfluent.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_Update.json
func ExampleOrganizationClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationClient().Update(ctx, "myResourceGroup", "myOrganization", &armconfluent.OrganizationClientUpdateOptions{Body: &armconfluent.OrganizationResourceUpdate{
		Tags: map[string]*string{
			"client": to.Ptr("dev-client"),
			"env":    to.Ptr("dev"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OrganizationResource = armconfluent.OrganizationResource{
	// 	Name: to.Ptr("myOrganization"),
	// 	Type: to.Ptr("Microsoft.Confluent/organizations"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Confluent/organizations/myOrganization"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armconfluent.OrganizationResourceProperties{
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		OfferDetail: &armconfluent.OfferDetail{
	// 			ID: to.Ptr("string"),
	// 			PlanID: to.Ptr("string"),
	// 			PlanName: to.Ptr("string"),
	// 			PublisherID: to.Ptr("string"),
	// 			Status: to.Ptr(armconfluent.SaaSOfferStatusStarted),
	// 			TermUnit: to.Ptr("string"),
	// 		},
	// 		OrganizationID: to.Ptr("string"),
	// 		ProvisioningState: to.Ptr(armconfluent.ProvisionStateSucceeded),
	// 		SsoURL: to.Ptr("string"),
	// 		UserDetail: &armconfluent.UserDetail{
	// 			EmailAddress: to.Ptr("contoso@microsoft.com"),
	// 			FirstName: to.Ptr("string"),
	// 			LastName: to.Ptr("string"),
	// 		},
	// 	},
	// 	SystemData: &armconfluent.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-09T14:28:47.284Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armconfluent.CreatedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/confluent/resource-manager/Microsoft.Confluent/stable/2021-12-01/examples/Organization_Delete.json
func ExampleOrganizationClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationClient().BeginDelete(ctx, "myResourceGroup", "myOrganization", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}