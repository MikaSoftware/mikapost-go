package model_manager

import (
    _ "github.com/jinzhu/gorm"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/base/utils"
    "github.com/mikasoftware/mikapost-go/model"
)

/* The structure of our manager. */

type BoxManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var boxManager *BoxManager


/* The mangaer functions */

func BoxManagerInstance() (*BoxManager) {
    if boxManager != nil {
        return boxManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        boxManager = &BoxManager{dao}
        return boxManager
    }
}

func (manager *BoxManager) GetByID(id uint64) (*model.Box, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.Box // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&org).Count(&count) // Find our user.
    return &org, count
}

// Function returns a `true` or `false` value depending on whether the `Box`
// object has permission to be access bey the user account.
func (manager *BoxManager) HasPermission(user *model.User, boxID uint64) bool {
    orm := manager.dao.GetORM() // Get our database layer.
    var count uint64
    var permission model.UserBoxPermission
    orm.Where("user_id = ? AND box_id = ?", user.ID, boxID).First(&permission).Count(&count)
    return count > 0
}

func (manager *BoxManager) FilterBy(user *model.User, pageIndex uint64, paginateBy uint64) ([]model.Box, uint64, uint64) {
    orm := manager.dao.GetORM()
    var boxes []model.Box
    var permissions []model.UserBoxPermission
    var boxIDs []int64

    // Lookup all the permissions associated with the user
    orm.Model(&user).Association("UserBoxPermissions").Find(&permissions)

    // Pluck all the `BoxID` values from our permissions.
    orm.Find(&permissions).Pluck("box_id", &boxIDs)

    // Lookup all the boxes that match the inputted BoxID values found. Also
    // apply our filtering to the ORM so the pagination code can properly
    // paginate the results.
    orm = orm.Where("id in (?)", boxIDs).Find(&boxes)

    // Make our paginated query.
    pagination := utils.Pagging(&utils.Param{
		DB:      orm,
		Page:    pageIndex,
		Limit:   paginateBy,
		OrderBy: []string{"id asc"},
		ShowSQL: false,
	}, &boxes)

    return boxes, pagination.TotalPage, pagination.TotalRecord
}
