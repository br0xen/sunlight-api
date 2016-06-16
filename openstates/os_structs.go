package openstates

import "time"

// OpenStates is the Wrapper for the OpenStates API
type OpenStates struct {
	APIVersion int
	APIKey     string
}

// StateMeta is all of the metadata for a state
type StateMeta struct {
	Name               string                   `json:"name"`
	Abbr               string                   `json:"abbreviation"`
	Chambers           map[string]StateChamber  `json:"chambers"`
	FeatureFlags       []string                 `json:"feature_flags"`
	CapitolTimezone    time.Location            `json:"-"`
	CapitolTimezoneStr string                   `json:"capitol_timezone"`
	ID                 string                   `json:"id"`
	LatestCSVDate      time.Time                `json:"-"`
	LatestCSVDateStr   string                   `json:"latest_csv_date"`
	LatestCSVURL       string                   `json:"latest_csv_url"`
	LatestJSONDate     time.Time                `json:"-"`
	LatestJSONDateStr  string                   `json:"latest_json_date"`
	LatestJSONURL      string                   `json:"latest_json_url"`
	LatestUpdate       time.Time                `json:"-"`
	LatestUpdateStr    string                   `json:"latest_update"`
	LegislatureName    string                   `json:"legislature_name"`
	LegislatureURL     string                   `json:"legislature_url"`
	SessionDetails     map[string]SessionDetail `json:"session_details"`
	Terms              []Term                   `json:"terms"`
}

// StateChamber is a Chamber in the state's government
type StateChamber struct {
	Name  string `json:"name"`  // The Chamber name
	Title string `json:"title"` // Title of a person in this chamber
}

// SessionDetail are details about a specific session
type SessionDetail struct {
	Type         string    `json:"type"`
	DisplayName  string    `json:"display_name"`
	StartDate    time.Time `json:"-"`
	StartDateStr string    `json:"start_date"`
}

// Term is a state term
type Term struct {
	EndYear   int      `json:"end_year"`
	StartYear int      `json:"start_year"`
	Name      string   `json:"name"`
	Sessions  []string `json:"sessions"`
}
