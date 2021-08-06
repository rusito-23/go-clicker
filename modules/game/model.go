package game

import (
	"github.com/jinzhu/gorm"
)

// Model -
// The game structure that will be saved in the DB
type Model struct {
	gorm.Model

	// The external unique identifier
	ExternalID string `gorm:"unique_index"`

	// The amount of clicks made by the player
	// Starts with 0
	ClickScore Score

	// The current status
	// Starts with `Created`
	Status Status
}

// Score -
type Score uint

// Status -
// Represents the possible status of a game
type Status string

const (
	// Created - The game was created but has not started yet
	Created = "created"
	// Started - The game did start and is in progress
	Started = "started"
	// Finished - The game finished successfully
	Finished = "finished"
)

// TableName -
// Specifies the table name for the Game Model
func (Model) TableName() string {
	return "games"
}
