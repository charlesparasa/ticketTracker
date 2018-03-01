package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	."bitbucket.org/tekion/ticketTracker/model"
	"bitbucket.org/tekion/ticketTracker/dao"
	."bitbucket.org/tekion/ticketTracker/config"
)

var config = Config{}


func CreateTicketEndpoint(w http.ResponseWriter, req *http.Request) {
	var ticketDao dao.TicketDAO
	var input Ticket
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		fmt.Errorf("invalid input params")
		return
	}

	if input.AlaramTriggeredTime == "" || input.TicketAssignedTime == "" || input.TikcetEstimatedTime == "" || input.TikcetStatus == "" {
		fmt.Errorf("invalid input params")
		return
	} else {
		ticketDao.Connect()
		err := ticketDao.Insert(input)
		if err != nil {
			fmt.Println("Ticket Not Inserted")
			return
		}
	}
	json.NewEncoder(w).Encode(input)
	respondWithJson(w, http.StatusCreated, input)

}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}