package model_manager

import (
    _ "github.com/jinzhu/gorm"
    
    "github.com/mikasoftware/mikapost-go/base/database"
    "github.com/mikasoftware/mikapost-go/base/service"
    "github.com/mikasoftware/mikapost-go/model"
)

/* The structure of our manager. */

type UserManager struct {
    dao *database.DataAcessObject
}


/* The global variables. */

var userManager *UserManager


/* The mangaer functions */

func UserManagerInstance() (*UserManager) {
    if userManager != nil {
        return userManager
    } else {
        // Get our database connection.
        dao := database.Instance()
        userManager = &UserManager{dao}
        return userManager
    }
}

func (manager *UserManager) GetByEmail(email string) (*model.User, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var user model.User // The model we will be returning.
    var count uint64
    orm.Where("email = ?", email).First(&user).Count(&count) // Find our user.
    return &user, count
}

func (manager *UserManager) GetByID(id uint64) (*model.User, uint64) {
    orm := manager.dao.GetORM() // Get our database layer.
    var user model.User // The model we will be returning.
    var count uint64
    orm.Where("id = ?", id).First(&user).Count(&count) // Find our user.
    return &user, count
}

func (manager *UserManager) Create(email string, password string, firstName string, lastName string) (*model.User, error) {
    // The model we will be creating.
    var user model.User

    // Secure our password so it's stored in an unreadable form.
    hashedPassword, _ := service.HashPassword(password)

    // Create our `User` object in our database.
    user = model.User {
        Email:        email,
        PasswordHash: hashedPassword,
        FirstName:    firstName,
        LastName:     lastName,
    }

    orm := manager.dao.GetORM() // Get our database layer.
    orm.Create(&user) // Create our object in the database.
    return &user, nil
}
