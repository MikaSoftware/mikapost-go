package config

import (
    "os"
    // "fmt"
    "flag"
)

/* global variable declaration */

var databaseConfigURL string
var signingSecretKey string
var address string
var is_running_unit_tests bool

/* function declaration */

func init() {
    if flag.Lookup("test.v") == nil {
        // fmt.Println("App is normal run")
        is_running_unit_tests = false
    } else {
        // fmt.Println("App is run under go test")
        is_running_unit_tests = true
    }
}


func GetSettingsVariableDatabaseURL() string {
    // Get the database configuration string from global variable (if it exists!).
    if databaseConfigURL != "" {
        return databaseConfigURL
    }

    // Get the database configuration text from the environment variables.
    if is_running_unit_tests {
        databaseConfigURL = os.Getenv("TEST_MIKAPOST_GORM_CONFIG")
    } else {
        databaseConfigURL = os.Getenv("MIKAPOST_GORM_CONFIG")
    }

    // Defensive Code: If the programmer forgot to setup our web-applications
    // environment variables then we must error right away.
    if databaseConfigURL == "" {
        if is_running_unit_tests {
            panic("`TEST_MIKAPOST_GORM_CONFIG` environment variable not setup!")
        } else {
            panic("`MIKAPOST_GORM_CONFIG` environment variable not setup!")
        }
    }

    // Return our newly created variable.
    return databaseConfigURL
}


func GetSettingsVariableSigningSecretKey() string {
    // Get the database configuration string from global variable (if it exists!).
    if signingSecretKey != "" {
        return signingSecretKey
    }

    // Get the database configuration text from the environment variables.
    signingSecretKey := os.Getenv("MIKAPOST_SECRET")

    // Defensive Code: If the programmer forgot to setup our web-applications
    // environment variables then we must error right away.
    if signingSecretKey == "" {
        panic("`MIKAPOST_SECRET` environment variable not setup!")
    }

    // Return our newly created variable.
    return signingSecretKey
}


func GetSettingsVariableAddress() string {
    // Get the database configuration string from global variable (if it exists!).
    if address != "" {
        return address
    }

    // Get the database configuration text from the environment variables.
    address := os.Getenv("MIKAPOST_ADDRESS")

    // Defensive Code: If the programmer forgot to setup our web-applications
    // environment variables then we must error right away.
    if address == "" {
        panic("`MIKAPOST_ADDRESS` environment variable not setup!")
    }

    // Return our newly created variable.
    return address
}
