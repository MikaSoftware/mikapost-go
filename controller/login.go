package controller

import (
    // "errors"
    "net/http"
    "github.com/go-chi/render"

    "github.com/mikasoftware/mikapost-go/model_manager"
    "github.com/mikasoftware/mikapost-go/serializer"
    "github.com/mikasoftware/mikapost-go/base/service"
)

// Function will authenticate the user loging credentials and return a JWT
// token with successfully authenticated or else return an error message.
func LoginFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST submission and unserialize it.
    data := &serializer.LoginRequest{}
    if err := render.Bind(r, data); err != nil {
        render.Render(w, r, serializer.ErrInvalidRequest(err))
		return
	}

    // Fetch our user object.
    user, _ := model_manager.UserManagerInstance().GetByEmail(data.Email)

    // Create the users authentication token.
    tokenString := service.GenerateJWTToken(user.ID)

    // Take our data and serialize it back into a response object to hand
    // back to the user.
    render.Status(r, http.StatusOK)
	render.Render(w, r, serializer.NewLoginResponse(
        tokenString,
        user.ID,
        user.Email,
        user.FirstName,
        user.LastName,
    ))
}
