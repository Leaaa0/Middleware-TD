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
)

type CalendarAdding struct {
	UcaId int    `json:"ucaId"`
	Name  string `json:"name"`
}

// CreateCalendar
// @Tags         calendar
// @Summary      Create a calendar.
// @Description  Create a calendar
// @Success      201            {object}  models.Calendar
// @Failure 	 400			"Cannot parse body to JSON data"
// @Failure 	 422 			"Incorrect JSON data : Name and ucaID required"
// @Failure      500            "Something went wrong"
// @Router       /calendars [post]
func CreateCalendar(w http.ResponseWriter, r *http.Request) {
	var calendarAdding CalendarAdding

	bodyBytes, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	err = json.NewDecoder(r.Body).Decode(&calendarAdding)
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
	if calendarAdding.Name == "" || calendarAdding.UcaId < 10000 { // ucaId need to be 6 digits number
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

	calendarCreated, err := calendars.CreateCalendar(calendarAdding.UcaId, calendarAdding.Name)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(calendarCreated)
	_, _ = w.Write(body)
	return
}
