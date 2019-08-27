// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build go1.8,linux

package plugin

import (
	"plugin"

	"github.com/statisticsnorway/drone-runtime/engine"
)

// Symbol the symbol name used to lookup the plugin provider value.
const Symbol = "Engine"

// Open returns a Factory dynamically loaded from a plugin.
func Open(path string) (engine.Engine, error) {
	lib, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	provider, err := lib.Lookup(Symbol)
	if err != nil {
		return nil, err
	}
	return provider.(func() (engine.Engine, error))()
}
