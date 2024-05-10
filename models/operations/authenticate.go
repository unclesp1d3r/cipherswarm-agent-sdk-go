// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk-go/models/components"
	"net/http"
)

type AuthenticateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful
	AuthenticationResult *components.AuthenticationResult
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

func (o *AuthenticateResponse) GetAuthenticationResult() *components.AuthenticationResult {
	if o == nil {
		return nil
	}
	return o.AuthenticationResult
}
