package main

import (
	"middleware/config/internal/controllers/calendars"
	"middleware/config/internal/helpers"
	_ "middleware/config/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) { // route /users
		r.Get("/", calendars.GetCalendar)     // GET /users - Récupérer tous les users
		r.Post("/", calendars.CreateCalendar) // POST /users - Créer un nouveau user

		r.Route("/{id}", func(r chi.Router) { // route /users/{id}
			r.Use(calendars.Context)                // Use Context method to get user ID
			r.Get("/", calendars.GetCalendar)       // GET /users/{id} - Récupérer un user
			r.Put("/", calendars.UpdateCalendar)    // PUT /users/{id} - Mettre à jour un user
			r.Delete("/", calendars.DeleteCalendar) // DELETE /users/{id} - Supprimer un user
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
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
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
