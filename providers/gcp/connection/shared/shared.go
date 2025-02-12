// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package shared

import (
	"go.mondoo.com/cnquery/v9/providers-sdk/v1/inventory"
)

type ConnectionType string

type GcpConnection interface {
	ID() uint32
	Name() string
	Type() ConnectionType
	Config() *inventory.Config
	Asset() *inventory.Asset
}

type Capabilities byte

const (
	Capability_Aws_Ebs Capabilities = 1 << iota
)
