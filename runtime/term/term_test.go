// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package term

import (
	"bytes"
	"testing"

	"github.com/statisticsnorway/drone-runtime/engine"
	"github.com/statisticsnorway/drone-runtime/runtime"
)

func TestWriteLine(t *testing.T) {
	var (
		buf   bytes.Buffer
		step  = &engine.Step{Metadata: engine.Metadata{Name: "test"}}
		line  = &runtime.Line{Number: 1, Message: "hello"}
		state = &runtime.State{Step: step}
	)

	WriteLine(&buf)(state, line)

	if got, want := buf.String(), "[test:1] hello"; got != want {
		t.Errorf("Want line %q, got %q", want, got)
	}
}

func TestWriteLinePretty(t *testing.T) {
	var (
		buf   bytes.Buffer
		step  = &engine.Step{Metadata: engine.Metadata{Name: "test"}}
		line  = &runtime.Line{Number: 1, Message: "hello"}
		state = &runtime.State{Step: step}
	)

	WriteLinePretty(&buf)(state, line)

	if got, want := buf.String(), "\x1b[32m[test:1]\x1b[0m hello"; got != want {
		t.Errorf("Want line %q, got %q", want, got)
	}
}
