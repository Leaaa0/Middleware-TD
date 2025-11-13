package calendars

import (
	"encoding/json"
	"middleware/config/internal/models"
	"middleware/config/internal/services/calendars"
	"net/http"

	"github.com/sirupsen/logrus"
)

type CalendarAdding struct {
	UcaId int    `json:"ucaId"`
	Name  string `json:"name"`
}

type CreateCalendarResponse struct {
	Message  string          `json:"message"`
	Calendar models.Calendar `json:"calendar"`
}

// CreateCalendar
// @Tags         calendar
// @Summary      Create a calendar.
// @Description  Create a calendar
// @Success      200            {object}  models.Calendar
// @Failure      500            "Something went wrong"
// @Router       /calendars [post]
func CreateCalendar(w http.ResponseWriter, r *http.Request) {
	var calendarAdding CalendarAdding

	err := json.NewDecoder(r.Body).Decode(&calendarAdding)
	if err != nil {
		logrus.Error("Error while decoding JSON body : ", err)
		http.Error(w, "JSON data incorrects", http.StatusBadRequest)
		return
	}
	if calendarAdding.Name == "" || calendarAdding.UcaId < 10000 { // ucaId doit être composé de 6 chiffres
		http.Error(w, " Name and ucaID required", http.StatusBadRequest)
		return
	}

	calendarCreated, err := calendars.CreateCalendar(calendarAdding.UcaId, calendarAdding.Name)
	if err != nil {
		http.Error(w, "Error while adding new calendar", http.StatusBadRequest)
		return
	}

	bodyResponse := CreateCalendarResponse{
		Message:  "Calendar added successfully",
		Calendar: *calendarCreated,
	}
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(bodyResponse)
	_, _ = w.Write(body)
	return
}
