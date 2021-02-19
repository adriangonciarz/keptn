// Code generated by go-swagger; DO NOT EDIT.

package service_approval

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateServiceApprovalHandlerFunc turns a function with the right signature into a create service approval handler
type CreateServiceApprovalHandlerFunc func(CreateServiceApprovalParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateServiceApprovalHandlerFunc) Handle(params CreateServiceApprovalParams) middleware.Responder {
	return fn(params)
}

// CreateServiceApprovalHandler interface for that can handle valid create service approval params
type CreateServiceApprovalHandler interface {
	Handle(CreateServiceApprovalParams) middleware.Responder
}

// NewCreateServiceApproval creates a new http.Handler for the create service approval operation
func NewCreateServiceApproval(ctx *middleware.Context, handler CreateServiceApprovalHandler) *CreateServiceApproval {
	return &CreateServiceApproval{Context: ctx, Handler: handler}
}

/*CreateServiceApproval swagger:route POST /project/{projectName}/stage/{stageName}/service/{serviceName}/approval Service approval createServiceApproval

Create service approval

*/
type CreateServiceApproval struct {
	Context *middleware.Context
	Handler CreateServiceApprovalHandler
}

func (o *CreateServiceApproval) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateServiceApprovalParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}