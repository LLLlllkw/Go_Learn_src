// Code generated by go-swagger; DO NOT EDIT.

package session

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

	"courseWork/cmd/models"
)

// NewGetPlaybackStatusesBySessionIdsParams creates a new GetPlaybackStatusesBySessionIdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlaybackStatusesBySessionIdsParams() *GetPlaybackStatusesBySessionIdsParams {
	return &GetPlaybackStatusesBySessionIdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlaybackStatusesBySessionIdsParamsWithTimeout creates a new GetPlaybackStatusesBySessionIdsParams object
// with the ability to set a timeout on a request.
func NewGetPlaybackStatusesBySessionIdsParamsWithTimeout(timeout time.Duration) *GetPlaybackStatusesBySessionIdsParams {
	return &GetPlaybackStatusesBySessionIdsParams{
		timeout: timeout,
	}
}

// NewGetPlaybackStatusesBySessionIdsParamsWithContext creates a new GetPlaybackStatusesBySessionIdsParams object
// with the ability to set a context for a request.
func NewGetPlaybackStatusesBySessionIdsParamsWithContext(ctx context.Context) *GetPlaybackStatusesBySessionIdsParams {
	return &GetPlaybackStatusesBySessionIdsParams{
		Context: ctx,
	}
}

// NewGetPlaybackStatusesBySessionIdsParamsWithHTTPClient creates a new GetPlaybackStatusesBySessionIdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlaybackStatusesBySessionIdsParamsWithHTTPClient(client *http.Client) *GetPlaybackStatusesBySessionIdsParams {
	return &GetPlaybackStatusesBySessionIdsParams{
		HTTPClient: client,
	}
}

/* GetPlaybackStatusesBySessionIdsParams contains all the parameters to send to the API endpoint
   for the get playback statuses by session ids operation.

   Typically these are written to a http.Request.
*/
type GetPlaybackStatusesBySessionIdsParams struct {

	/* SessionIds.

	   连线 ID 的数组
	*/
	SessionIds *models.SessionIds

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get playback statuses by session ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaybackStatusesBySessionIdsParams) WithDefaults() *GetPlaybackStatusesBySessionIdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get playback statuses by session ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlaybackStatusesBySessionIdsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) WithTimeout(timeout time.Duration) *GetPlaybackStatusesBySessionIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) WithContext(ctx context.Context) *GetPlaybackStatusesBySessionIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) WithHTTPClient(client *http.Client) *GetPlaybackStatusesBySessionIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSessionIds adds the sessionIds to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) WithSessionIds(sessionIds *models.SessionIds) *GetPlaybackStatusesBySessionIdsParams {
	o.SetSessionIds(sessionIds)
	return o
}

// SetSessionIds adds the sessionIds to the get playback statuses by session ids params
func (o *GetPlaybackStatusesBySessionIdsParams) SetSessionIds(sessionIds *models.SessionIds) {
	o.SessionIds = sessionIds
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlaybackStatusesBySessionIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.SessionIds != nil {
		if err := r.SetBodyParam(o.SessionIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
