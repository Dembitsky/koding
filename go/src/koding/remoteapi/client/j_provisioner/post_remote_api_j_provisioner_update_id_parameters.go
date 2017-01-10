package j_provisioner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJProvisionerUpdateIDParams creates a new PostRemoteAPIJProvisionerUpdateIDParams object
// with the default values initialized.
func NewPostRemoteAPIJProvisionerUpdateIDParams() *PostRemoteAPIJProvisionerUpdateIDParams {
	var ()
	return &PostRemoteAPIJProvisionerUpdateIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProvisionerUpdateIDParamsWithTimeout creates a new PostRemoteAPIJProvisionerUpdateIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProvisionerUpdateIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerUpdateIDParams {
	var ()
	return &PostRemoteAPIJProvisionerUpdateIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProvisionerUpdateIDParamsWithContext creates a new PostRemoteAPIJProvisionerUpdateIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProvisionerUpdateIDParamsWithContext(ctx context.Context) *PostRemoteAPIJProvisionerUpdateIDParams {
	var ()
	return &PostRemoteAPIJProvisionerUpdateIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProvisionerUpdateIDParams contains all the parameters to send to the API endpoint
for the post remote API j provisioner update ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJProvisionerUpdateIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerUpdateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) WithContext(ctx context.Context) *PostRemoteAPIJProvisionerUpdateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJProvisionerUpdateIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) WithID(id string) *PostRemoteAPIJProvisionerUpdateIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j provisioner update ID params
func (o *PostRemoteAPIJProvisionerUpdateIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProvisionerUpdateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
