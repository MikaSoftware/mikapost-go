package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)


// Model represents a collection of "Thing" objects. An alias for this model
// would be "Folder". The purpose
type Box struct {
    ID                   uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Name                 string `gorm:"not null; size:63;"`
    ShortDescription     string `gorm:"size:127;"`
    LongDescription      string `gorm:"type:text"`
    Status               uint8 `gorm:"DEFAULT: 1; not null;"`
    CreatedAt            time.Time
    UpdatedAt            time.Time
    User                 User `gorm:"foreignkey:UserID"` // Model
    UserID               uint64 `gorm:"index; not null;"`
    UserBoxPermissions   []UserBoxPermission // Model
}

// Give custom table name in our database.
func (u Box) TableName() string {
    return "mika_boxes"
}

// Status:
// 1 = Active (Private)
// 2 = Active (Public)
// 3 = Archived
// 4 = Deleted
