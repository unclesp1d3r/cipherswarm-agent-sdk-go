// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"net/http"
)

type UpdateAgentRequestBody struct {
	// The id of the agent
	ID int64 `json:"id"`
	// The hostname of the agent
	Name string `json:"name"`
	// The signature of the client
	ClientSignature string `json:"client_signature"`
	// The operating system of the agent
	OperatingSystem string   `json:"operating_system"`
	Devices         []string `json:"devices"`
}

func (o *UpdateAgentRequestBody) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *UpdateAgentRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateAgentRequestBody) GetClientSignature() string {
	if o == nil {
		return ""
	}
	return o.ClientSignature
}

func (o *UpdateAgentRequestBody) GetOperatingSystem() string {
	if o == nil {
		return ""
	}
	return o.OperatingSystem
}

func (o *UpdateAgentRequestBody) GetDevices() []string {
	if o == nil {
		return []string{}
	}
	return o.Devices
}

type UpdateAgentRequest struct {
	// id
	ID          int64                   `pathParam:"style=simple,explode=false,name=id"`
	RequestBody *UpdateAgentRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateAgentRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *UpdateAgentRequest) GetRequestBody() *UpdateAgentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateAgentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful
	Agent *components.Agent
}

func (o *UpdateAgentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAgentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAgentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAgentResponse) GetAgent() *components.Agent {
	if o == nil {
		return nil
	}
	return o.Agent
}
