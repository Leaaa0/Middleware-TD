package calendars

import (
	"encoding/json"
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/calendars"
	"net/http"
)

// GetCalendars
// @Tags         calendars
// @Summary      Get all calendars.
// @Description  Get all calendars.
// @Success      200            {array}  models.Calendar
// @Failure      500             "Something went wrong"
// @Router       /calendars [get]
func GetCalendars(w http.ResponseWriter, _ *http.Request) {
	// calling service
	calendars, err := calendars.GetAllCalendars()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(calendars)
	_, _ = w.Write(body)
	return
}
