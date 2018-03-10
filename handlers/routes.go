package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/tickets", createTicket).Methods("POST")
	/*r.HandleFunc("/movies", updateTicket).Methods("POST")
	r.HandleFunc("/movies", deleteTicket).Methods("DELETE")
	r.HandleFunc("/movies", readTicket).Methods("GET")
	r.HandleFunc("/movies/{id}", findTicketByID).Methods("GET")*/
	if err := http.ListenAndServe(":12345", r); err != nil {
		log.Fatal(err)
	}
}

