package dto

// CoreUserSignUpReq represents the request to sign up a user in the core-user-server service
type CoreUserSignUpReq struct {
	Username       string `json:"username"`
	UID            string `json:"uid"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	PhoneNo        string `json:"phoneNo"`
	Email          string `json:"email"`
	SelectedLocale string `json:"selectedLocale"`
}

// CoreUserSignUpRes represents the response to sign up a user in the core-user-server service
type CoreUserSignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Success  bool   `json:"success"`
}
