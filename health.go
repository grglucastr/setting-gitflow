package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {

	app.logger.Debug("GET /health")

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
