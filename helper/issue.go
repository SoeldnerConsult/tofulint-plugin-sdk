package helper

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/hashicorp/hcl/v2"
)

// Issue is a stub that has the same structure as the actually used issue object.
// This is only used for testing, as the mock Runner doesn't depend on the actual Issue structure.
type Issue struct {
	Rule    tflint.Rule
	Message string
	Range   hcl.Range
}

// Issues is a list of Issue.
type Issues []*Issue
