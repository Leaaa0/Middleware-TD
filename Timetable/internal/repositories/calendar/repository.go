package calendars

import (
	"middleware/timetable/internal/helpers"
	"middleware/timetable/internal/models"

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
		err = rows.Scan(&data.Id, &data.UcaId, &data.Name)
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
	err = row.Scan(&calendar.Id, &calendar.UcaId, &calendar.Name)
	if err != nil {
		return nil, err
	}
	return &calendar, err
}

func CreateCalendar(ucaId int, name string) (*models.Calendar, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO calendars (id, ucaId, name) VALUES (?,?,?)", id.String(), ucaId, name)
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	return &models.Calendar{&id, ucaId, name}, nil
}

func DeleteCalendar(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM calendars WHERE id=?", id.String())
	helpers.CloseDB(db)

	if err != nil {
		return err
	}

	return nil
}

func UpdateCalendar(id uuid.UUID, ucaId int, name string) (*models.Calendar, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("UPDATE calendars SET ucaId=?, name=? WHERE id=?", ucaId, name, id.String())
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	return &models.Calendar{&id, ucaId, name}, nil
}
