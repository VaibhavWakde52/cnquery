// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package resources

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/applications"
	"go.mondoo.com/cnquery/v9/llx"
	"go.mondoo.com/cnquery/v9/providers-sdk/v1/util/convert"
	"go.mondoo.com/cnquery/v9/providers/ms365/connection"
	"go.mondoo.com/cnquery/v9/types"
)

func (m *mqlMicrosoftApplication) id() (string, error) {
	return m.Id.Data, nil
}

func (a *mqlMicrosoft) applications() ([]interface{}, error) {
	conn := a.MqlRuntime.Connection.(*connection.Ms365Connection)
	graphClient, err := graphClient(conn)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	resp, err := graphClient.Applications().Get(ctx, &applications.ApplicationsRequestBuilderGetRequestConfiguration{})
	if err != nil {
		return nil, transformError(err)
	}

	res := []interface{}{}
	apps := resp.GetValue()
	for _, app := range apps {
		mqlResource, err := CreateResource(a.MqlRuntime, "microsoft.application",
			map[string]*llx.RawData{
				"id":              llx.StringData(convert.ToString(app.GetId())),
				"appId":           llx.StringData(convert.ToString(app.GetAppId())),
				"createdDateTime": llx.TimeDataPtr(app.GetCreatedDateTime()),
				"displayName":     llx.StringData(convert.ToString(app.GetDisplayName())),
				"publisherDomain": llx.StringData(convert.ToString(app.GetPublisherDomain())),
				"signInAudience":  llx.StringData(convert.ToString(app.GetSignInAudience())),
				"identifierUris":  llx.ArrayData(convert.SliceAnyToInterface(app.GetIdentifierUris()), types.String),
			})
		if err != nil {
			return nil, err
		}
		res = append(res, mqlResource)
	}

	return res, nil
}
