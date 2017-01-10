package j_group

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

// NewPostRemoteAPIJGroupCountParams creates a new PostRemoteAPIJGroupCountParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupCountParams() *PostRemoteAPIJGroupCountParams {
	var ()
	return &PostRemoteAPIJGroupCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupCountParamsWithTimeout creates a new PostRemoteAPIJGroupCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupCountParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupCountParams {
	var ()
	return &PostRemoteAPIJGroupCountParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupCountParamsWithContext creates a new PostRemoteAPIJGroupCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupCountParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupCountParams {
	var ()
	return &PostRemoteAPIJGroupCountParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupCountParams contains all the parameters to send to the API endpoint
for the post remote API j group count operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupCountParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group count params
func (o *PostRemoteAPIJGroupCountParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group count params
func (o *PostRemoteAPIJGroupCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group count params
func (o *PostRemoteAPIJGroupCountParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group count params
func (o *PostRemoteAPIJGroupCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j group count params
func (o *PostRemoteAPIJGroupCountParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJGroupCountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j group count params
func (o *PostRemoteAPIJGroupCountParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
