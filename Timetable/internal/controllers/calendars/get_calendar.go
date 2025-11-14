package calendars

import (
	"encoding/json"
	"middleware/timetable/internal/helpers"
	"middleware/timetable/internal/services/calendars"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetCalendar
// @Tags         calendars
// @Summary      Get a calendar.
// @Description  Get a calendar.
// @Param        id           	path      string  true  "calendar UUID formatted ID"
// @Success      200            {object}  models.calendar
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /calendars/{id} [get]
func GetCalendar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	calendarId, _ := ctx.Value("calendarId").(uuid.UUID) // getting key set in context.go

	calendar, err := calendars.GetCalendarById(calendarId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(calendar)
	_, _ = w.Write(body)
	return
}
