package model

import (
    _ "github.com/jinzhu/gorm"
)


type ShareableLink struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Code                string `gorm:"not null; UNIQUE_INDEX; size:60;"`
    Box                 User `gorm:"foreignkey:BoxID"` // Model
    BoxID               uint64 `gorm:"index; ;"`
    Thing               User `gorm:"foreignkey:ThingID"` // Model
    ThingID             uint64 `gorm:"index;"`
    Type                uint8 `gorm:"DEFAULT: 1;"`
}

// Give custom table name in our database.
func (u ShareableLink) TableName() string {
    return "mika_shareable_links"
}

// Type:
// 1 = Read / Write
// 2 = Read Only
// 3 = Disbled
