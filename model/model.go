package model

type Ticket struct {
	TicketID        		string   	 `bson:"_id" json:"id"`
	TicketAssignedTime  	string    	 `bson:"ticketAssignedTime" json:"ticketAssignedTime"`
	TikcetEstimatedTime  	string   	 `bson:"tikcetEstimatedTime" json:"tikcetEstimatedTime"`
	AlaramTriggeredTime   	string 	 	 `bson:"alaramTriggeredTime" json:"alaramTriggeredTime"`
	TikcetStatus			string 		 `bson:"tikcetStatus" json:"ticketStatus"`
}
const (
	Assigned   = "assigned"
	Inprogress = "inprogress"
	Closed     = "closed"
)








