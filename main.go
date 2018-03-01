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
	//router.HandleFunc("/people", model.GetPeopleEndpoint).Methods("GET")
	//router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/tickets", CreateTicketEndpoint).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
	if err := http.ListenAndServe(":12345", router); err != nil {
		log.Fatal(err)
	}
	fmt.Println("sever Started")
}
