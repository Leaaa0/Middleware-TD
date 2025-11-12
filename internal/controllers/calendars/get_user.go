package calendars

import (
	"encoding/json"
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/users"
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
	userId, _ := ctx.Value("userId").(uuid.UUID) // getting key set in context.go

	user, err := calendars.GetCalendarById(userId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
	return
}
