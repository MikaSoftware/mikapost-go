package thing_serializer

import (
    "fmt"
    "net/http"
    "time"

    "github.com/mikasoftware/mikapost-go/model"
)


// --- PROTECTED ---

// Individual thing list response payload.
type TimeSeriesDataListItemPayload struct {
    ID                  uint64 `json:"id,omitempty"; form:"int"`
    Timestamp           time.Time `json:"timestamp,omitempty"`
    Value               float64 `json:"value"; form:"float"`
    ThingID             uint64 `json:"thing_id"`
}

// The array of payload structures.
type TimeSeriesDataListDetailsPayload []*TimeSeriesDataListItemPayload

// Full paginated list response payload.
type TimeSeriesDataListPayload struct {
    PageIndex uint64 `json:"page"`
    PagesCount uint64 `json:"pages"`
    TotalRecords uint64 `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Details TimeSeriesDataListDetailsPayload `json:"details"`
}

// Constructor creates a TimeSeriesDataListItemPayload payload from the
// TimeSeriesData model data.
func NewTimeSeriesDataListItemPayload(object *model.TimeSeriesDatum) *TimeSeriesDataListItemPayload {
	resp := &TimeSeriesDataListItemPayload{
        ID: object.ID,
        Timestamp: object.Timestamp,
        Value: object.Value,
        ThingID: object.ThingID,
    }
	return resp
}

// Required function to be implemented by "Render" class for the
// "TimeSeriesDataListItemPayload" struct.
func (rd *TimeSeriesDataListItemPayload) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewTimeSeriesDataListDetailsPayload(things []model.TimeSeriesDatum) TimeSeriesDataListDetailsPayload {
	list := TimeSeriesDataListDetailsPayload{}
	for _, thing := range things {
		list = append(list, NewTimeSeriesDataListItemPayload(&thing))
	}
	return list
}


// Constructor creates paginated list response from the "TimeSeriesData" queryset.
func NewTimeSeriesDataListResponse(things []model.TimeSeriesDatum, pageIndex uint64, pagesCount uint64, totalRecords uint64) *TimeSeriesDataListPayload {

    // The following will generate the next or previous URL.
    var next string = ""
    var previous string = ""
    if pageIndex == 0 {
        previous = ""
    } else {
        if pagesCount > 1 {
            previous = fmt.Sprintf("/api/v1/data?page=%v", pageIndex-1)
        }
    }
    if pageIndex > pagesCount {
        next = fmt.Sprintf("/api/v1/data?page=%v", pageIndex+1)
    }

    // Generate our payload.
	resp := &TimeSeriesDataListPayload{
        PageIndex: pageIndex,
        PagesCount: pagesCount,
        TotalRecords: totalRecords,
        Next: next,
        Previous: previous,
        Details: NewTimeSeriesDataListDetailsPayload(things),
    }

    // Return our generated payload.
	return resp
}

// Required function to be implemented by "Render" class for the
// "TimeSeriesDataListPayload" struct.
func (rd *TimeSeriesDataListPayload) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
