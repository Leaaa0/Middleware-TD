package main

import (
	"middleware/timetable/internal/controllers/events"
	"middleware/timetable/internal/helpers"
	_ "middleware/timetable/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/events", func(r chi.Router) { // route /events
		r.Get("/", events.GetEvents) // GET /events - Get all the events

		r.Route("/{id}", func(r chi.Router) { // route /events/{id}
			r.Use(events.Context)       // Use Context method to get event ID
			r.Get("/", events.GetEvent) // GET /events/{id} - Get an event
		})
	})

	logrus.Info("[INFO] Timetable API server started. Now listening on *:8092")
	logrus.Fatalln(http.ListenAndServe(":8092", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS events (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
    		dateStart INT NOT NULL,
			dateEnd VARCHAR(255) NOT NULL,
    		summary VARCHAR(255) NOT NULL,
    		location VARCHAR(255) NOT NULL,
    		description BLOB,
    		lastModified INT NOT NULL
 		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
