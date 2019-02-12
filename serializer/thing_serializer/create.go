package thing_serializer

import (
    "context"
    "errors"
    "net/http"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/model"
    "github.com/mikasoftware/mikapost-go/model_manager"
)


// ThingCreateRequest is the request payload for Thing data model.
type ThingCreateRequest struct {
    BoxID                uint64 `json:"box_id,string"`
    Name                 string `json:"name"`
    ShortDescription     string `json:"short_description,omitempty"`
    LongDescription      string `json:"long_description,omitempty"`
    UnitOfMeasure        string `json:"unit_of_measure,omitempty"`
    Status               uint8 `json:"status,string"`
    StreetAddress        string `json:"street_address;omitempty"`
    StreetAddressExtra   string `json:"street_address_extra;omitempty"`
    City                 string `json:"city;omitempty"`
    Province             string `json:"province;omitempty"`
    Country              string `json:"country;omitempty"`
    Postal               string `json:"postal;omitempty"`
    IsAddressVisible     bool `json:"is_address_visible;omitempty"`
}


// Function will create Thing data model from the input payload.
func (data *ThingCreateRequest) Save(ctx context.Context) (*model.Thing, error) {
    // Extract the current user from the request context.
    user := ctx.Value("user").(*model.User)

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // The model we will be creating.
    var thing model.Thing

    // Create our object in the database.
    thing = model.Thing {
        BoxID:              data.BoxID,
        Name:               data.Name,
        ShortDescription:   data.ShortDescription,
        LongDescription:    data.LongDescription,
        UnitOfMeasure:      data.UnitOfMeasure,
        User:               *user,
        Status:             data.Status,
        ShareKey:           "",
        StreetAddress:      data.StreetAddress,
        StreetAddressExtra: data.StreetAddressExtra,
        City:               data.City,
        Province:           data.Province,
        Country:            data.Country,
        Postal:             data.Postal,
        IsAddressVisible:   data.IsAddressVisible,
    }
    db.Create(&thing)

    // Create our permission.
    userThingPermission := model.UserThingPermission{
        UserID: user.ID,
        ThingID:  thing.ID,
        Type:   1,
    }
    db.Create(&userThingPermission)

    // Return our newly created `Thing` object.
    return &thing, nil
}


// Function will validate the input payload.
func (data *ThingCreateRequest) Bind(r *http.Request) error {
    // // Extract the current user from the request context.
    user := r.Context().Value("user").(*model.User)

    if data.BoxID == 0 {
        return errors.New("Please fill in the `box_id`.")
    } else {
        hasPermission := model_manager.ThingManagerInstance().HasBoxPermission(user, data.BoxID)
        if hasPermission == false {
            return errors.New("Please select `box_id` that you have permission for.")
        }
    }

    // Validate "Name" field.
    if data.Name == "" {
        return errors.New("Please fill in the name.")
    }

    // Validate "UnitOfMeasure" field.
    if data.UnitOfMeasure == "" {
        return errors.New("Please fill in the unit of measure.")
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
