package alerts

import (
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/alerts"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteAlert
// @Tags         alert
// @Summary      Delete an alert.
// @Description  Delete an alert
// @Success      204
// @Failure 	 404			"Calendar not found
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /alerts [delete]
func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alertId, _ := ctx.Value("alertId").(uuid.UUID)

	err := alerts.DeleteAlert(alertId)
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
