package utils

type Message string

const (
	CREATED Message = "Has Been Created Successfully"
	UPDATED Message = "Has Been Modified Successfully"
	DELETED Message = "Has Been Destroyed Successfully"

	SIGNIN Message = "Your Sign In Has Been Successfully"
	SIGNUP Message = "Your Sign Up Has Been Successfully"
	LOGOUT Message = "Your Logout Has Been Successfully"
)

type ResponseSuccess struct {
	Code         int64       `json:"Code,omitempty"`
	Message      Message     `json:"Message,omitempty"`
	TokenAccess  string      `json:"TokenAccess,omitempty"`
	TokenRefresh string      `json:"TokenRefresh,omitempty"`
	Page         int64       `json:"Page,omitempty"`
	PerPage      int64       `json:"PerPage,omitempty"`
	TotalItems   int64       `json:"TotalItems,omitempty"`
	Items        interface{} `json:"Items,omitempty"`
}

type ResponseError struct {
	Code    int64       `json:"Code,omitempty"`
	Message Message     `json:"Message,omitempty"`
	Errors  interface{} `json:"Errors,omitempty"`
}
