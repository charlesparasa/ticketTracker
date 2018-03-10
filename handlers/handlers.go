package handlers


import (
	"encoding/json"
	"fmt"
	"net/http"
	."bitbucket.org/tekion/ticketTracker/model"
	"bitbucket.org/tekion/ticketTracker/dao"
	."bitbucket.org/tekion/ticketTracker/config"
	"github.com/gorilla/sessions"
	"log"
)

var config = Config{}

var store = sessions.NewCookieStore([]byte("ticket-session"))
func createTicket(w http.ResponseWriter, req *http.Request) {

	var ticketDao dao.TicketDAO
	var input Ticket
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		fmt.Errorf("invalid input params")
		return
	}
	session, err := store.Get(req, input.TicketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

		alarmTime := input.TicketAssignedTime + input.TicketEstimatedTime
		if alarmTime > 24.0 {
			log.Fatalln("Invalid Time")
		} else {
			input.AlaramTriggeredTime = alarmTime
		}
		fmt.Println("input" ,input )
		ticketDao.Connect()
		err = ticketDao.Insert(input)
		if err != nil {
			fmt.Println("Ticket Not Inserted" ,err)
			respondWithJson(w, http.StatusInternalServerError, err)
			return
		}

	// Set some session values.
	session.Values["AlarmTriggerTime"] = input.AlaramTriggeredTime
	session.Values["TicketID"] = input.TicketID
	session.Values["TicketStatus"] = input.TikcetStatus
	// Save it before we write to the response/return from the handler.
	session.Save(req, w)
		respondWithJson(w, http.StatusCreated, input)

}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}