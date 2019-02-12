package model_manager

import (
    _ "github.com/jinzhu/gorm"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/base/utils"
    "github.com/mikasoftware/mikapost-go/model"
)

/* The structure of our manager. */

type ThingManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var thingManager *ThingManager


/* The mangaer functions */

func ThingManagerInstance() (*ThingManager) {
    if thingManager != nil {
        return thingManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        thingManager = &ThingManager{dao}
        return thingManager
    }
}

// Function returns the `Thing` model data based on the inputted `id` parameter.
func (manager *ThingManager) GetByID(id uint64) (*model.Thing, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Thing // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&org).Count(&count) // Find our user.
    return &org, count
}

// Function returns `true` or `false` depending on whether a `UserBoxPermission`
// object exists for the inputted parameters.
func (manager *ThingManager) HasBoxPermission(user *model.User, boxID uint64) bool {
    orm := manager.dao.GetORM() // Get our database layer.
    var count uint64 = 0
    var permission model.UserBoxPermission
    orm.Where("user_id = ? AND box_id = ?", user.ID, boxID).First(&permission).Count(&count)
    return count > 0
}

// Function returns `true` or `false` depending on whether a
// `UserThingPermission` object exists for the inputted parameters.
func (manager *ThingManager) HasThingPermission(user *model.User, thingID uint64) bool {
    orm := manager.dao.GetORM() // Get our database layer.
    var count uint64 = 0
    var permission model.UserThingPermission
    orm.Where("user_id = ? AND thing_id = ?", user.ID, thingID).First(&permission).Count(&count)
    return count > 0
}

func (manager *ThingManager) FilterByUser(user *model.User, pageIndex uint64, paginateBy uint64) ([]model.Thing, uint64, uint64) {
    orm := manager.dao.GetORM()
    var things []model.Thing
    var permissions []model.UserThingPermission
    var thingIDs []int64

    // Lookup all the permissions associated with the user
    orm.Model(&user).Association("UserThingPermissions").Find(&permissions)

    // Pluck all the `ThingID` values from our permissions, these ID values are
    // IDs to `Things` which the user has permission for.
    orm.Find(&permissions).Pluck("thing_id", &thingIDs)

    // Lookup all the things that match the ID values found in the array of ID
    // values. Also apply our filtering to the ORM so the pagination code
    // can properly paginate the results.
    orm = orm.Where("id in (?)", thingIDs).Find(&things)

    // Make our paginated query.
    pagination := utils.Pagging(&utils.Param{
		DB:      orm,
		Page:    pageIndex,
		Limit:   paginateBy,
		OrderBy: []string{"id asc"},
		ShowSQL: false,
	}, &things)

    return things, pagination.TotalPage, pagination.TotalRecord
}
