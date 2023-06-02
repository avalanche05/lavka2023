package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Courier struct {
	gorm.Model
	CourierType  string
	Regions      pq.Int32Array  `gorm:"type:integer[]"`
	WorkingHours pq.StringArray `gorm:"type:text[]"`
}

type Order struct {
	gorm.Model
	Weight        float32
	Regions       int32
	DeliveryHours pq.StringArray `gorm:"type:text[]"`
	Cost          int32
	CompletedTime *time.Time
	CourierId     int64
}
