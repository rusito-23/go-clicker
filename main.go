package main

import (
    // "github.com/gin-gonic/gin"
    "go-friends/db"
    "log"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

    // Get database credentials and specs
    var db, err = db.Initialize()

    if err != nil {
        log.Fatalf("Could not initialize database: %v", err)
        return
    }

    // Close database on exit
    defer db.Close()

    // TODO: Create gin routers
}
