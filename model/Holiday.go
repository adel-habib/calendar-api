package model

import (
	"fmt"
	"time"
)

type Holiday struct {
	Id      uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Regions []Region  `json:"-" gorm:"many2many:holiday_regions;"`
}

func (h Holiday) String() string {
	return fmt.Sprintf("{Id:%d, Name:%s, Abb.:%s, Regions: %s} \n", h.Id, h.Name, h.Date, h.Regions)
}
