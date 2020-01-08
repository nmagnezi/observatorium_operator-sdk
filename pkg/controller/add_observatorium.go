package controller

import (
	"github.com/observatorium/observatorium-operator/pkg/controller/observatorium"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, observatorium.Add)
}
