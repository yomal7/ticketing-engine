package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello from the ticketing engine\n"))
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status": "ok",
		}

		writeJSON(w, http.StatusOK, data)
	})

	mux.HandleFunc("GET /broken", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"bad": math.NaN()})
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(v)
	if err != nil {
		log.Println("error marshalling response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"internal server error"}`))
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(payload)
}
