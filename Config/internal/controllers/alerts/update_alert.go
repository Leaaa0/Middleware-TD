package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"
	"middleware/config/internal/services/alerts"
	"net/http"

	"github.com/gofrs/uuid"
)

type AlertModifying struct {
	Resource     string `json:"resource"`
	AllResources bool   `json:"allResources"`
	Mail         string `json:"mail"`
}

// UpdateAlert
// @Tags         alert
// @Summary      Update an alert.
// @Description  Update an alert
// @Success      200            {object}  models.Alert
// @Failure 	 400 			"Cannot parse body to JSON data"
// @Failure 	 404			"Alert not found"
// @Failure      422            "Cannot parse id"
// @Failure 	 422 			"Incorrect JSON data : email required"
// @Failure      500            "Something went wrong"
// @Router       /alerts [put]
func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("alertId").(uuid.UUID)

	var alertModifying AlertModifying

	bodyBytes, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.NewDecoder(r.Body).Decode(&alertModifying)
	if err != nil {
		err = &models.ErrorBadRequest{
			Message: fmt.Sprintf("Cannot parse data as JSON data. Body received : ", string(bodyBytes)),
		}
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}
	if alertModifying.Mail == "" {
		err = &models.ErrorUnprocessableEntity{
			Message: "Incorrect JSON data : Mail required",
		}
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}
	if alertModifying.AllResources == false {
		if alertModifying.Resource == "" {
			err = &models.ErrorUnprocessableEntity{
				Message: "At least one ressource must be watched",
			}
			body, status := helpers.RespondError(err)
			w.WriteHeader(status)
			if body != nil {
				_, _ = w.Write(body)
			}
			return
		}

		_, err = uuid.FromString(alertModifying.Resource)
		if err != nil {
			err = &models.ErrorUnprocessableEntity{
				Message: "Resource ID incorrect : must be an UUID",
			}
			body, status := helpers.RespondError(err)
			w.WriteHeader(status)
			if body != nil {
				_, _ = w.Write(body)
			}
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alertUpdated)
	_, _ = w.Write(body)
	return
}
