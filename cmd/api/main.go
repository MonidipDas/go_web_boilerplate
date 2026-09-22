package main

import (
	"log"
	"net/http"
	"time"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	//[]byte is used here to typecast a string to slice of bytes.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("200 OK"))
	w.Write([]byte(`{"message":"200 OK"}`))
}

func main() {
	// using mux (multiplexer)
	// http.HandleFunc("GET /healthz", healthz)
	// http.HandleFunc("GET /healthz", healthz)
	// err := http.ListenAndServe(":8080", nil)
	// err := http.ListenAndServe(":8090", mux)
	// if err != nil {
	// 	log.Fatalf("Server Failed : %v", err)
	// }

	// mux := http.NewServeMux()
	// mux.HandleFunc("GET /healthz", healthz)
	// err := http.ListenAndServe(":8090", mux)
	// if err != nil {
	// 	log.Fatalf("Server Failed : %v", err)
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)

	srv := &http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Println("Starting server on :8090...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server Failed : %v", err)
	}
}
