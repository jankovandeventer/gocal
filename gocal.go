package main

import (
	// "fmt"
	"time"
	"net/mail"
	"net/http"
	"log"
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	current_month := time.Now().Month()
	renderTemplate(w, r, current_month)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, current_month time.Month) {
	tmpl, err := template.ParseFiles("templates/month.html")
	err = tmpl.ExecuteTemplate(w, "month.html", current_month)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main () {
	http.HandleFunc("/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}