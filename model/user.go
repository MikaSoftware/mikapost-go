package model

import (
    "time"
    _ "github.com/jinzhu/gorm"
)

//a struct to rep user account
type User struct {
    ID                          uint64 `gorm:"primary_key; AUTO_INCREMENT; UNIQUE_INDEX;"`
    Email                       string `gorm:"not null; unique; size:255;"`
    PasswordHash                string `gorm:"size:511`
    FirstName                   string `gorm:"type:varchar(127)”`
    LastName                    string `gorm:"type:varchar(127)”`
    Status                      uint8 `gorm:"DEFAULT: 1;"` // 1 = Active, 2 = Inactive
    GroupID                     uint8 `gorm:"DEFAULT: 1;"` // 1 = Regular, 2 = Admin
    BillingStreetAddress        string `gorm:"size:127;"`
    BillingStreetAddressExtra   string `gorm:"size:127;"`
    BillingCity                 string `gorm:"size:127;"`
    BillingProvince             string `gorm:"size:127;"`
    BillingCountry              string `gorm:"size:127;"`
    BillingPostal               string `gorm:"size:31;"`
    ShippingStreetAddress       string `gorm:"size:127;"`
    ShippingStreetAddressExtra  string `gorm:"size:127;"`
    ShippingCity                string `gorm:"size:127;"`
    ShippingProvince            string `gorm:"size:127;"`
    ShippingCountry             string `gorm:"size:127;"`
    ShippingPostal              string `gorm:"size:31;"`
    CreatedAt                   time.Time
    UpdatedAt                   time.Time
    UserBoxPermissions          []UserBoxPermission // Model
    UserThingPermissions        []UserThingPermission // Model
}

// Give custom table name in our database.
func (u User) TableName() string {
    return "mika_users"
}
