package main

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	body := map[string]string{
		"status": "UP",
	}

	resp, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
