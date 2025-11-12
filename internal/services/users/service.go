package calendars

import (
	"database/sql"
	"fmt"
	"middleware/config/internal/models"
	repository "middleware/config/internal/repositories/users"

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
