// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"net/http"
)

type CheckForCrackerUpdateRequest struct {
	// operating_system
	OperatingSystem *string `queryParam:"style=form,explode=true,name=operating_system"`
	// version
	Version *string `queryParam:"style=form,explode=true,name=version"`
}

func (o *CheckForCrackerUpdateRequest) GetOperatingSystem() *string {
	if o == nil {
		return nil
	}
	return o.OperatingSystem
}

func (o *CheckForCrackerUpdateRequest) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

type CheckForCrackerUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// operating system not found
	CrackerUpdate *components.CrackerUpdate
}

func (o *CheckForCrackerUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckForCrackerUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckForCrackerUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CheckForCrackerUpdateResponse) GetCrackerUpdate() *components.CrackerUpdate {
	if o == nil {
		return nil
	}
	return o.CrackerUpdate
}
