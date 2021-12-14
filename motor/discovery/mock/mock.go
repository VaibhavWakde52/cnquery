package mock

import (
	"go.mondoo.io/mondoo/motor/asset"
	"go.mondoo.io/mondoo/motor/discovery/common"
	"go.mondoo.io/mondoo/motor/motorid"
	"go.mondoo.io/mondoo/motor/motorid/hostname"
	"go.mondoo.io/mondoo/motor/platform"
	"go.mondoo.io/mondoo/motor/transports"
	"go.mondoo.io/mondoo/motor/transports/resolver"
)

type Resolver struct{}

func (r *Resolver) Name() string {
	return "Mock Resolver"
}

func (r *Resolver) AvailableDiscoveryTargets() []string {
	return []string{}
}

func (r *Resolver) Resolve(tc *transports.TransportConfig, cfn common.CredentialFn, sfn common.QuerySecretFn, userIdDetectors ...transports.PlatformIdDetector) ([]*asset.Asset, error) {
	assetInfo := &asset.Asset{
		State:       asset.State_STATE_ONLINE,
		Connections: []*transports.TransportConfig{tc},
	}

	// use hostname as name if asset name was not explicitly provided
	if assetInfo.Name == "" {
		assetInfo.Name = tc.Host
	}

	assetInfo.Connections = []*transports.TransportConfig{tc}

	assetInfo.Platform = &platform.Platform{
		Kind: transports.Kind_KIND_BARE_METAL,
	}

	m, err := resolver.NewMotorConnection(tc, cfn)
	if err != nil {
		return nil, err
	}
	defer m.Close()

	// determine platform information
	p, err := m.Platform()
	if err == nil {
		assetInfo.Platform = p
	}

	platformIds, assetMetadata, err := motorid.GatherIDs(m.Transport, p, userIdDetectors)
	if err != nil {
		return nil, err
	}
	assetInfo.PlatformIds = platformIds
	if assetMetadata.Name != "" {
		assetInfo.Name = assetMetadata.Name
	}

	// use hostname as asset name
	if p != nil && assetInfo.Name == "" {
		// retrieve hostname
		hostname, err := hostname.Hostname(m.Transport, p)
		if err == nil && len(hostname) > 0 {
			assetInfo.Name = hostname
		}
	}
	return []*asset.Asset{assetInfo}, nil
}
