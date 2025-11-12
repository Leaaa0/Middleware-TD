package calendars

import (
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllCalendars() ([]models.Calendar, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM calendars")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	calendars := []models.Calendar{}
	for rows.Next() {
		var data models.Calendar
		err = rows.Scan(&data.Id, &data.Name)
		if err != nil {
			return nil, err
		}
		calendars = append(calendars, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return calendars, err
}

func GetCalendarById(id uuid.UUID) (*models.Calendar, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM calendars WHERE id=?", id.String())
	helpers.CloseDB(db)

	var calendar models.Calendar
	err = row.Scan(&calendar.Id, &calendar.Name)
	if err != nil {
		return nil, err
	}
	return &calendar, err
}
