package main

import (
	"log"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"strconv"
)

type Fields struct{
	n string
	t string
}

const (
	IDENTIFIER string = "id INT PRIMARY KEY"
	VARCHAR string = "VARCHAR(255)"
	TEXT string = "TEXT"
	TIMESTAMP string = "TIMESTAMP"
	BOOL string = "SMALLINT(1)"
)

const (
	EventTable string = "event"
)


var event_fields = []Fields{
	{"subject", VARCHAR},
	{"Location", TEXT},
	{"starttime", TIMESTAMP},
	{"endtime", TIMESTAMP},
	{"whole_day", BOOL},
	{"note", TEXT},
	{"inviter", VARCHAR},
	{"invitees", VARCHAR},
	{"TESTCOLUMN", VARCHAR},
	{"TESTCOLUMN2", BOOL},
}

func Example() {
	database, _ := sql.Open("sqlite3", "./janko.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS " + EventTable + " (id INTEGER PRIMARY_KEY, subject TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO event (subject) VALUES (?)")
	statement.Exec("Test Event")
	rows, _ := database.Query("SELECT id, subject FROM event")
	var id int
	var subject string
	for rows.Next() {
		rows.Scan(&id, &subject)
		fmt.Println(strconv.Itoa(id) + ": " + subject)
	}
}

func Connect() *sql.DB {
	database, _ := sql.Open("sqlite3", "./janko.db")
	return database
}

func CreateEvent(n string) {
	c := Connect()
	q,_ := c.Prepare("INSERT INTO event (subject) VALUES (?)")
	q.Exec(n)
}

func GetEventList() {
	c := Connect()
	rows, _ := c.Query("SELECT * FROM event")
	var id int
	var subject string
	for rows.Next() {
		rows.Scan(&id, &subject)
		fmt.Println(strconv.Itoa(id) + ": " + subject)
	}
}

func IsNewColumn(column string) bool {
	c := Connect()
	q := fmt.Sprintf("SELECT %s FROM %s", column, EventTable)
	_, err := c.Query(q)
	if err != nil {
		return true
	}
	return false
}

func install() {
	c := Connect()
	q := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)",  EventTable, IDENTIFIER)
	stmt, err := c.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	for _, val := range event_fields {
		isNew := IsNewColumn(val.n)
		if isNew {
			q := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", EventTable, val.n, val.t)
			stmt, _ := c.Prepare(q)
			stmt.Exec()
		}
	}
}

func main() {
	install()
	GetEventList()
}