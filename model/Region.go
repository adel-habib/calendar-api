package model

import "fmt"

type Region struct {
	Id           uint64  `json:"Id" gorm:"primaryKey;autoIncrement;not null;"`
	Name         string  `json:"Name"`
	ShortName    string  `json:"Abb."`
	ParentId     *uint   `json:"-"`
	ParentRegion *Region `json:"-" gorm:"foreignKey:ParentId"`
}

func (r Region) String() string {
	return fmt.Sprintf("{Id:%d, Name:%s, Abb.:%s}", r.Id, r.Name, r.ShortName)
}
