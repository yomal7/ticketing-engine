package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello from the ticketing engine\n"))
	})

	mux.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"status": "ok",
		}

		writeJSON(w, http.StatusOK, data)
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
