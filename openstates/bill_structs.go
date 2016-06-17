package openstates

import "time"

// Bill is a Bill
type Bill struct {
	Title           string       `json:"title"`
	CreatedAtStr    string       `json:"created_at"`
	CreatedAt       time.Time    `json:"-"`
	UpdatedAtStr    string       `json:"updated_at"`
	UpdatedAt       time.Time    `json:"-"`
	Chamber         string       `json:"chamber"`
	State           string       `json:"state"`
	Session         string       `json:"session"`
	Subjects        []string     `json:"subjects"`
	Type            []string     `json:"type"`
	ID              string       `json:"id"`
	BillID          string       `json:"bill_id"`
	HasDetail       bool         `json:"-"`
	ActionDates     *ActionDates `json:"action_dates"`
	Actions         []Action     `json:"actions"`
	AlternateTitles []string     `json:"alternate_titles"`
	Documents       []string     `json:"documents"`
	Level           string       `json:"level"`
	ScrapedSubjects []string     `json:"scraped_subjects"`
	Sources         []Source     `json:"sources"`
	Sponsors        []Sponsor    `json:"sponsors"`
	Versions        []Version    `json:"versions"`
	Votes           []Vote       `json:"votes"`
}

// ActionDates are the dates that actions happened on the bill
type ActionDates struct {
	PassedUpperStr string    `json:"passed_upper"`
	PassedUpper    time.Time `json:"-"`
	PassedLowerStr string    `json:"passed_lower"`
	PassedLower    time.Time `json:"-"`
	LastStr        string    `json:"last"`
	Last           time.Time `json:"-"`
	SignedStr      string    `json:"signed"`
	Signed         time.Time `json:"-"`
	FirstStr       string    `json:"first"`
	First          time.Time `json:"-"`
}

// Action on a bill
type Action struct {
	Date    time.Time
	DateStr string   `json:"date"`
	Action  string   `json:"action"`
	Type    []string `json:"type"`
	Actor   string   `json:"actor"`
}

// Sponsor is a sponsor on a bill
type Sponsor struct {
	LegID string `json:"leg_id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
}

// Version is a version of a bill
type Version struct {
	URL      string `json:"url"`
	MimeType string `json:"mimetype"`
	DocID    string `json:"doc_id"`
	Name     string `json:"name"`
}

// Vote is a vote on a bill
type Vote struct {
	OtherCount int          `json:"other_count"`
	Threshold  string       `json:"+threshold"`
	OtherVotes []Legislator `json:"other_votes"`
	YesCount   int          `json:"yes_count"`
	YesVotes   []Legislator `json:"yes_votes"`
	NoCount    int          `json:"no_count"`
	NoVotes    []Legislator `json:"no_votes"`
	Motion     string       `json:"motion"`
	Chamber    string       `json:"chamber"`
	State      string       `json:"state"`
	Session    string       `json:"session"`
	Sources    []Source     `json:"sources"`
	Passed     bool         `json:"passed"`
	DateStr    string       `json:"date"`
	Date       time.Time    `json:"-"`
	VoteID     string       `json:"vote_id"`
	Type       string       `json:"type"`
	ID         string       `json:"id"`
	BillID     string       `json:"bill_id"`
}
