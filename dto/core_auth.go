package dto

const coreAuthContentType = "application/json"

// CoreAuthSignUpReq represents the request to sign up a user in the core-auth-server service
type CoreAuthSignUpReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phoneNo"`
}

// CoreAuthSignUpReqHeader represents the request header to sign up a user in the core-auth-server service
type CoreAuthSignUpReqHeader struct {
	ContentType  string
	AcceptLocale string
}

// ToMap converts the CoreAuthSignUpReqHeader to a map
func (h CoreAuthSignUpReqHeader) ToMap() map[string]string {
	m := make(map[string]string)
	m["Accept-Locale"] = h.AcceptLocale
	if h.ContentType == "" {
		m["Content-Type"] = coreAuthContentType
	} else {
		m["Content-Type"] = h.ContentType
	}
	return m
}

// CoreAuthSignUpRes represents the response to sign up a user in the core-auth-server service
type CoreAuthSignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Token    string `json:"token"`
}

// CoreAuthDeleteUserReq represents the request to delete a user in the core-auth-server service
type CoreAuthDeleteUserReq struct {
	UID string `json:"uid"`
}

// CoreAuthDeleteUserRes represents the response to delete a user in the core-auth-server service
type CoreAuthDeleteUserRes struct {
	Success bool `json:"success"`
}
