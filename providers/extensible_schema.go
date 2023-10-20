// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package providers

import (
	"sync"

	"github.com/rs/zerolog/log"
	"go.mondoo.com/cnquery/v9/providers-sdk/v1/resources"
)

type extensibleSchema struct {
	resources.Schema

	loaded        map[string]struct{}
	runtime       *Runtime
	lastRefreshed int64
	lockAll       sync.Mutex // only used in getting all schemas
	lockAdd       sync.Mutex // only used when adding a schema
}

func (x *extensibleSchema) loadAllSchemas() {
	x.lockAll.Lock()
	defer x.lockAll.Unlock()

	// If another goroutine started to load this before us, it will be locked until
	// we complete to load everything and then it will be dumped into this
	// position. At this point, if it has been loaded we can return safely, since
	// we don't unlock until we are finished loading.
	if x.lastRefreshed >= LastProviderInstall {
		return
	}
	x.lastRefreshed = LastProviderInstall

	providers, err := ListActive()
	if err != nil {
		log.Error().Err(err).Msg("failed to list all providers, can't load additional schemas")
		return
	}

	for name := range providers {
		if name == BuiltinCoreID {
			continue
		}
		schema, err := x.runtime.coordinator.LoadSchema(name)
		if err != nil {
			log.Error().Err(err).Msg("load schema failed")
		} else {
			x.Add(name, schema)
		}
	}

	// We are loading the core provider last, so it overrides all other schemas.
	// It will ensure that core fields are preferred.
	if _, ok := providers[BuiltinCoreID]; ok {
		schema, err := x.runtime.coordinator.LoadSchema(BuiltinCoreID)
		if err != nil {
			log.Error().Err(err).Msg("load schema failed")
		} else {
			x.Add(BuiltinCoreID, schema)
		}
	}
}

func (x *extensibleSchema) Close() {
	x.loaded = map[string]struct{}{}
	x.Schema.Resources = nil
}

func (x *extensibleSchema) Lookup(name string) *resources.ResourceInfo {
	if found, ok := x.Resources[name]; ok {
		return found
	}
	if x.lastRefreshed >= LastProviderInstall {
		return nil
	}

	x.loadAllSchemas()
	return x.Resources[name]
}

func (x *extensibleSchema) LookupField(resource string, field string) (*resources.ResourceInfo, *resources.Field) {
	found, ok := x.Resources[resource]
	if !ok {
		if x.lastRefreshed >= LastProviderInstall {
			return nil, nil
		}

		x.loadAllSchemas()

		found, ok = x.Resources[resource]
		if !ok {
			return nil, nil
		}
		return found, found.Fields[field]
	}

	fieldObj, ok := found.Fields[field]
	if ok {
		return found, fieldObj
	}
	if x.lastRefreshed >= LastProviderInstall {
		return found, nil
	}

	x.loadAllSchemas()
	return found, found.Fields[field]
}

func (x *extensibleSchema) Add(name string, schema *resources.Schema) {
	if schema == nil {
		return
	}
	if name == "" {
		log.Error().Msg("tried to add a schema with no name")
		return
	}

	x.lockAdd.Lock()
	defer x.lockAdd.Unlock()

	if _, ok := x.loaded[name]; ok {
		return
	}

	x.loaded[name] = struct{}{}
	x.Schema.Add(schema)
}

func (x *extensibleSchema) AllResources() map[string]*resources.ResourceInfo {
	if x.lastRefreshed < LastProviderInstall {
		x.loadAllSchemas()
	}

	return x.Resources
}
