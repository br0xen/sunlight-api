package openstates

import "time"

// Event is an event
type Event struct {
	Documents    []Document    `json:"documents"`
	End          time.Time     `json:"-"`
	EndStr       string        `json:"end"`
	Description  string        `json:"description"`
	State        string        `json:"state"`
	Agenda       string        `json:"+agenda"`
	CreatedAtStr string        `json:"created_at"`
	CreatedAt    time.Time     `json:"-"`
	When         time.Time     `json:"-"`
	WhenStr      string        `json:"when"`
	UpdatedAtStr string        `json:"updated_at"`
	UpdatedAt    time.Time     `json:"-"`
	Sources      []Source      `json:"sources"`
	Participants []Participant `json:"participants"`
	Session      string        `json:"session"`
	Location     string        `json:"location"`
	RelatedBills []EventBill   `json:"related_bills"`
	Timezone     string        `json:"timezone"`
	Type         string        `json:"type"`
	ID           string        `json:"id"`
	Chamber      string        `json:"+chamber"`
}

// EventBill is a bill associated with an event
type EventBill struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	ID          string `json:"id"`
	BillID      string `json:"bill_id"`
	Chamber     string `json:"+chamber"`
}

// Document is a document for an event (agenda, etc)
type Document struct {
	URL      string `json:"url"`
	MimeType string `json:"+mimetype"`
	Name     string `json:"name"`
	Type     string `json:"+type"`
}

// Participant is a participant in an event
type Participant struct {
	Chamber         string `json:"chamber"`
	ParticipantType string `json:"participant_type"`
	Name            string `json:"participant"`
	ID              string `json:"id"`
	Type            string `json:"type"`
}
