package dto

// SignUpReq represents the request to sign up a user
type SignUpReq struct {
	Username  string `json:"username" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	PhoneNo   string `json:"phoneNo" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

// SignUpReqHeader represents the request header to sign up a user
type SignUpReqHeader struct {
	AcceptLocale string
}

// SignUpRes represents the response to sign up a user
type SignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Token    string `json:"token"`
}
