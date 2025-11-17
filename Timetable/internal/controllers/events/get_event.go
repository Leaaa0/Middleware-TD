package events

import (
	"encoding/json"
	"middleware/timetable/internal/helpers"
	"middleware/timetable/internal/services/events"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetEvent
// @Tags         events
// @Summary      Get an event.
// @Description  Get an event.
// @Param        id           	path      string  true  "event UUID formatted ID"
// @Success      200            {object}  models.events
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /events/{id} [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, _ := ctx.Value("eventId").(uuid.UUID) // getting key set in context.go

	event, err := events.GetEventById(eventId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(event)
	_, _ = w.Write(body)
	return
}
