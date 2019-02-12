package model

import (
    _ "github.com/jinzhu/gorm"
)


type UserBoxPermission struct {
    ID                  uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    User                User `gorm:"foreignkey:UserID"` // Model
    UserID              uint64 `gorm:"index; not null;"`
    Box                 User `gorm:"foreignkey:BoxID"` // Model
    BoxID               uint64 `gorm:"index; not null;"`
    Type                uint8 `gorm:"DEFAULT: 1;"`
}


// Give custom table name in our database.
func (u UserBoxPermission) TableName() string {
    return "mika_user_box_permissions"
}


// Type:
// 1 = Read / Write (Owner)
// 2 = Read / Write
// 3 = Read Only
