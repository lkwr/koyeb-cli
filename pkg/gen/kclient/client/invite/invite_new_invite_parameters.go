// Code generated by go-swagger; DO NOT EDIT.

package invite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/koyeb/koyeb-cli/pkg/gen/kclient/models"
)

// NewInviteNewInviteParams creates a new InviteNewInviteParams object
// with the default values initialized.
func NewInviteNewInviteParams() *InviteNewInviteParams {
	var ()
	return &InviteNewInviteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInviteNewInviteParamsWithTimeout creates a new InviteNewInviteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInviteNewInviteParamsWithTimeout(timeout time.Duration) *InviteNewInviteParams {
	var ()
	return &InviteNewInviteParams{

		timeout: timeout,
	}
}

// NewInviteNewInviteParamsWithContext creates a new InviteNewInviteParams object
// with the default values initialized, and the ability to set a context for a request
func NewInviteNewInviteParamsWithContext(ctx context.Context) *InviteNewInviteParams {
	var ()
	return &InviteNewInviteParams{

		Context: ctx,
	}
}

// NewInviteNewInviteParamsWithHTTPClient creates a new InviteNewInviteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInviteNewInviteParamsWithHTTPClient(client *http.Client) *InviteNewInviteParams {
	var ()
	return &InviteNewInviteParams{
		HTTPClient: client,
	}
}

/*InviteNewInviteParams contains all the parameters to send to the API endpoint
for the invite new invite operation typically these are written to a http.Request
*/
type InviteNewInviteParams struct {

	/*Body*/
	Body *models.AccountInviteUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the invite new invite params
func (o *InviteNewInviteParams) WithTimeout(timeout time.Duration) *InviteNewInviteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite new invite params
func (o *InviteNewInviteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite new invite params
func (o *InviteNewInviteParams) WithContext(ctx context.Context) *InviteNewInviteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite new invite params
func (o *InviteNewInviteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite new invite params
func (o *InviteNewInviteParams) WithHTTPClient(client *http.Client) *InviteNewInviteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite new invite params
func (o *InviteNewInviteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invite new invite params
func (o *InviteNewInviteParams) WithBody(body *models.AccountInviteUserRequest) *InviteNewInviteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invite new invite params
func (o *InviteNewInviteParams) SetBody(body *models.AccountInviteUserRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InviteNewInviteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}