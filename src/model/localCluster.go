package model

import (
	"gorm.io/gorm"
	"time"
)

type ClusterType int

const (
	Normal     ClusterType = 1
	MultiWorld ClusterType = 2
)

type LocalCluster struct {
	gorm.Model
	Name        string
	Description string
	AppId       int
	Type        ClusterType
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
