package model

import (
    _ "github.com/jinzhu/gorm"
)


// Model used to grant access of the time-series data found inside a `thing`
// to any other users in the system. The purpose of this class is to grant
// access to users besides the owner.
type UserThingPermission struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    User                User `gorm:"foreignkey:UserID"` // Model
    UserID              uint64 `gorm:"index; not null;"`
    Thing               User `gorm:"foreignkey:ThingID"` // Model
    ThingID             uint64 `gorm:"index; not null;"`
    Type                uint8 `gorm:"DEFAULT: 1;"`
}

// Give custom table name in our database.
func (u UserThingPermission) TableName() string {
    return "mika_user_thing_permissions"
}

// Type:
// 1 = Read / Write
// 2 = Read Only
