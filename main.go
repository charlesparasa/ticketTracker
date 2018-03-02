package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	."bitbucket.org/tekion/ticketTracker/controller"
	"fmt"
	config "bitbucket.org/tekion/ticketTracker/config"
	dbUtil "bitbucket.org/tekion/ticketTracker/dao"
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
		fmt.Println(err)
	} else {
		fmt.Println("I am here")
	}

}

var conf config.Config
var daoUtil dbUtil.TicketDAO
// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	fmt.Println("init")
	conf.Read()
	daoUtil.Server = conf.Server
	daoUtil.Database = conf.Database
	daoUtil.Connect()
}