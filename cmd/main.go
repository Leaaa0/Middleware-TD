package main

import (
	"middleware/config/internal/controllers/users"
	"middleware/config/internal/helpers"
	_ "middleware/config/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) { // route /users
		r.Get("/", users.GetUsers)    // GET /users - Récupérer tous les users
		r.Post("/", users.CreateUser) // POST /users - Créer un nouveau user

		r.Route("/{id}", func(r chi.Router) { // route /users/{id}
			r.Use(users.Context)            // Use Context method to get user ID
			r.Get("/", users.GetUser)       // GET /users/{id} - Récupérer un user
			r.Put("/", users.UpdateUser)    // PUT /users/{id} - Mettre à jour un user
			r.Delete("/", users.DeleteUser) // DELETE /users/{id} - Supprimer un user
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
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
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
