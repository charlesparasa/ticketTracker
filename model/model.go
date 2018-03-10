package model



type Ticket struct {
	TicketID        		string   	 	 `bson:"_id" json:"id"`
	TicketAssignedTime  	float32    	 	 `bson:"ticketAssignedTime" json:"ticketAssignedTime"`
	TicketEstimatedTime  	float32   	 	 `bson:"ticketEstimatedTime" json:"ticketEstimatedTime"`
	AlaramTriggeredTime   	float32 	 	 `bson:"alaramTriggeredTime" json:"alaramTriggeredTime"`
	TicketStatusDescription string		     `bson:"ticketStatusDescription" json:"ticketStatusDescription"`
	TikcetStatus			string 		     `bson:"tikcetStatus" json:"ticketStatus"`

}
const (
	Assigned   = "assigned"
	Inprogress = "inprogress"
	Closed     = "closed"
)








