package model

import "gorm.io/gorm"

type WorldState int

const (
	Running WorldState = 1
	Stopped WorldState = 2
)

type World struct {
	gorm.Model
	Name        string
	Description string
	ClusterId   int
	ProcessId   int
	Status      WorldState
}
