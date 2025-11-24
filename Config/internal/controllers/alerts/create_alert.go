package alerts

import (
	"encoding/json"
	"middleware/config/internal/models"
	"middleware/config/internal/services/alerts"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type AlertAdding struct {
	Resource     *uuid.UUID `json:"resource"`
	AllResources bool       `json:"allResources"`
	Mail         string     `json:"mail"`
}

type CreateAlertResponse struct {
	Message string       `json:"message"`
	Alert   models.Alert `json:"alert"`
}

// CreatAlert
// @Tags         alert
// @Summary      Create an alert.
// @Description  Create an alert
// @Success      200            {object}  models.Alert
// @Failure      500            "Something went wrong"
// @Router       /alerts [post]
func CreatAlert(w http.ResponseWriter, r *http.Request) {
	var alertAdding AlertAdding

	err := json.NewDecoder(r.Body).Decode(&alertAdding)
	if err != nil {
		logrus.Error("Error while decoding JSON body : ", err)
		http.Error(w, "JSON data incorrects", http.StatusBadRequest)
		return
	}
	if alertAdding.Mail == "" { // ucaId doit être composé de 6 chiffres
		http.Error(w, "Contact mail required", http.StatusBadRequest)
		return
	}
	if alertAdding.AllResources == false && alertAdding.Resource.String() == "" {
		http.Error(w, "Resource id required if AllResources is false", http.StatusBadRequest)
		return
	}

	alertCreated, err := alerts.CreateAlert(alertAdding.Resource.String(), alertAdding.AllResources, alertAdding.Mail)
	if err != nil {
		http.Error(w, "Error while adding new alert", http.StatusBadRequest)
		return
	}

	bodyResponse := CreateAlertResponse{
		Message: "Alert added successfully",
		Alert:   *alertCreated,
	}
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(bodyResponse)
	_, _ = w.Write(body)
	return
}
