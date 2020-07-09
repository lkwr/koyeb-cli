// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLogsTailStackRevisionLogsParams creates a new LogsTailStackRevisionLogsParams object
// with the default values initialized.
func NewLogsTailStackRevisionLogsParams() *LogsTailStackRevisionLogsParams {
	var ()
	return &LogsTailStackRevisionLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLogsTailStackRevisionLogsParamsWithTimeout creates a new LogsTailStackRevisionLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLogsTailStackRevisionLogsParamsWithTimeout(timeout time.Duration) *LogsTailStackRevisionLogsParams {
	var ()
	return &LogsTailStackRevisionLogsParams{

		timeout: timeout,
	}
}

// NewLogsTailStackRevisionLogsParamsWithContext creates a new LogsTailStackRevisionLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLogsTailStackRevisionLogsParamsWithContext(ctx context.Context) *LogsTailStackRevisionLogsParams {
	var ()
	return &LogsTailStackRevisionLogsParams{

		Context: ctx,
	}
}

// NewLogsTailStackRevisionLogsParamsWithHTTPClient creates a new LogsTailStackRevisionLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLogsTailStackRevisionLogsParamsWithHTTPClient(client *http.Client) *LogsTailStackRevisionLogsParams {
	var ()
	return &LogsTailStackRevisionLogsParams{
		HTTPClient: client,
	}
}

/*LogsTailStackRevisionLogsParams contains all the parameters to send to the API endpoint
for the logs tail stack revision logs operation typically these are written to a http.Request
*/
type LogsTailStackRevisionLogsParams struct {

	/*Sha*/
	Sha string
	/*StackID*/
	StackID string
	/*Start*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) WithTimeout(timeout time.Duration) *LogsTailStackRevisionLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) WithContext(ctx context.Context) *LogsTailStackRevisionLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) WithHTTPClient(client *http.Client) *LogsTailStackRevisionLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSha adds the sha to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) WithSha(sha string) *LogsTailStackRevisionLogsParams {
	o.SetSha(sha)
	return o
}

// SetSha adds the sha to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) SetSha(sha string) {
	o.Sha = sha
}

// WithStackID adds the stackID to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) WithStackID(stackID string) *LogsTailStackRevisionLogsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithStart adds the start to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) WithStart(start *string) *LogsTailStackRevisionLogsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the logs tail stack revision logs params
func (o *LogsTailStackRevisionLogsParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *LogsTailStackRevisionLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sha
	if err := r.SetPathParam("sha", o.Sha); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
