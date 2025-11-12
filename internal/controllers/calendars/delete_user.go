package calendars

import (
	"encoding/json"
	"net/http"
)

// DeleteCalendar
// @Tags         calendars
// @Summary      Delete a calendar.
// @Description  Delete a calendar
// @Success      200            {object}  models.User
// @Failure      500            "Something went wrong"
// @Router       /calendars [delete]
func DeleteCalendar(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Test suppression")
	_, _ = w.Write(body)
	return
}
