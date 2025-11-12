package calendars

import (
	"encoding/json"
	"net/http"
)

// UpdateCalendar
// @Tags         calendars
// @Summary      Update a calendar.
// @Description  Update a calendar
// @Success      200            {object}  models.Calendar
// @Failure      500            "Something went wrong"
// @Router       /calendars [put]
func UpdateCalendar(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Test modification")
	_, _ = w.Write(body)
	return
}
