package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/auth/data"
	"github.com/kibzrael/raelelectronics/auth/utils"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	}

	var input LoginInput
	if err = json.Unmarshal(body, &input); err != nil {
		utils.ApiPanic(&res, &err)
		return
	}
	if input.Email == "" || input.Password == "" {
		err = errors.New("email and password arguments are required")
		utils.ApiPanic(&res, &err)
		return
	}

	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	var user data.User
	err = db.Get(&user, "SELECT * FROM users WHERE email=$1", input.Email)

	switch err {
	case nil:
		if !utils.VerifyPassword(user.Password, input.Password) {
			res.WriteHeader(http.StatusBadRequest)
			response := map[string]any{"message": "invalid password for the given account", "status": http.StatusBadRequest}
			utils.JsonResponse(&res, response)
			return
		}
	case sql.ErrNoRows:
		res.WriteHeader(http.StatusBadRequest)
		response := map[string]any{"message": "invalid password for the given account", "status": http.StatusBadRequest}
		utils.JsonResponse(&res, response)
		return
	default:
		utils.ApiPanic(&res, &err)
	}

	loggedAt := time.Now()
	db.Exec("UPDATE users SET logged_at=$1 WHERE uid=$2", loggedAt, user.Uid)
	user.LoggedAt = &loggedAt

	user.Password = nil
	response := map[string]any{"message": "login successful", "user": user}
	utils.JsonResponse(&res, response)
}
