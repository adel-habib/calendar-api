package service

import (
	"calendar-api/model"
	"calendar-api/persistance"
	"fmt"
	"time"
)

func easterDate(year int) time.Time {

	a := year % 19
	b := year % 4
	c := year % 7
	var k = year / 100
	var p = k / 3
	var q = k / 4

	m := (15 + k - p - q) % 30
	d := (19*a + m) % 30
	n := (4 + k - q) % 7
	e := (2*b + 4*c + 6*d + n) % 7
	easter := 22 + d + e
	easterDate := time.Date(year, time.March, easter, 0, 0, 0, 0, time.Local)
	return easterDate
}

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func CalculateGermanHolidays(year int) {
	var holidays []model.Holiday
	var DERegion model.Region
	var federalStates []model.Region
	persistance.DB.Where("name = ?", model.DE).First(&DERegion)
	persistance.DB.Where("parent_id = ?", DERegion.Id).Find(&federalStates)
	federalStates = append(federalStates, DERegion)

	DAY := time.Hour * 24

	// Fixed Holidays
	NEUJAHR := Date(year, time.January, 1)
	// HEILIGE_DER_DREI_KOENIGE := Date(year, time.January, 6)
	// FRAUEN_TAG := Date(year, time.March, 8)
	TagDerArbeit := Date(year, time.May, 1)
	// MARIA_HIMMELFAHRT := Date(year, time.August, 15)
	// WELT_KINDER_TAG := Date(year, time.September, 20)
	TagDerDeutschenEinheit := Date(year, time.October, 3)
	// REFORMATIONSTAG := Date(year, time.October, 31)
	// ALLERHEILIGEN := Date(year, time.October, 1)
	ErsterWeihnachtstag := Date(year, time.December, 25)
	ZweiterWeihnachtstag := ErsterWeihnachtstag.Add(1 * DAY)

	// Easter-based Holidays
	OsterSonntag := easterDate(year)
	OsterMontag := OsterSonntag.Add(DAY)
	KarFreitag := OsterSonntag.Add(-2 * DAY)
	ChristiHimmelFahrt := OsterSonntag.Add(39 * DAY)
	PfingstSonntag := OsterSonntag.Add(49 * DAY)
	PfingstMontag := PfingstSonntag.Add(1 * DAY)
	// FRONLEICHNAM := OsterSonntag.Add(60 * DAY)
	// National (federal) model, as in 2022
	federalHolidays := [9]string{
		"Neujahr",
		"Karfreitag",
		"Ostermontag",
		"Christi Himmelfahrt",
		"Pfingstmontag",
		"1. Mai",
		" Tag der Deutschen Einheit",
		"erster Weihnachtstag",
		"zweiter Weihnachtstag",
	}
	federalHolidaysDates := [9]time.Time{
		NEUJAHR,
		KarFreitag,
		OsterMontag,
		ChristiHimmelFahrt,
		PfingstMontag, TagDerArbeit,
		TagDerDeutschenEinheit,
		ErsterWeihnachtstag, ZweiterWeihnachtstag,
	}
	for index, date := range federalHolidaysDates {
		holidays = append(holidays, model.Holiday{Date: date, Name: federalHolidays[index], Regions: federalStates})
	}
	rows, err := persistance.SaveHolidays(holidays)
	if err != nil {
		fmt.Println(err)
	}
	println(rows)
}
