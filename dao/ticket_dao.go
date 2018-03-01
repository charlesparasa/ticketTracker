package dao

import (
	"log"
	. "bitbucket.org/tekion/ticketTracker/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TicketDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Ticket"
)

// Establish a connection to database
func (m *TicketDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of Ticket
func (m *TicketDAO) FindAll() ([]Ticket, error) {
	var ticket []Ticket
	err := db.C(COLLECTION).Find(bson.M{}).All(&ticket)
	return ticket, err
}

// Find a ticket by its id
func (m *TicketDAO) FindById(id string) (Ticket, error) {
	var ticket Ticket
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&ticket)
	return ticket, err
}

// Insert a ticket into database
func (m *TicketDAO) Insert(ticket Ticket) error {
	err := db.C(COLLECTION).Insert(&ticket)
	return err
}

// Delete an existing ticket
func (m *TicketDAO) Delete(ticket Ticket) error {
	err := db.C(COLLECTION).Remove(&ticket)
	return err
}

// Update an existing ticket
func (m *TicketDAO) Update(ticket Ticket) error {
	err := db.C(COLLECTION).UpdateId(ticket.TicketID, &ticket)
	return err
}

