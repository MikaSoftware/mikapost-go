package serializer

import (
    "net/http"
)


// ProfileResponse is the response payload for User data model.
type ProfileResponse struct {
    UserID uint64 `json:"user_id,omitempty" form:"int"`
    Email string `json:"email" form:"email"`
    FirstName string `json:"first_name,omitempty"`
    LastName string `json:"last_name,omitempty"`
}


// Function will create our output payload.
func NewProfileResponse(userID uint64, email string, firstName string, lastName string) *ProfileResponse {
	resp := &ProfileResponse{
        UserID: userID,
        Email: email,
        FirstName: firstName,
        LastName: lastName,
    }
	return resp
}

func (rd *ProfileResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
