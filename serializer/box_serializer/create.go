package box_serializer

import (
    "context"
    "errors"
    "net/http"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/model"
)


// BoxCreateRequest is the request payload for Box data model.
type BoxCreateRequest struct {
    Name                 string `json:"name"`
    ShortDescription     string `json:"short_description,omitempty"`
    LongDescription      string `json:"long_description,omitempty"`
    Status               uint8 `json:"status,string"`
}


// Function will create Box data model from the input payload.
func (data *BoxCreateRequest) Save(ctx context.Context) (*model.Box, error) {
    // Extract the current user from the request context.
    user := ctx.Value("user").(*model.User)

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our `Box` object in our database from our struct.
    box := model.Box{
        Name:               data.Name,
        ShortDescription:   data.ShortDescription,
        LongDescription:    data.LongDescription,
        User:               *user,
        Status:             data.Status,
    }
    db.Create(&box)

    // Create our permission.
    userBoxPermission := model.UserBoxPermission{
        UserID: user.ID,
        BoxID:  box.ID,
        Type:   1,
    }
    db.Create(&userBoxPermission)

    // Return our newly created `Box` object.
    return &box, nil
}


// Function will validate the input payload.
func (data *BoxCreateRequest) Bind(r *http.Request) error {
    // Validate "Name" field.
    if data.Name == "" {
        return errors.New("Please fill in the name.")
    }

    // Validate "Status" field.
    validStatus := false
    statusOptionsArr := [4]uint8{1, 2, 3, 4}
    for _, v := range statusOptionsArr {
        if v == data.Status {
            validStatus = true
        }
    }
    if validStatus == false {
        return errors.New("Please select a valid status option.")
    }

    // Return with no errors.
	return nil
}
