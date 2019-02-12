package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)


// Model represents a single unit of data for a particular "Thing" model.
type TimeSeriesDatum struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Timestamp           time.Time `gorm:"not null;"`
    Value               float64 `gorm:""`
    Thing               Thing `gorm:"foreignkey:ThingID"` // Model
    ThingID             uint64 `gorm:"index; not null;"`
}

// Give custom table name in our database.
func (u TimeSeriesDatum) TableName() string {
    return "mika_time_series_data"
}
