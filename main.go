package main

import (
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	//[]byte is used here to typecast a string to slice of bytes.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("200 OK"))
	w.Write([]byte(`{"message":"200 OK"}`))
}

func main() {
	http.HandleFunc("GET/healthz", healthz)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server Failed : %v", err)
	}

}
