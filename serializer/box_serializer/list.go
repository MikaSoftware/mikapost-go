package box_serializer

import (
    "fmt"
    "net/http"

    "github.com/mikasoftware/mikapost-go/model"
)


// --- PROTECTED ---

// Individual box list response payload.
type BoxListItemPayload struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    ShortDescription    string `json:"short_description,omitempty"`
    LongDescription     string `json:"long_description,omitempty"`
}

// The array of payload structures.
type BoxListDetailsPayload []*BoxListItemPayload

// Full paginated list response payload.
type BoxListPayload struct {
    PageIndex uint64 `json:"page"`
    PagesCount uint64 `json:"pages"`
    TotalRecords uint64 `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Details BoxListDetailsPayload `json:"details"`
}

// Constructor creates a BoxListItemPayload payload from the
// Box model data.
func NewBoxListItemPayload(object *model.Box) *BoxListItemPayload {
	resp := &BoxListItemPayload{
        ID: object.ID,
        Name: object.Name,
        ShortDescription: object.ShortDescription,
        LongDescription: object.LongDescription,
    }
	return resp
}

// Required function to be implemented by "Render" class for the
// "BoxListItemPayload" struct.
func (rd *BoxListItemPayload) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewBoxListDetailsPayload(boxs []model.Box) BoxListDetailsPayload {
	list := BoxListDetailsPayload{}
	for _, box := range boxs {
		list = append(list, NewBoxListItemPayload(&box))
	}
	return list
}


// Constructor creates paginated list response from the "Box" queryset.
func NewBoxListResponse(boxs []model.Box, pageIndex uint64, pagesCount uint64, totalRecords uint64) *BoxListPayload {

    // The following will generate the next or previous URL.
    var next string = ""
    var previous string = ""
    if pageIndex == 0 {
        previous = ""
    } else {
        if pagesCount > 1 {
            previous = fmt.Sprintf("/api/v1/boxes?page=%v", pageIndex-1)
        }
    }
    if pageIndex > pagesCount {
        next = fmt.Sprintf("/api/v1/boxes?page=%v", pageIndex+1)
    }

    // Generate our payload.
	resp := &BoxListPayload{
        PageIndex: pageIndex,
        PagesCount: pagesCount,
        TotalRecords: totalRecords,
        Next: next,
        Previous: previous,
        Details: NewBoxListDetailsPayload(boxs),
    }

    // Return our generated payload.
	return resp
}

// Required function to be implemented by "Render" class for the
// "BoxListPayload" struct.
func (payload *BoxListPayload) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
