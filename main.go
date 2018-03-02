package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	."bitbucket.org/tekion/ticketTracker/controller"
	"fmt"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/tickets", CreateTicketEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
	if err := http.ListenAndServe(":12345", router); err != nil {
		log.Fatal(err)
	}
	fmt.Println("sever Started")
}
