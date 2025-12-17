package calendars

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"
	"middleware/config/internal/services/calendars"
	"net/http"

	"github.com/gofrs/uuid"
)

type CalendarModifying struct {
	UcaId int    `json:"UcaId"`
	Name  string `json:"Name"`
}

type UpdateCalendarResponse struct {
	Message  string          `json:"Message"`
	Calendar models.Calendar `json:"Calendar"`
}

// UpdateCalendar
// @Tags         calendars
// @Summary      Update a calendar.
// @Description  Update a calendar
// @Success      200            {object}  models.Alert
// @Failure 	 400 			"Cannot parse body to JSON data"
// @Failure 	 404			"Calendar not found"
// @Failure      422            "Cannot parse id"
// @Failure 	 422 			"Incorrect JSON data : Name and ucaID required"
// @Failure      500            "Something went wrong"
// @Router       /calendars [put]
func UpdateCalendar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	calendarId, _ := ctx.Value("calendarId").(uuid.UUID)

	var calendarModifying CalendarModifying

	bodyBytes, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.NewDecoder(r.Body).Decode(&calendarModifying)
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
	if calendarModifying.Name == "" || calendarModifying.UcaId < 10000 { // ucaId need to be 6 digits number
		err = &models.ErrorUnprocessableEntity{
			Message: "Incorrect JSON data : Name and ucaID required",
		}
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	calendarUpdated, err := calendars.UpdateCalendar(calendarId, calendarModifying.UcaId, calendarModifying.Name)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(calendarUpdated)
	_, _ = w.Write(body)
	return
}
