package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

type Message struct {
	Msg string
}

func main() {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		m := Message{}
		err := d.Decode(&m)
		if err != nil {
			json.NewEncoder(w).Encode(errors.New("Unable to decode request body"))
			return
		}
		log.Printf("Received message: %v\n", m.Msg)

		json.NewEncoder(w).Encode(m)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check called")
		w.Write([]byte("OK"))
	})
	r.Get("/healthfail", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Health check failed")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	})

	srv := &http.Server{
		Addr:    "0.0.0.0:7777",
		Handler: r,
	}

	go func() {
		log.Println("Listening on http://localhost:7777")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Setup handler for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	<-c

	log.Println("Shutting down")
}
