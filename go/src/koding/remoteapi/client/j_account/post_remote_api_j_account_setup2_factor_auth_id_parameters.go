package j_account

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

// NewPostRemoteAPIJAccountSetup2FactorAuthIDParams creates a new PostRemoteAPIJAccountSetup2FactorAuthIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountSetup2FactorAuthIDParams() *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	var ()
	return &PostRemoteAPIJAccountSetup2FactorAuthIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountSetup2FactorAuthIDParamsWithTimeout creates a new PostRemoteAPIJAccountSetup2FactorAuthIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountSetup2FactorAuthIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	var ()
	return &PostRemoteAPIJAccountSetup2FactorAuthIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountSetup2FactorAuthIDParamsWithContext creates a new PostRemoteAPIJAccountSetup2FactorAuthIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountSetup2FactorAuthIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	var ()
	return &PostRemoteAPIJAccountSetup2FactorAuthIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountSetup2FactorAuthIDParams contains all the parameters to send to the API endpoint
for the post remote API j account setup2 factor auth ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountSetup2FactorAuthIDParams struct {

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

// WithTimeout adds the timeout to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) WithID(id string) *PostRemoteAPIJAccountSetup2FactorAuthIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account setup2 factor auth ID params
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountSetup2FactorAuthIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
