package openstates

import "time"

// Committee is a committee
type Committee struct {
	Level        string            `json:"level"`
	CreatedAtStr string            `json:"created_at"`
	CreatedAt    time.Time         `json:"-"`
	UpdatedAtStr string            `json:"updated_at"`
	UpdatedAt    time.Time         `json:"-"`
	ParentID     string            `json:"parent_id"`
	State        string            `json:"state"`
	Subcommittee string            `json:"subcommittee"`
	Name         string            `json:"committee"`
	Chamber      string            `json:"chamber"`
	ID           string            `json:"id"`
	HasDetail    bool              `json:"-"`
	Members      []CommitteeMember `json:"members"`
	Sources      []Source          `json:"sources"`
}

// CommitteeMember is a member of a committee
type CommitteeMember struct {
	LegID string `json:"leg_id"`
	Role  string `json:"role"`
	Name  string `json:"name"`
}
