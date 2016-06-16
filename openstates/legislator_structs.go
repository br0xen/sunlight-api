package openstates

// Legislator is a legislator
type Legislator struct {
	LegID      string `json:"leg_id"`
	NameOnBill string `json:"name"` // Shows up when pulling bill detail
}
