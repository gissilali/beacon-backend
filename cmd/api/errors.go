package main

type envelope map[string]interface{}

type ErrorEnvelope struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

const (
	ErrorCodeUnsupportedRequest = "UNSUPPORTED_REQUEST"
	ErrorInvalidForm            = "INVALID_FORM"
	ErrorInvalidCredentials     = "INVALID_CREDENTIALS"
)
