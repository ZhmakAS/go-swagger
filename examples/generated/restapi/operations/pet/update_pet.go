package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// UpdatePetHandlerFunc turns a function with the right signature into a update pet handler
type UpdatePetHandlerFunc func(UpdatePetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePetHandlerFunc) Handle(params UpdatePetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdatePetHandler interface for that can handle valid update pet params
type UpdatePetHandler interface {
	Handle(UpdatePetParams, interface{}) middleware.Responder
}

// NewUpdatePet creates a new http.Handler for the update pet operation
func NewUpdatePet(ctx *middleware.Context, handler UpdatePetHandler) *UpdatePet {
	return &UpdatePet{Context: ctx, Handler: handler}
}

/*UpdatePet swagger:route PUT /pets pet updatePet

Update an existing pet

*/
type UpdatePet struct {
	Context *middleware.Context
	Params  UpdatePetParams
	Handler UpdatePetHandler
}

func (o *UpdatePet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewUpdatePetParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
