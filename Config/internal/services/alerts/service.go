package alerts

import (
	"database/sql"
	"fmt"
	"middleware/config/internal/models"
	repository "middleware/config/internal/repositories/alerts"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAlerts() ([]models.Alert, error) {
	var err error
	// calling repository
	alerts, err := repository.GetAllAlerts()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving alerts : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while retrieving alerts",
		}
	}

	return alerts, nil
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	alert, err := repository.GetAlertById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "alert not found",
			}
		}
		logrus.Errorf("error retrieving alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while retrieving alert %s", id.String()),
		}
	}

	return alert, err
}

func CreateAlert(resource string, allResources bool, mail string) (*models.Alert, error) {
	var err error
	// calling repository
	if allResources {
		resource = ""
	}
	alertCreated, err := repository.CreateAlert(resource, allResources, mail)
	// managing errors
	if err != nil {
		logrus.Errorf("error adding new alert : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Something went wrong while adding new alert",
		}
	}

	return alertCreated, nil
}

func DeleteAlert(id uuid.UUID) error {
	err := repository.DeleteAlert(id)
	if err != nil {
		logrus.Errorf("error deleting alert %s : %s", id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting alert %s", id.String()),
		}
	}

	return err
}

func UpdateAlert(id uuid.UUID, resource string, allResource bool, mail string) (*models.Alert, error) {
	if allResource {
		resource = ""
	}
	alertUpdated, err := repository.UpdateAlert(id, resource, allResource, mail)
	if err != nil {
		logrus.Errorf("error deleting alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Something went wrong while deleting alert %s", id.String()),
		}
	}

	return alertUpdated, err
}
