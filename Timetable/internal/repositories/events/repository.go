package events

import (
	"middleware/timetable/internal/helpers"
	"middleware/timetable/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllEvents() ([]models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM events")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	events := []models.Event{}
	for rows.Next() {
		var data models.Event
		err = rows.Scan(&data.Id, &data.UcaId, &data.Name)
		if err != nil {
			return nil, err
		}
		events = append(events, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return events, err
}

func GetEventById(id uuid.UUID) (*models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM events WHERE id=?", id.String())
	helpers.CloseDB(db)

	var event models.Event
	err = row.Scan(&event.Id, &event.UcaId, &event.Name)
	if err != nil {
		return nil, err
	}
	return &event, err
}
