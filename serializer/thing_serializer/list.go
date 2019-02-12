package thing_serializer

import (
    "fmt"
    "net/http"

    "github.com/mikasoftware/mikapost-go/model"
)


// --- PROTECTED ---

// Individual thing list response payload.
type ThingListItemPayload struct {
    ID                  uint64 `json:"id,omitempty" form:"int"`
    Name                string `json:"name,omitempty"`
    ShortDescription    string `json:"short_description,omitempty"`
    UnitOfMeasure       string `json:"unit_of_measure,omitempty"`
}

// The array of payload structures.
type ThingListDetailsPayload []*ThingListItemPayload

// Full paginated list response payload.
type ThingListPayload struct {
    PageIndex uint64 `json:"page"`
    PagesCount uint64 `json:"pages"`
    TotalRecords uint64 `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Details ThingListDetailsPayload `json:"details"`
}

// Constructor creates a ThingListItemPayload payload from the
// Thing model data.
func NewThingListItemPayload(object *model.Thing) *ThingListItemPayload {
	resp := &ThingListItemPayload{
        ID: object.ID,
        Name: object.Name,
        ShortDescription: object.ShortDescription,
        UnitOfMeasure: object.UnitOfMeasure,
    }
	return resp
}

// Required function to be implemented by "Render" class for the
// "ThingListItemPayload" struct.
func (rd *ThingListItemPayload) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func NewThingListDetailsPayload(things []model.Thing) ThingListDetailsPayload {
	list := ThingListDetailsPayload{}
	for _, thing := range things {
		list = append(list, NewThingListItemPayload(&thing))
	}
	return list
}


// Constructor creates paginated list response from the "Thing" queryset.
func NewThingListResponse(things []model.Thing, pageIndex uint64, pagesCount uint64, totalRecords uint64) *ThingListPayload {

    // The following will generate the next or previous URL.
    var next string = ""
    var previous string = ""
    if pageIndex == 0 {
        previous = ""
    } else {
        if pagesCount > 1 {
            previous = fmt.Sprintf("/api/v1/things?page=%v", pageIndex-1)
        }
    }
    if pageIndex > pagesCount {
        next = fmt.Sprintf("/api/v1/things?page=%v", pageIndex+1)
    }

    // Generate our payload.
	resp := &ThingListPayload{
        PageIndex: pageIndex,
        PagesCount: pagesCount,
        TotalRecords: totalRecords,
        Next: next,
        Previous: previous,
        Details: NewThingListDetailsPayload(things),
    }

    // Return our generated payload.
	return resp
}

// Required function to be implemented by "Render" class for the
// "ThingListPayload" struct.
func (rd *ThingListPayload) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
