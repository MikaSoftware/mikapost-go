package thing_serializer

import (
    "context"
    "errors"
    "net/http"
    "time"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/model"
    "github.com/mikasoftware/mikapost-go/model_manager"
)


// TimeSeriesDatumCreateRequest is the request payload for TimeSeriesDatum data model.
type TimeSeriesDatumCreateRequest struct {
    Timestamp   time.Time `json:"timestamp,omitempty"`
    Value       float64 `json:"value,omitempty,string"`
    ThingID     uint64 `json:"thing_id,omitempty,string"`
}


// Function will create TimeSeriesDatum data model from the input payload.
func (data *TimeSeriesDatumCreateRequest) Save(ctx context.Context) (*model.TimeSeriesDatum, error) {
    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // // Extract the current user from the request context.
    // user := ctx.Value("user").(*model.User)

    // The model we will be creating.
    var thing model.TimeSeriesDatum

    // Create our object in the database.
    thing = model.TimeSeriesDatum {
        Timestamp: data.Timestamp,
        Value:     data.Value,
        ThingID:   data.ThingID,
    }
    db.Create(&thing)

    // Return our newly created `Thing` object.
    return &thing, nil
}


// Function will validate the input payload.
func (data *TimeSeriesDatumCreateRequest) Bind(r *http.Request) error {
    // Extract the current user from the request context.
    user := r.Context().Value("user").(*model.User)

    // Validate "ThingID" field.
    if data.ThingID == 0 {
        return errors.New("Please fill in the `thing_id`.")
    } else {
        hasPermission := model_manager.TimeSeriesDatumManagerInstance().HasPermission(user.ID, data.ThingID)
        if hasPermission == false {
            return errors.New("Please select `thing_id` that you have permission for.")
        }
    }

    // Return with no errors.
	return nil
}
