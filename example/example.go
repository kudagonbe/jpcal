package main

import (
	"log"

	"github.com/kudagonbe/jpcal"
)

func main() {
	const year int = 2020
	const month int = 1

	var ds jpcal.Days

	// Get all days in a year.
	ds, _ = jpcal.AllDays(year)
	log.Println("====================", "AllDays", "====================")
	for _, v := range ds {
		log.Println(v)
	}

	// Get all days in a month.
	ds, _ = jpcal.AllDaysYM(year, month)
	log.Println("====================", "AllDaysYM", "====================")
	for _, v := range ds {
		log.Println(v)
	}

	// Get all national holidays in a year.
	ds, _ = jpcal.Holidays(year)
	log.Println("====================", "Holidays", "====================")
	for _, v := range ds {
		log.Println(v)
	}

	// Get all national holidays in a month.
	ds, _ = jpcal.HolidaysYM(year, month)
	log.Println("====================", "HolidaysYM", "====================")
	for _, v := range ds {
		log.Println(v)
	}

	// Get specific type days in a year.
	ds, _ = jpcal.SpecificTypeDays(year, jpcal.TypeSaturday, jpcal.TypeSunday)
	log.Println("====================", "SpecificTypeDays", "====================")
	for _, v := range ds {
		log.Println(v)
	}

	// Get specific type days in a month.
	ds, _ = jpcal.SpecificTypeDaysYM(year, month, jpcal.TypeWeekDay, jpcal.TypeNationalHoliday)
	log.Println("====================", "SpecificTypeDaysYM", "====================")
	for _, v := range ds {
		log.Println(v)
	}
}
