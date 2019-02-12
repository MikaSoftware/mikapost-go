package box_serializer

import (
    "net/http"

    "github.com/mikasoftware/mikapost-go/model"
)


// --- PROTECTED ---

// BoxDetailResponse is the response payload for Box data model.
type BoxDetailResponse struct {
    ID                   uint64 `json:"id,omitempty" form:"int"`
    Name                 string `json:"name,omitempty"`
    ShortDescription     string `json:"short_description,omitempty"`
    LongDescription      string `json:"description,omitempty"`
    Status               uint8 `json:"status"`
}

// Function will create our output payload.
func NewBoxDetailResponse(box *model.Box) *BoxDetailResponse {
	resp := &BoxDetailResponse{
        ID:                 box.ID,
        Name:               box.Name,
        ShortDescription:   box.ShortDescription,
        LongDescription:    box.LongDescription,
        Status:             box.Status,
    }
	return resp
}

func (rd *BoxDetailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
