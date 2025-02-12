// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"go.mondoo.com/cnquery/v9/llx"
	"go.mondoo.com/cnquery/v9/providers-sdk/v1/util/convert"
	"go.mondoo.com/cnquery/v9/providers/azure/connection"
	"go.mondoo.com/cnquery/v9/types"

	azureres "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func (a *mqlAzureSubscription) resourceGroups() ([]interface{}, error) {
	conn := a.MqlRuntime.Connection.(*connection.AzureConnection)

	ctx := context.Background()
	token := conn.Token()
	subId := a.SubscriptionId.Data

	client, err := azureres.NewResourceGroupsClient(subId, token, &arm.ClientOptions{})
	if err != nil {
		return nil, err
	}

	pager := client.NewListPager(&azureres.ResourceGroupsClientListOptions{})
	res := []interface{}{}
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, rg := range page.Value {
			mqlAzure, err := CreateResource(a.MqlRuntime, "azure.subscription.resourcegroup",
				map[string]*llx.RawData{
					"id":                llx.StringData(convert.ToString(rg.ID)),
					"name":              llx.StringData(convert.ToString(rg.Name)),
					"location":          llx.StringData(convert.ToString(rg.Location)),
					"tags":              llx.MapData(convert.PtrMapStrToInterface(rg.Tags), types.String),
					"type":              llx.StringData(convert.ToString(rg.Type)),
					"provisioningState": llx.StringData(convert.ToString(rg.Properties.ProvisioningState)),
					"managedBy":         llx.StringData(convert.ToString(rg.ManagedBy)),
				},
			)
			if err != nil {
				return nil, err
			}
			res = append(res, mqlAzure)
		}
	}

	return res, nil
}

func (a *mqlAzureSubscriptionResourcegroup) id() (string, error) {
	return a.Id.Data, nil
}
