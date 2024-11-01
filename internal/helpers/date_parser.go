package helpers

import (
	"log"
	"time"
)

func DateParser(date_str string) time.Time {
	parsedDate, err := time.Parse("2006-01-02", date_str)
	if err != nil {
		log.Println("Error parsing date:", err)
		panic(err)
	}
	return parsedDate
}
