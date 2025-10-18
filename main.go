package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {

	server := http.NewServeMux()

	fmt.Println("Hello World!!!")

	server.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		body := map[string]string{
			"status": "UP",
		}

		resp, err := json.Marshal(body)

		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(resp)

	})

	err := http.ListenAndServe(":4000", server)

	if err != nil {
		os.Exit(1)
	}
}
