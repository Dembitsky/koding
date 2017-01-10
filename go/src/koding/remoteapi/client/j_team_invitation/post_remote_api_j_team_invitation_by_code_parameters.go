package j_team_invitation

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

// NewPostRemoteAPIJTeamInvitationByCodeParams creates a new PostRemoteAPIJTeamInvitationByCodeParams object
// with the default values initialized.
func NewPostRemoteAPIJTeamInvitationByCodeParams() *PostRemoteAPIJTeamInvitationByCodeParams {
	var ()
	return &PostRemoteAPIJTeamInvitationByCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJTeamInvitationByCodeParamsWithTimeout creates a new PostRemoteAPIJTeamInvitationByCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJTeamInvitationByCodeParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJTeamInvitationByCodeParams {
	var ()
	return &PostRemoteAPIJTeamInvitationByCodeParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJTeamInvitationByCodeParamsWithContext creates a new PostRemoteAPIJTeamInvitationByCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJTeamInvitationByCodeParamsWithContext(ctx context.Context) *PostRemoteAPIJTeamInvitationByCodeParams {
	var ()
	return &PostRemoteAPIJTeamInvitationByCodeParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJTeamInvitationByCodeParams contains all the parameters to send to the API endpoint
for the post remote API j team invitation by code operation typically these are written to a http.Request
*/
type PostRemoteAPIJTeamInvitationByCodeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j team invitation by code params
func (o *PostRemoteAPIJTeamInvitationByCodeParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJTeamInvitationByCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j team invitation by code params
func (o *PostRemoteAPIJTeamInvitationByCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j team invitation by code params
func (o *PostRemoteAPIJTeamInvitationByCodeParams) WithContext(ctx context.Context) *PostRemoteAPIJTeamInvitationByCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j team invitation by code params
func (o *PostRemoteAPIJTeamInvitationByCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j team invitation by code params
func (o *PostRemoteAPIJTeamInvitationByCodeParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJTeamInvitationByCodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j team invitation by code params
func (o *PostRemoteAPIJTeamInvitationByCodeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJTeamInvitationByCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
