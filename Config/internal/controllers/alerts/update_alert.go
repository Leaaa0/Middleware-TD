package alerts

import (
	"encoding/json"
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"
	"middleware/config/internal/services/alerts"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type AlertModifying struct {
	Resource     string `json:"resource"`
	AllResources bool   `json:"allResources"`
	Mail         string `json:"mail"`
}

type UpdateAlertResponse struct {
	Message string       `json:"message"`
	Alert   models.Alert `json:"alert"`
}

// UpdateAlert
// @Tags         alert
// @Summary      Update an alert.
// @Description  Update an alert
// @Success      200            {object}  models.Alert
// @Failure      400            "Bad request"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /alerts [put]
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("alertId").(uuid.UUID)

	var alertModifying AlertModifying

	err := json.NewDecoder(r.Body).Decode(&alertModifying)
	if err != nil {
		logrus.Error("Error while decoding JSON body : ", err)
		http.Error(w, "JSON data incorrects", http.StatusBadRequest)
		return
	}
	if alertModifying.Mail == "" {
		http.Error(w, "Contact mail required", http.StatusBadRequest)
		return
	}
	if alertModifying.AllResources == false {
		if alertModifying.Resource == "" {
			http.Error(w, "At least one ressource must be watched", http.StatusBadRequest)
			return
		}

		_, err = uuid.FromString(alertModifying.Resource)
		if err != nil {
			http.Error(w, "Resource ID incorrect : must be an UUID", http.StatusBadRequest)
			return
		}
	}

	alertUpdated, err := alerts.UpdateAlert(alertId, alertModifying.Resource, alertModifying.AllResources, alertModifying.Mail)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	bodyResponse := UpdateAlertResponse{
		Message: "Alert updated successfully",
		Alert:   *alertUpdated,
	}
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(bodyResponse)
	_, _ = w.Write(body)
	return
}
