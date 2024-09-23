package routes

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kibzrael/raelelectronics/auth/data"
	"github.com/kibzrael/raelelectronics/auth/utils"
)

type RegisterInput struct {
	Uid      string    `json:"uid"`
	FullName string    `json:"full_name" db:"full_name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	LoggedAt time.Time `json:"logged_at" db:"logged_at"`
}

func RegisterHandler(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	}

	var input RegisterInput
	err = json.Unmarshal(body, &input)
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	} else if input.Email == "" || input.Password == "" || input.FullName == "" {
		err = errors.New("email, password, and full_name arguments are required")
		utils.ApiPanic(&res, &err)
		return
	}

	db := ctx.Value(data.DB_CONTEXT).(*sqlx.DB)

	input.Uid = uuid.New().String()
	input.LoggedAt = time.Now()
	input.Password, err = utils.HashPassword(input.Password)
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	}

	_, err = db.NamedExec("INSERT INTO users (uid, full_name, email, password, logged_at) VALUES(:uid, :full_name, :email, :password, :logged_at)", &input)
	if err != nil {
		utils.ApiPanic(&res, &err)
		return
	}

	input.Password = ""
	response := map[string]any{"message": "registration successful", "user": input}
	res.WriteHeader(http.StatusCreated)
	utils.JsonResponse(&res, response)

}
