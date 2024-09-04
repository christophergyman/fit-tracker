package main

import (
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)


// handleSubmit handles the form submission
func handleSubmit(c echo.Context) error {
	// Retrieve form data
	workout := c.FormValue("workout")
	datetime := c.FormValue("datetime")
	note := c.FormValue("note")

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO workouts(workout, datetime, notes) VALUES(?, ?, ?)")
	if err != nil {
	    log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(workout, datetime, note)
	if err != nil {
	    log.Fatal(err)
	}

	// Print form data to the console (or you can save it to a database, etc.)
	fmt.Printf("workout: %s\n", workout)
	fmt.Printf("datetime: %s\n", datetime)
	fmt.Printf("note: %s\n", note)

	// Send a response back to the client
	return c.String(http.StatusOK, "Form submitted successfully!")
}

func getWorkouts(c echo.Context) error {
	//connec to db
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
	    log.Fatal(err)
	}
	defer db.Close()

	//get all rows from workouts table
	rows, err := db.Query("SELECT * FROM workouts")
	if err != nil {
	    log.Fatal(err)
	}
	defer rows.Close()

	//append the values into workouts array
	var workouts []string
	for rows.Next() {
	    var id int
	    var workout, datetime, notes string
	    err = rows.Scan(&id, &workout, &datetime, &notes)
	    if err != nil {
		log.Fatal(err)
	    }

	    workouts = append(workouts, fmt.Sprintf("ID: %d, Workout: %s, Datetime: %s, Notes: %s", id, workout, datetime, notes))
	}
	return c.String(http.StatusOK, fmt.Sprintf("Workouts: \n%s", strings.Join(workouts, "\n")))
}

func main() {
	// assinging forms and their webpages
	e := echo.New()
	e.File("/", "public/index.html")
	e.File("/form", "public/form.html")
	e.POST("/submit", handleSubmit)
	e.GET("/getWorkouts", getWorkouts)
	e.Logger.Fatal(e.Start(":1323"))
}
