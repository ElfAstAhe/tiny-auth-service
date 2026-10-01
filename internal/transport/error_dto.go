package transport

// ErrorDTO model info
// @Description Error dto
type ErrorDTO struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
} //@name ErrorDTO

// NewErrorDTO acts as a factory constructor creating a standardized transport-level error response instance.
func NewErrorDTO(status int, message string) *ErrorDTO {
	return &ErrorDTO{
		Code:    status,
		Message: message,
	}
}

// NewErrorDTOFromError converts a raw error instance into a structured transport-level data transfer object payload.
func NewErrorDTOFromError(code int, err error) *ErrorDTO {
	return NewErrorDTO(code, err.Error())
}
