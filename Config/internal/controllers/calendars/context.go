package calendars

import (
	"context"
	"fmt"
	"middleware/config/internal/helpers"
	"middleware/config/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
)

// Context
/* This method is used to get ressource ID from url
*
* In REST, urls are formed like this : calendars/{specific_collection_ressource_id}/another_collection/{another_collection_ressource_id}...
* In this example, it could be calendars/{calendar_id} to get specific calendar infos or calendars/{calendar_id}/events/{event_id} to get specific calendar's specific event
 */
func Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calendarId, err := uuid.FromString(chi.URLParam(r, "id"))
		if err != nil {
			body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id"))})

			w.WriteHeader(status)
			if body != nil {
				_, _ = w.Write(body)
			}
			return
		}

		ctx := context.WithValue(r.Context(), "calendarId", calendarId) // We fill context with a Key-valued variable
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
