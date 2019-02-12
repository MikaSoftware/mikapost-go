package thing_serializer

import (
    "net/http"
    "time"

    "github.com/mikasoftware/mikapost-go/model"
)


// --- PROTECTED ---

// TimeSeriesDatumDetailResponse is the response payload for TimeSeriesDatum data model.
type TimeSeriesDatumDetailResponse struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Timestamp           time.Time `json:"timestamp,omitempty"`
    Value               float64 `json:"value,omitempty"`
    ThingID             uint64 `json:"thing_id,omitempty"`
}

// Function will create our output payload.
func NewTimeSeriesDatumDetailResponse(datum *model.TimeSeriesDatum) *TimeSeriesDatumDetailResponse {
	resp := &TimeSeriesDatumDetailResponse{
        ID:            datum.ID,
        Timestamp:     datum.Timestamp,
        Value:         datum.Value,
        ThingID:       datum.ThingID,
    }
	return resp
}

func (rd *TimeSeriesDatumDetailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
