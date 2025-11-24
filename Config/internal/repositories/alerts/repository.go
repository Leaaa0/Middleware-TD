package alerts

import (
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllAlerts() ([]models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM alerts")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	alerts := []models.Alert{}
	for rows.Next() {
		var data models.Alert
		err = rows.Scan(&data.Id, &data.Resource, &data.AllResources, &data.Mail)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return alerts, err
}

func GetAlertById(id uuid.UUID) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM alerts WHERE id=?", id.String())
	helpers.CloseDB(db)

	var alert models.Alert
	err = row.Scan(&alert.Id, &alert.Resource, &alert.AllResources, &alert.Mail)
	if err != nil {
		return nil, err
	}
	return &alert, err
}

func CreateAlert(resource string, allResources bool, mail string) (*models.Alert, error) {

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	if resource == "" {
		_, err = db.Exec("INSERT INTO alerts (id, resource, allResources, mail) VALUES (?,?,?,?)", id.String(), nil, allResources, mail)
	} else {
		_, err = db.Exec("INSERT INTO alerts (id, resource, allResources, mail) VALUES (?,?,?,?)", id.String(), resource, allResources, mail)
	}
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	resourceUuid, _ := uuid.FromString(resource)
	return &models.Alert{&id, &resourceUuid, allResources, mail}, nil
}

func DeleteAlert(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM alerts WHERE id=?", id.String())
	helpers.CloseDB(db)

	if err != nil {
		return err
	}

	return nil
}

func UpdateAlert(id uuid.UUID, resource string, allResource bool, mail string) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	if resource == "" {
		_, err = db.Exec("UPDATE alerts SET resource=?, allResources=?, mail=? WHERE id=?", nil, allResource, mail, id.String())
	} else {
		_, err = db.Exec("UPDATE alerts SET resource=?, allResources=?, mail=? WHERE id=?", resource, allResource, mail, id.String())
	}
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	resourceUuid, _ := uuid.FromString(resource)
	return &models.Alert{&id, &resourceUuid, allResource, mail}, nil
}
