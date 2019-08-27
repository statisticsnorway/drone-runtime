// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runtime

import (
	"testing"

	"github.com/statisticsnorway/drone-runtime/engine"
)

func TestWithHooks(t *testing.T) {
	h := &Hook{}
	r := New(WithHooks(h))
	if r.hook != h {
		t.Errorf("Option does not set runtime hooks")
	}
}

func TestWithConfig(t *testing.T) {
	c := &engine.Spec{}
	r := New(WithConfig(c))
	if r.config != c {
		t.Errorf("Option does not set runtime configuration")
	}
}
