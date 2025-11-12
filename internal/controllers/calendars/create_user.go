package calendars

import (
	"encoding/json"
	"net/http"
)

// CreateCalendar
// @Tags         calendar
// @Summary      Create a calendar.
// @Description  Create a calendar
// @Success      200            {object}  models.User
// @Failure      500            "Something went wrong"
// @Router       /calendars [post]
func CreateCalendar(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Test création")
	_, _ = w.Write(body)
	return
}
