package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	status := true
	if data == nil {
		status = false
	}
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(struct {
		Status bool `json:"status"`
		Data   interface{} `json:"data"`
	}{
		Status: status,
		Data:   data,
	})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
	responseLogger(data)
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

func responseLogger(d interface{}) {
	log.Printf(`response: %+v`, d)
}
