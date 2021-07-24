package db

import (
    "go-friends/models"
)

// FindUserByUsername - Search for a particular user
func (db DB) FindUserByUsername(username string) models.User {
    var user models.User
    db.conn.First(&user, username)
    return user
}

// CreateNewUser - Inserts a new user into the database
func (db DB) CreateNewUser(user models.User) error {
    db.conn.Create(user)
    return nil
}
