package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) listSettings(w http.ResponseWriter, r *http.Request) {

	app.logger.Debug("GET /settings")

	settings := map[string]any{
		"toggleOne":   true,
		"toggleTwo":   false,
		"toggleThree": true,
	}

	body := map[string]any{
		"settings": settings,
		"version":  "1.0.0",
	}

	resp, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
