package main

import (
	"middleware/config/internal/controllers/alerts"
	"middleware/config/internal/controllers/calendars"
	"middleware/config/internal/helpers"
	_ "middleware/config/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/calendars", func(r chi.Router) { // route /calendars
		r.Get("/", calendars.GetCalendars)    // GET /calendars - Get all calendars
		r.Post("/", calendars.CreateCalendar) // POST /calendars - Create a new calendar

		r.Route("/{id}", func(r chi.Router) { // route /calendars/{id}
			r.Use(calendars.Context)                // Use Context method to get calendar ID
			r.Get("/", calendars.GetCalendar)       // GET /calendars/{id} - Get a calendar
			r.Put("/", calendars.UpdateCalendar)    // PUT /calendars/{id} - Update a calendar
			r.Delete("/", calendars.DeleteCalendar) // DELETE /calendars/{id} - Delete a calendar
		})
	})
	r.Route("/alerts", func(r chi.Router) { // route /alerts
		r.Get("/", alerts.GetAlerts)    // GET /alerts - Get all alerts
		r.Post("/", alerts.CreateAlert) // POST /alerts - Create a new alert

		r.Route("/{id}", func(r chi.Router) { // route /alerts/{id}
			r.Use(alerts.Context)             // Use Context method to get alert ID
			r.Get("/", alerts.GetAlert)       // GET /alerts/{id} - Get an alert
			r.Put("/", alerts.UpdateAlert)    // PUT /alerts/{id} - Update an alert
			r.Delete("/", alerts.DeleteAlert) // DELETE /alerts/{id} - Delete an alert
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8091")
	logrus.Fatalln(http.ListenAndServe(":8091", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS calendars (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
    		ucaId INT NOT NULL,
			name VARCHAR(255) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS alerts (
		    id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
		    resource VARCHAR(255),
		    allResources BOOLEAN NOT NULL,
		    mail VARCHAR(255)
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
