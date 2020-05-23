package main

import (
	"fmt"
	"io/ioutil"
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
	e.SetDescription(fmt.Sprintf("%s's birthday", name))
	e.SetSummary(fmt.Sprintf("%s's birthday", name))
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
func main() {
	cal := ics.NewCalendar()
	cal.SetName("Foo bar baz")
	cal.SetProductId("Oh yeah!")
	villagersList := villagers()
	addBirthday(cal, "Ellie", &villagersList)
	addBirthday(cal, "Spike", &villagersList)
	addBirthday(cal, "Canberra", &villagersList)
	fmt.Println(cal.Serialize())

	// save file
	ioutil.WriteFile("birthdays.ics", []byte(cal.Serialize()), 0644)
}
