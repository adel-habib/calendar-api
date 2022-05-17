package persistance

import (
	"calendar-api/model"
	"time"
)

func GetHolidayById(id int) model.Holiday {
	var h model.Holiday
	DB.Where("id = ?", id).First(&h)
	return h
}

func GetHolidaysByRegionName(name string) []model.Holiday {
	var h []model.Holiday
	DB.Preload("Regions where name = ?", name).Find(&h)
	return h
}
func GetHolidaysByRegionNameAndYear(name string, year int) []model.Holiday {
	var h []model.Holiday
	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(year, time.December, 31, 0, 0, 0, 0, time.Local)
	DB.Preload("Regions", "name = ?", name).Where("date BETWEEN ? and ?", start, end).Find(&h)
	return h
}

func SaveHolidays(h []model.Holiday) (int64, error) {
	result := DB.Create(&h)
	return result.RowsAffected, result.Error
}
