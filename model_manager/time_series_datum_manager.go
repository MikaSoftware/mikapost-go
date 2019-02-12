package model_manager

import (
    _ "github.com/jinzhu/gorm"

    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/base/utils"
    "github.com/mikasoftware/mikapost-go/model"
)

/* The structure of our manager. */

type TimeSeriesDatumManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var time_series_dataManager *TimeSeriesDatumManager


/* The mangaer functions */

func TimeSeriesDatumManagerInstance() (*TimeSeriesDatumManager) {
    if time_series_dataManager != nil {
        return time_series_dataManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        time_series_dataManager = &TimeSeriesDatumManager{dao}
        return time_series_dataManager
    }
}

func (manager *TimeSeriesDatumManager) GetByID(id uint64) (*model.TimeSeriesDatum, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var org model.TimeSeriesDatum // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&org).Count(&count) // Find our user.
    return &org, count
}


func (manager *TimeSeriesDatumManager) HasPermission(userID uint64, thingID uint64) bool {
    // Setup our variables.
    orm := manager.dao.GetORM() // Get our database layer.
    var count uint64 = 0

    //--------------------------------------------------------------------------
    // CASE 1 OF 2: User has permission for `Thing` object.
    //--------------------------------------------------------------------------
    var thingPermission model.UserThingPermission
    orm.Where("user_id = ? AND thing_id = ?", userID, thingID).First(&thingPermission).Count(&count)
    if count > 0 {
        return true
    }

    //--------------------------------------------------------------------------
    // CASE 2 OF 2: User has permission for `Box` object which belongs to the
    //              `Thing` object.
    //--------------------------------------------------------------------------
    var thing model.Thing
    orm.Where("id = ?", thingID).First(&thing).Count(&count)
    if count == 0 {
        return false
    }

    var boxPermission model.UserBoxPermission
    orm.Where("box_id = ? and user_id = ?", thing.BoxID, userID).First(&boxPermission).Count(&count)
    return count > 0
}


func (manager *TimeSeriesDatumManager) FilterByThing(thingID uint64, pageIndex uint64, paginateBy uint64) ([]model.TimeSeriesDatum, uint64, uint64) {
    orm := manager.dao.GetORM()
    var timeSeriesData []model.TimeSeriesDatum

    // (C) Find all "Time-Series Data" that belong to the list of "Things".
    orm = orm.Where("thing_id = ?", thingID).Find(&timeSeriesData)

    // Make our paginated query.
    pagination := utils.Pagging(&utils.Param{
		DB:      orm,
		Page:    pageIndex,
		Limit:   paginateBy,
		OrderBy: []string{"id asc"},
		ShowSQL: false,
	}, &timeSeriesData)

    return timeSeriesData, pagination.TotalPage, pagination.TotalRecord
}
