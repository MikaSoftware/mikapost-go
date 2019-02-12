package serializer

import (
    "errors"
    "net/http"
    
    "github.com/mikasoftware/mikapost-go/base/service"
    "github.com/mikasoftware/mikapost-go/model_manager"
)

// Input payload.
type LoginRequest struct {
    Email string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
}

// Function will validate the input payload.
func (data *LoginRequest) Bind(r *http.Request) error {
    if data.Email == "" {
        return errors.New("Missing email.")
    }
    // Check to see if the user exists in the database.
    user, count := model_manager.UserManagerInstance().GetByEmail(data.Email)
    if count == 0 {
        return errors.New("Email or password is incorrect. (1)")
    }
    if data.Password == "" {
        return errors.New("Missing password.")
    }
    // Try user password.
    passwordsMatch := service.CheckPasswordHash(data.Password, user.PasswordHash)
    if passwordsMatch == false {
        return errors.New("Email or password is incorrect. (2)")
    }
	return nil
}

// Output payload.
type LoginResponse struct {
    TokenString string `json:"token" form:"string"`
    UserID uint64 `json:"user_id,omitempty" form:"int"`
    Email string `json:"email" form:"email"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}

// Function will create our output payload.
func NewLoginResponse(tokenString string, userID uint64, email string, firstName string, lastName string) *LoginResponse {
	resp := &LoginResponse{
        TokenString: tokenString,
        UserID: userID,
        Email: email,
        FirstName: firstName,
        LastName: lastName,
    }
	return resp
}

func (rd *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
