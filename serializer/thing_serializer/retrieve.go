package thing_serializer

import (
    "net/http"

    "github.com/mikasoftware/mikapost-go/model"
)


// --- PROTECTED ---

// ThingDetailResponse is the response payload for Thing data model.
type ThingDetailResponse struct {
    ID                   uint64 `json:"id"; form:"int"`
    BoxID                uint64 `json:"box_id"; form:"int"`
    Name                 string `json:"name"`
    ShortDescription     string `json:"short_description,omitempty"`
    LongDescription      string `json:"long_description,omitempty"`
    UnitOfMeasure        string `json:"unit_of_measure"`
    Status               uint8 `json:"status"`
    UserID               uint64 `json:"user_id"`
    ShareKey             string `json:"shared_key;omitempty"`
    StreetAddress        string `json:"street_address;omitempty"`
    StreetAddressExtra   string `json:"street_address_extra;omitempty"`
    City                 string `json:"city;omitempty"`
    Province             string `json:"province;omitempty"`
    Country              string `json:"country;omitempty"`
    Postal               string `json:"postal;omitempty"`
    IsAddressVisible     bool `json:"is_address_visible;omitempty"`
}

// Function will create our output payload.
func NewThingDetailResponse(thing *model.Thing) *ThingDetailResponse {
	resp := &ThingDetailResponse{
        ID:                 thing.ID,
        BoxID:              thing.BoxID,
        Name:               thing.Name,
        ShortDescription:   thing.ShortDescription,
        LongDescription:    thing.LongDescription,
        UnitOfMeasure:      thing.UnitOfMeasure,
        Status:             thing.Status,
        UserID:             thing.UserID,
        ShareKey:           thing.ShareKey,
        StreetAddress:      thing.StreetAddress,
        StreetAddressExtra: thing.StreetAddressExtra,
        City:               thing.City,
        Province:           thing.Province,
        Country:            thing.Country,
        Postal:             thing.Postal,
        IsAddressVisible:   thing.IsAddressVisible,
    }
	return resp
}

func (rd *ThingDetailResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
