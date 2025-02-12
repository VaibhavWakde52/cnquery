// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"go.mondoo.com/cnquery/v9/llx"
	"go.mondoo.com/cnquery/v9/providers-sdk/v1/util/convert"
	"go.mondoo.com/cnquery/v9/providers/aws/connection"
)

func (a *mqlAwsAccount) id() (string, error) {
	conn := a.MqlRuntime.Connection.(*connection.AwsConnection)
	return "aws.account/" + conn.AccountId(), nil
}

func (a *mqlAwsAccount) aliases() ([]interface{}, error) {
	conn := a.MqlRuntime.Connection.(*connection.AwsConnection)
	client := conn.Iam("") // no region for iam, use configured region

	res, err := client.ListAccountAliases(context.TODO(), &iam.ListAccountAliasesInput{})
	if err != nil {
		return nil, err
	}
	result := []interface{}{}
	for i := range res.AccountAliases {
		result = append(result, res.AccountAliases[i])
	}
	return result, nil
}

func (a *mqlAwsAccount) organization() (*mqlAwsOrganization, error) {
	conn := a.MqlRuntime.Connection.(*connection.AwsConnection)
	client := conn.Organizations("") // no region for orgs, use configured region

	org, err := client.DescribeOrganization(context.TODO(), &organizations.DescribeOrganizationInput{})
	if err != nil {
		return nil, err
	}
	res, err := CreateResource(a.MqlRuntime, "aws.organization",
		map[string]*llx.RawData{
			"arn":                llx.StringData(convert.ToString(org.Organization.Arn)),
			"featureSet":         llx.StringData(string(org.Organization.FeatureSet)),
			"masterAccountId":    llx.StringData(convert.ToString(org.Organization.MasterAccountId)),
			"masterAccountEmail": llx.StringData(convert.ToString(org.Organization.MasterAccountEmail)),
		})
	return res.(*mqlAwsOrganization), err
}
