package main

import (
	"log"
	"time"
	"net/mail"
	"net/http"
	"html/template"
)

type Entry struct {
	Subject string
	Location string
	StartTime time.Time
	EndTime time.Time
	WholeDay  bool
	Note string
	Inviter mail.Address
	Invitees []mail.Address
}

type Data struct {
	Current time.Time
	Days []time.Time
	Entries []Entry
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	data :=Data {
		Current: time.Now(),
		Days: daysInMonth(),
	}
	renderTemplate(w, r, data)
}

func daysIn(month time.Month, year int) int {
	return time.Date(year, month+1, 0,0,0,0,0, time.UTC).Day()
}

func daysInMonth() []time.Time {
	count := daysIn(time.Now().Month(), time.Now().Year())
	firstDay := time.Date(time.Now().Year(), time.Now().Month(), 1, 0,0,0,0,time.UTC)
	start := 1;
	start = start+1 - int(firstDay.Weekday())
	var days []time.Time
	for i:=start; i <= count; i++ {
		days = append(days, time.Date(time.Now().Year(), time.Now().Month(), i, 0,0,0,0,time.UTC))
	}
	return days
}

func renderTemplate(w http.ResponseWriter, r *http.Request, data Data) {
	tmpl, err := template.ParseFiles("templates/month.html")
	err = tmpl.ExecuteTemplate(w, "month.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main () {
	http.HandleFunc("/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}