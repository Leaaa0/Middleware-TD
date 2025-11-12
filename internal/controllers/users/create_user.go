package users

import (
	"encoding/json"
	"net/http"
)

// CreateUser
// @Tags         users
// @Summary      Create a user.
// @Description  Create a user
// @Success      200            {object}  models.User
// @Failure      500            "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal("Test création")
	_, _ = w.Write(body)
	return
}
