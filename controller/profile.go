package controller

import (
	"net/http"
	"github.com/go-chi/render"
	
    "github.com/mikasoftware/mikapost-go/model"
    "github.com/mikasoftware/mikapost-go/serializer"
)


func ProfileRetrieveFunc(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*model.User)

    // Take our data and serialize it back into a response object to hand
    // back to the user.
    render.Status(r, http.StatusOK)
	render.Render(w, r, serializer.NewProfileResponse(
        user.ID,
        user.Email,
        user.FirstName,
        user.LastName,
    ))
}
