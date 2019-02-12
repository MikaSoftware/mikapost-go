package serializer

import (
    "errors"
    "net/http"

    "github.com/mikasoftware/mikapost-go/model_manager"
)

// RegistrationRequest is the request payload for User data model.
//
// NOTE: It's good practice to have well defined request and response payloads
// so you can manage the specific inputs and outputs for clients, and also gives
// you the opportunity to transform data on input or output, for example
// on request, we'd like to protect certain fields and on output perhaps
// we'd like to include a computed field based on other values that aren't
// in the data model. Also, check out this awesome blog post on struct composition:
// http://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/
type RegistrationRequest struct {
    Email string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
	FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}

func (data *RegistrationRequest) Bind(r *http.Request) error {
    if data.Email == "" {
        return errors.New("Please fill in the email.")
    }
    _, count := model_manager.UserManagerInstance().GetByEmail(data.Email)
    if count > 0 {
        return errors.New("Email is not unique. Please enter another email.")
    }
    if data.Password == "" {
        return errors.New("Please fill in the password.")
    }
    if data.FirstName == "" {
        return errors.New("Please fill in the first name.")
    }
    if data.LastName == "" {
        return errors.New("Please fill in the last name.")
    }
	return nil
}

// RegistrationResponse is the output payload for the API endpoint.
type RegistrationResponse struct {
    TokenString string `json:"token"`
    UserID uint64 `json:"user_id,omitempty" form:"int"`
    Email string `json:"email"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}

// Function will create our output payload.
func NewRegistrationResponse(tokenString string, userID uint64, email string, firstName string, lastName string) *RegistrationResponse {
	resp := &RegistrationResponse{
        TokenString: tokenString,
        UserID: userID,
        Email: email,
        FirstName: firstName,
        LastName: lastName,
    }
	return resp
}

func (rd *RegistrationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
