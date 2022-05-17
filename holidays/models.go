package holidays

import (
	"fmt"
	"time"
)

type Region struct {
	Id           uint64  `json:"Id" gorm:"primaryKey;autoIncrement;not null;"`
	Name         string  `json:"Name"`
	ShortName    string  `json:"Abb."`
	ParentId     *uint   `json:"-"`
	ParentRegion *Region `json:"-" gorm:"foreignKey:ParentId"`
}

type Holiday struct {
	Id      uint64    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Regions []Region  `json:"-" gorm:"many2many:holiday_regions;"`
}

func (r Region) String() string {
	return fmt.Sprintf("{Id:%d, Name:%s, Abb.:%s}", r.Id, r.Name, r.ShortName)
}
