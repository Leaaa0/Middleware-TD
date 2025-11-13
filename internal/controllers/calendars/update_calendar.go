package calendars

import (
	"encoding/json"
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/calendars"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type CalendarModifying struct {
	UcaId int    `json:"ucaId"`
	Name  string `json:"name"`
}

// UpdateCalendar
// @Tags         calendars
// @Summary      Update a calendar.
// @Description  Update a calendar
// @Success      200            {object}  models.Calendar
// @Failure      400            "Bad request"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /calendars [put]
func UpdateCalendar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	calendarId, _ := ctx.Value("calendarId").(uuid.UUID)

	var calendarModifying CalendarModifying

	err := json.NewDecoder(r.Body).Decode(&calendarModifying)
	if err != nil {
		logrus.Error("Error while decoding JSON body : ", err)
		http.Error(w, "JSON data incorrects", http.StatusBadRequest)
		return
	}
	if calendarModifying.Name == "" || calendarModifying.UcaId < 10000 { // ucaId doit être composé de 6 chiffres
		http.Error(w, " Name and ucaID required", http.StatusBadRequest)
		return
	}

	err = calendars.UpdateCalendar(calendarId, calendarModifying.UcaId, calendarModifying.Name)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Calendar modified successfully")
	_, _ = w.Write(body)
	return
}
