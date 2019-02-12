package serializer

import (
    "context"
    // "errors"
    "net/http"
    "time"
    "github.com/go-chi/render"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/model"
    // "github.com/mikasoftware/mikapost-go/model_manager"
)


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//

// Individual time_series_data list response payload.
type TimeSeriesDatumListItemResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Timestamp           time.Time `json:"timestamp,omitempty"`
    Value               float64 `json:"value,omitempty"`
    ThingID             uint64 `json:"thing_id,omitempty"`
}

// Constructor creates a TimeSeriesDatumListItemResponse payload from the
// TimeSeriesDatum model data.
func NewTimeSeriesDatumListItemResponse(object *model.TimeSeriesDatum) *TimeSeriesDatumListItemResponse {
	resp := &TimeSeriesDatumListItemResponse{
        ID: object.ID,
        Timestamp: object.Timestamp,
        Value: object.Value,
        ThingID: object.ThingID,
    }
	return resp
}

func (rd *TimeSeriesDatumListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

// Constructor creates a JSON response payload from the array of TimeSeriesDatum
// model data objects.
func NewTimeSeriesDatumListResponse(time_series_datas []model.TimeSeriesDatum) []render.Renderer {
	list := []render.Renderer{}
	for _, time_series_data := range time_series_datas {
		list = append(list, NewTimeSeriesDatumListItemResponse(&time_series_data))
	}
	return list
}


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


// Function will create TimeSeriesDatum data model from the input payload.
func (data *TimeSeriesDatumRequest) Save(ctx context.Context) (*model.TimeSeriesDatum, error) {
    // // Extract the current user from the request context.
    // user := ctx.Value("user").(*model.User)

    // The model we will be creating.
    var time_series_data model.TimeSeriesDatum

    // Create our `User` object in our database.
    time_series_data = model.TimeSeriesDatum {
        Timestamp:          data.Timestamp,
        Value:              data.Value,
        ThingID:            data.ThingID,
    }

    // Get our database connection.
    dao := database.Instance()
    db := dao.GetORM()

    // Create our object in the database.
    db.Create(&time_series_data)

    return &time_series_data, nil
}

// TimeSeriesDatumRequest is the request payload for TimeSeriesDatum data model.
type TimeSeriesDatumRequest struct {
    Timestamp           time.Time `json:"timestamp,omitempty"`
    Value               float64 `json:"value,omitempty"`
    ThingID             uint64 `json:"thing_id,omitempty"`
}

// Function will validate the input payload.
func (data *TimeSeriesDatumRequest) Bind(r *http.Request) error {
    // // Extract the current user from the request context.
    // user := r.Context().Value("user").(*model.User)

    // Return with no errors.
	return nil
}


//----------------------------------------------------------------------------//
//                                  DETAILS                                   //
//----------------------------------------------------------------------------//


// TimeSeriesDatumDetailResponse is the response payload for TimeSeriesDatum data model.
type TimeSeriesDatumDetailResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Timestamp           time.Time `json:"timestamp,omitempty"`
    Value               float64 `json:"value,omitempty"`
    ThingID             uint64 `json:"thing_id,omitempty"`
}

// Function will create our output payload.
func NewTimeSeriesDatumDetailResponse(time_series_data *model.TimeSeriesDatum) *TimeSeriesDatumDetailResponse {
	resp := &TimeSeriesDatumDetailResponse{
        ID:            time_series_data.ID,
        Timestamp:     time_series_data.Timestamp,
        Value:         time_series_data.Value,
        ThingID:       time_series_data.ThingID,
    }
	return resp
}

func (rd *TimeSeriesDatumDetailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
