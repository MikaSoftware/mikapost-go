package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)


// Model to represent any "Thing" which has time-series data associated with it,
// for example financial data like "account balance over time", or "Internet of
// Things (IoT) data like "relative humidity data over time", etc.
type Thing struct {
    ID                   uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                 string `gorm:"not null; size:63;"`
    ShortDescription     string `gorm:"size:127;"`
    LongDescription      string `gorm:"type:text"`
    UnitOfMeasure        string `gorm:"type:varchar(127); not null;"`
    Status               uint8 `gorm:"DEFAULT: 1; not null;"`
    ShareKey             string `gorm:"size:127;"`
    StreetAddress        string `gorm:"size:127;"`
    StreetAddressExtra   string `gorm:"size:127;"`
    City                 string `gorm:"size:127;"`
    Province             string `gorm:"size:127;"`
    Country              string `gorm:"size:127;"`
    Postal               string `gorm:"size:31;"`
    IsAddressVisible     bool `gorm:"type:boolean;DEFAULT: false;"`
    CreatedAt            time.Time
    UpdatedAt            time.Time
    Box                  Box `gorm:"foreignkey:BoxID"` // Model
    BoxID                uint64 `gorm:"index; not null;"`
    User                 User `gorm:"foreignkey:UserID"` // Model
    UserID               uint64 `gorm:"index; not null;"`
    UserThingPermissions []UserThingPermission // Model
}

// Give custom table name in our database.
func (u Thing) TableName() string {
    return "mika_things"
}

// Status:
// 1 = Active (Private)
// 2 = Active (Public)
// 3 = Inactive,
// 4 = Deleted
