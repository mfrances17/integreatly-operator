package controller

import (
	"github.com/aerogear/mobile-security-service-operator/pkg/controller/mobilesecurityserviceapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mobilesecurityserviceapp.Add)
}
