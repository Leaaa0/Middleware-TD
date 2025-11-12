package users

import (
	"encoding/json"
	"net/http"
)

// UpdateUser
// @Tags         users
// @Summary      Update a user.
// @Description  Update a user
// @Success      200            {object}  models.User
// @Failure      500            "Something went wrong"
// @Router       /users [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Test modification")
	_, _ = w.Write(body)
	return
}
