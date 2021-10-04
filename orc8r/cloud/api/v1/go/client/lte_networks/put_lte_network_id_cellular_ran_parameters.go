// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutLTENetworkIDCellularRanParams creates a new PutLTENetworkIDCellularRanParams object
// with the default values initialized.
func NewPutLTENetworkIDCellularRanParams() *PutLTENetworkIDCellularRanParams {
	var ()
	return &PutLTENetworkIDCellularRanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDCellularRanParamsWithTimeout creates a new PutLTENetworkIDCellularRanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDCellularRanParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDCellularRanParams {
	var ()
	return &PutLTENetworkIDCellularRanParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDCellularRanParamsWithContext creates a new PutLTENetworkIDCellularRanParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDCellularRanParamsWithContext(ctx context.Context) *PutLTENetworkIDCellularRanParams {
	var ()
	return &PutLTENetworkIDCellularRanParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDCellularRanParamsWithHTTPClient creates a new PutLTENetworkIDCellularRanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDCellularRanParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDCellularRanParams {
	var ()
	return &PutLTENetworkIDCellularRanParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDCellularRanParams contains all the parameters to send to the API endpoint
for the put LTE network ID cellular ran operation typically these are written to a http.Request
*/
type PutLTENetworkIDCellularRanParams struct {

	/*Config
	  New RAN configuration for the network

	*/
	Config *models.NetworkRanConfigs
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDCellularRanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) WithContext(ctx context.Context) *PutLTENetworkIDCellularRanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDCellularRanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) WithConfig(config *models.NetworkRanConfigs) *PutLTENetworkIDCellularRanParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) SetConfig(config *models.NetworkRanConfigs) {
	o.Config = config
}

// WithNetworkID adds the networkID to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) WithNetworkID(networkID string) *PutLTENetworkIDCellularRanParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID cellular ran params
func (o *PutLTENetworkIDCellularRanParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDCellularRanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}