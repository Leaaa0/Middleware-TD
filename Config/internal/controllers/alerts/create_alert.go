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

type AlertAdding struct {
	Resource     string `json:"resource"`
	AllResources bool   `json:"allResources"`
	Mail         string `json:"mail"`
}

// CreateAlert
// @Tags         alert
// @Summary      Create an alert.
// @Description  Create an alert
// @Success      201            {object}  models.Alert
// @Failure		 400
// @Failure      422
// @Failure      500            "Something went wrong"
// @Router       /alerts [post]
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var alertAdding AlertAdding

	bodyBytes, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.NewDecoder(r.Body).Decode(&alertAdding)
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
	if alertAdding.Mail == "" {
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
	if alertAdding.AllResources == false {
		if alertAdding.Resource == "" {
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

		_, err = uuid.FromString(alertAdding.Resource)
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

	alertCreated, err := alerts.CreateAlert(alertAdding.Resource, alertAdding.AllResources, alertAdding.Mail)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(alertCreated)
	_, _ = w.Write(body)
	return
}
