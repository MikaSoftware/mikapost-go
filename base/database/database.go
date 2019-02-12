package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"

    "github.com/mikasoftware/mikapost-go/base/config"
    "github.com/mikasoftware/mikapost-go/model"
)

/* Database Structure */

type DataAcessObject struct {
    dbPool *gorm.DB
}

/* Global variable. */

var dao *DataAcessObject

/* Private initializer */

func init() {
    Instance()
}

/* Function declaration */

// Function will return an instance of our database access layer (DAO) or the
// function will lazily load the DAO and then return the DAO.
func Instance() (*DataAcessObject) {
    // Lazily load the database connection if it was not created before.
    if dao != nil {
        return dao
    }

    // Get the database configuration text from the environment variables.
    databaseConfigString := config.GetSettingsVariableDatabaseURL()

    // The following code will connect our application to the `postgres` database.
    db, err := gorm.Open("postgres", databaseConfigString)
    if err != nil {
        fmt.Println(err)
        panic("Failed to connect database")
    }
    // defer db.Close() // Handle this in `main.go` so do not uncomment this!

    // PLEASE READ FOR MORE INFORAMTION:
    // http://doc.gorm.io/

    // // Automatically delete previous database schema.
    // db.Debug().DropTableIfExists(&model.TimeSeriesDatum{})
    // db.Debug().DropTableIfExists(&model.UserThingPermission{})
    // db.Debug().DropTableIfExists(&model.Thing{})
    // db.Debug().DropTableIfExists(&model.UserBoxPermission{})
    // db.Debug().DropTableIfExists(&model.Box{})
    // db.Debug().DropTableIfExists(&model.User{})

    // Automatically migrate our database schema.
    db.Debug().AutoMigrate(&model.User{})
    db.Debug().AutoMigrate(&model.Box{})
    db.Debug().AutoMigrate(&model.UserBoxPermission{})
    db.Debug().AutoMigrate(&model.Thing{})
    db.Debug().AutoMigrate(&model.UserThingPermission{})
    db.Debug().AutoMigrate(&model.TimeSeriesDatum{})

    // Keep an instance of our new object.
    dao = &DataAcessObject{
        dbPool: db,
    }

    //Return our database connector.
    return dao
}

func (instance *DataAcessObject) DropAndCreateDatabase() {
    // Automatically delete previous database schema.
    instance.dbPool.Debug().DropTableIfExists(&model.TimeSeriesDatum{})
    instance.dbPool.Debug().DropTableIfExists(&model.Thing{})
    instance.dbPool.Debug().DropTableIfExists(&model.UserThingPermission{})
    instance.dbPool.Debug().DropTableIfExists(&model.UserBoxPermission{})
    instance.dbPool.Debug().DropTableIfExists(&model.Box{})
    instance.dbPool.Debug().DropTableIfExists(&model.User{})

    // Automatically migrate our database schema.
    instance.dbPool.Debug().AutoMigrate(&model.User{})
    instance.dbPool.Debug().AutoMigrate(&model.Box{})
    instance.dbPool.Debug().AutoMigrate(&model.UserBoxPermission{})
    instance.dbPool.Debug().AutoMigrate(&model.Thing{})
    instance.dbPool.Debug().AutoMigrate(&model.UserThingPermission{})
    instance.dbPool.Debug().AutoMigrate(&model.TimeSeriesDatum{})
}

func (instance *DataAcessObject) GetORM() (*gorm.DB) {
    return instance.dbPool
}
