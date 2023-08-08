package shared

import (
	"log"
	"time"
)

var monthMap = map[time.Month]string{
	time.January:   "Jan",
	time.February:  "Fev",
	time.March:     "Mar",
	time.April:     "Abr",
	time.May:       "Mai",
	time.June:      "Jun",
	time.July:      "Jul",
	time.August:    "Ago",
	time.September: "Set",
	time.October:   "Out",
	time.November:  "Nov",
	time.December:  "Dez",
}

var loc, _ = time.LoadLocation("America/Sao_Paulo")

func FormateDate(date string) string {
	var timestampedDate, err = time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	timestampedDate = timestampedDate.In(loc)
	return timestampedDate.Format("01 " + monthMap[timestampedDate.Month()] + " 2006 - 15:04H")
}
