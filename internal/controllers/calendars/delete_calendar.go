package calendars

import (
	"encoding/json"
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/calendars"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteCalendar
// @Tags         calendars
// @Summary      Delete a calendar.
// @Description  Delete a calendar
// @Success      200            {object}  models.User
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Calendar deleted successfully")
	_, _ = w.Write(body)
	return
}
