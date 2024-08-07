// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

// AuthenticateResponseBody - successful
type AuthenticateResponseBody struct {
	Authenticated bool  `json:"authenticated"`
	AgentID       int64 `json:"agent_id"`
}

func (o *AuthenticateResponseBody) GetAuthenticated() bool {
	if o == nil {
		return false
	}
	return o.Authenticated
}

func (o *AuthenticateResponseBody) GetAgentID() int64 {
	if o == nil {
		return 0
	}
	return o.AgentID
}

type AuthenticateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful
	Object *AuthenticateResponseBody
}

func (o *AuthenticateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AuthenticateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AuthenticateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AuthenticateResponse) GetObject() *AuthenticateResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
