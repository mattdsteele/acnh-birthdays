package acnh

import (
	"fmt"
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
func addBirthday(cal *ics.Calendar, name string, villagers *VillagerInfo) error {
	villager, err := villagers.Villager(name)
	if err != nil {
		return err
	}
	b := villager.Birthday
	days := strings.Split(b, "/")
	newEvent(cal, name, days[1], days[0])
	return nil
}

func AcnhGopher(w http.ResponseWriter, r *http.Request) {
	// verify a villagers param
	v := r.URL.Query().Get("villagers")
	if v == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "text/calendar")
	cal := ics.NewCalendar()
	cal.SetName("\"ACNH Villager Birthdays\"")
	cal.SetDescription("\"ACNH Villager Birthdays\"")
	cal.SetXWRCalName("\"ACNH Villager Birthdays\"")
	cal.SetXWRCalDesc("\"ACNH Villager Birthdays\"")
	cal.SetProductId(newUUID())
	villagersList := villagers()
	for _, v := range strings.Split(v, ",") {
		err := addBirthday(cal, v, &villagersList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	fmt.Fprintf(w, cal.Serialize())
}
