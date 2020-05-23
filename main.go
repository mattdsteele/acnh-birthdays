package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
)

func annualEvent(month, day string) string {
	return "FREQ=YEARLY;BYMONTH=" + month + ";BYMONTHDAY=" + day
}
func newUUID() string {
	u := uuid.New()
	text, _ := u.MarshalBinary()
	id := fmt.Sprintf("%x@%s", text, "acnh-calendar")
	return id

}
func newEvent(cal *ics.Calendar, name, monthS, dayS string) {
	e := cal.AddEvent(newUUID())
	e.SetDescription(fmt.Sprintf("%s's birthday", strings.Title(name)))
	e.SetSummary(fmt.Sprintf("%s's birthday", strings.Title(name)))
	today, _ := time.Parse("2006-1-2", "2020-"+monthS+"-"+dayS)
	e.SetAllDayEndAt(today.AddDate(0, 0, 1))
	e.SetAllDayStartAt(today)
	e.SetDtStampTime(today)
	e.AddProperty(ics.ComponentProperty("RRULE"), annualEvent(monthS, dayS))
}
func addBirthday(cal *ics.Calendar, name string, villagers *VillagerInfo) {
	ellie, _ := villagers.Villager(name)
	b := ellie.Birthday
	days := strings.Split(b, "/")
	newEvent(cal, name, days[1], days[0])
}
func mainx() {
	cal := ics.NewCalendar()
	cal.SetName("Foo bar baz")
	cal.SetProductId("Oh yeah!")
	villagersList := villagers()
	addBirthday(cal, "ellie", &villagersList)
	addBirthday(cal, "spike", &villagersList)
	addBirthday(cal, "canberra", &villagersList)
	fmt.Println(cal.Serialize())

	// save file
	ioutil.WriteFile("birthdays.ics", []byte(cal.Serialize()), 0644)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/calendar")
		cal := ics.NewCalendar()
		cal.SetName("Foo bar baz")
		cal.SetProductId("Oh yeah!")
		villagersList := villagers()
		addBirthday(cal, "ellie", &villagersList)
		addBirthday(cal, "spike", &villagersList)
		addBirthday(cal, "canberra", &villagersList)
		fmt.Fprintf(w, cal.Serialize())
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
