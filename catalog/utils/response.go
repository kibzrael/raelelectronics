package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonResponse(res *http.ResponseWriter, response any) {
	(*res).Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(response)
	if err != nil {
		ApiPanic(res, &err)
		return
	}
	(*res).Write(data)
}

func ApiPanic(res *http.ResponseWriter, err *error) {
	(*res).WriteHeader(http.StatusInternalServerError)

	response := map[string]any{"message": "Error", "error": (*err).Error(), "status": http.StatusInternalServerError}
	data, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		fmt.Fprintln(*res, "Error Occured:", *err)
	}
	(*res).Header().Set("Content-Type", "application/json")
	(*res).Write(data)
}
