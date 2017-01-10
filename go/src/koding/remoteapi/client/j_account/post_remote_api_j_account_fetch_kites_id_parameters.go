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

// NewPostRemoteAPIJAccountFetchKitesIDParams creates a new PostRemoteAPIJAccountFetchKitesIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountFetchKitesIDParams() *PostRemoteAPIJAccountFetchKitesIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchKitesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountFetchKitesIDParamsWithTimeout creates a new PostRemoteAPIJAccountFetchKitesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountFetchKitesIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchKitesIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchKitesIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountFetchKitesIDParamsWithContext creates a new PostRemoteAPIJAccountFetchKitesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountFetchKitesIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountFetchKitesIDParams {
	var ()
	return &PostRemoteAPIJAccountFetchKitesIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountFetchKitesIDParams contains all the parameters to send to the API endpoint
for the post remote API j account fetch kites ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountFetchKitesIDParams struct {

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

// WithTimeout adds the timeout to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountFetchKitesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountFetchKitesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJAccountFetchKitesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) WithID(id string) *PostRemoteAPIJAccountFetchKitesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account fetch kites ID params
func (o *PostRemoteAPIJAccountFetchKitesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountFetchKitesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
