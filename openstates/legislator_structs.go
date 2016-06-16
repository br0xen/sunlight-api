package openstates

import "time"

// Legislator is a legislator
type Legislator struct {
	FirstName     string            `json:"first_name"`
	LastName      string            `json:"last_name"`
	MiddleName    string            `json:"middle_name"`
	District      string            `json:"district"`
	Chamber       string            `json:"chamber"`
	URL           string            `json:"url"`
	CreatedAtStr  string            `json:"created_at"`
	CreatedAt     time.Time         `json:"-"`
	UpdatedAtStr  string            `json:"updated_at"`
	UpdatedAt     time.Time         `json:"-"`
	Email         string            `json:"email"`
	Active        bool              `json:"active"`
	State         string            `json:"state"`
	Offices       []Office          `json:"offices"`
	OfficeAddress string            `json:"office_address"`
	VotesmartID   string            `json:"votesmart_id"`
	FullName      string            `json:"full_name"`
	LegID         string            `json:"leg_id"`
	Party         string            `json:"party"`
	Suffixes      string            `json:"suffixes"`
	ID            string            `json:"id"`
	PhotoURL      string            `json:"photo_url"`
	Fax           string            `json:"+fax"`
	Level         string            `json:"level"`
	Phone         string            `json:"+phone"`
	OldRoles      map[string][]Role `json:"old_roles"`
	Roles         []Role            `json:"roles"`

	NameOnBill string `json:"name"` // Shows up when pulling bill detail
}

// Office is a legislators office
type Office struct {
	Fax     string `json:"fax"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Type    string `json:"type"`
	Email   string `json:"email"`
}

// Role is a legislators role
type Role struct {
	Term         string    `json:"term"`
	StartDateStr string    `json:"start_date"`
	StartDate    time.Time `json:"-"`
	EndDateStr   string    `json:"end_date"`
	EndDate      time.Time `json:"-"`
	District     string    `json:"district"`
	Chamber      string    `json:"chamber"`
	State        string    `json:"state"`
	Party        string    `json:"party"`
	Type         string    `json:"type"`
	CommitteeID  string    `json:"committee_id"`
	SubCommittee string    `json:"subcommittee"`
	Committee    string    `json:"committee"`
	Position     string    `json:"position"`
}
