package handlers

import "github.com/hbd/go-crm/controllers"

// Controller .
type Controller struct {
	*controllers.Controller
}

// NewController .
func NewController() *Controller {
	return &Controller{controllers.NewController()}
}
