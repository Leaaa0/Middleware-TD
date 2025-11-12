package users

import (
	"encoding/json"
	"net/http"
)

// DeleteUser
// @Tags         users
// @Summary      Delete a user.
// @Description  Delete a user
// @Success      200            {object}  models.User
// @Failure      500            "Something went wrong"
// @Router       /users [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Test suppression")
	_, _ = w.Write(body)
	return
}
