package alerts

import (
	"encoding/json"
	"middleware/config/internal/helpers"
	"middleware/config/internal/services/alerts"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteAlert
// @Tags         alert
// @Summary      Delete an alert.
// @Description  Delete an alert
// @Success      200            {object}  models.Alert
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

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Alert deleted successfully")
	_, _ = w.Write(body)
	return
}
