package calendars

import (
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/calendars"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteCalendar
// @Tags         calendars
// @Summary      Delete a calendar.
// @Description  Delete a calendar
// @Success      204
// @Failure 	 404			"Calendar not found
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /calendars [delete]
func DeleteCalendar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	calendarId, _ := ctx.Value("calendarId").(uuid.UUID)

	err := calendars.DeleteCalendar(calendarId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
