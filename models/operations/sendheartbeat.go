// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"net/http"
)

type SendHeartbeatRequest struct {
	// id
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *SendHeartbeatRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type SendHeartbeatResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful, but with server feedback
	AgentHeartbeatResponse *components.AgentHeartbeatResponse
}

func (o *SendHeartbeatResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendHeartbeatResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendHeartbeatResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SendHeartbeatResponse) GetAgentHeartbeatResponse() *components.AgentHeartbeatResponse {
	if o == nil {
		return nil
	}
	return o.AgentHeartbeatResponse
}
