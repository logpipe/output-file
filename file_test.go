package main

import (
	"github.com/logpipe/logpipe/engine"
	"testing"
)

func TestFileOutput(t *testing.T) {

	output := &FileOutput{path: "data.log", delim: '\n'}
	engine.DebugOutput(output)
	engine.Wait()
}
