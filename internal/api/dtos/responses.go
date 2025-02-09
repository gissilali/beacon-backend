package dtos

const (
	ErrorCodeUnsupportedRequest = "UNSUPPORTED_REQUEST"
	ErrorInvalidForm            = "INVALID_FORM"
	ErrorInvalidCredentials     = "INVALID_CREDENTIALS"
	ErrorFailedToIssueTokens    = "FAILED_TO_ISSUE_TOKEN"
	ErrorFailedToSaveTokens     = "FAILED_TO_SAVE_TOKEN"
	ErrorInvalidRefreshToken    = "INVALID_REFRESH_TOKEN"
	ErrorGeneratingAuthTokens   = "FAILED_TO_GENERATE_AUTH_TOKEN"
)

type Response map[string]interface{}

type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}
