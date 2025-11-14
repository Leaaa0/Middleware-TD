package calendars

import (
	"database/sql"
	"fmt"
	"middleware/timetable/internal/models"
	repository "middleware/timetable/internal/repositories/calendar"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllCalendars() ([]models.Calendar, error) {
	var err error
	// calling repository
	calendars, err := repository.GetAllCalendars()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving calendars : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving calendars",
		}
	}

	return calendars, nil
}

func GetCalendarById(id uuid.UUID) (*models.Calendar, error) {
	calendar, err := repository.GetCalendarById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "calendar not found",
			}
		}
		logrus.Errorf("error retrieving calendar %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving calendar %s", id.String()),
		}
	}

	return calendar, err
}

func CreateCalendar(ucaId int, name string) (*models.Calendar, error) {
	var err error
	// calling repository
	calendarCreated, err := repository.CreateCalendar(ucaId, name)
	// managing errors
	if err != nil {
		logrus.Errorf("error adding new calendar : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while adding new calendar",
		}
	}

	return calendarCreated, nil
}

func DeleteCalendar(id uuid.UUID) error {
	err := repository.DeleteCalendar(id)
	if err != nil {
		logrus.Errorf("error deleting calendar %s : %s", id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting calendar %s", id.String()),
		}
	}

	return err
}

func UpdateCalendar(id uuid.UUID, ucaId int, name string) (*models.Calendar, error) {
	calendarUpdated, err := repository.UpdateCalendar(id, ucaId, name)
	if err != nil {
		logrus.Errorf("error deleting calendar %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting calendar %s", id.String()),
		}
	}

	return calendarUpdated, err
}
