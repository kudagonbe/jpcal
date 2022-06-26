package main

import (
	"log"

	"github.com/kudagonbe/jpcal"
)

func main() {
	ds, err := jpcal.SpecificTypeDays(2019, jpcal.TypeNationalHoliday)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ds {
		log.Println(v)
	}
	log.Printf("len: %d", ds.Len())
}
