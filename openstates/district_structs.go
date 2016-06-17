package openstates

// District is a district
type District struct {
	Abbr        string               `json:"abbr"`
	BoundaryID  string               `json:"boundary_id"`
	Chamber     string               `json:"chamber"`
	ID          string               `json:"id"`
	Legislators []DistrictLegislator `json:"legislators"`
	Name        string               `json:"name"`
	NumSeats    int                  `json:"num_seats"`
}

// DistrictLegislator is a legislator entry in a district
type DistrictLegislator struct {
	FullName string `json:"full_name"`
	LegID    string `json:"leg_id"`
}

// DistrictBoundary describes the boundary of a district
type DistrictBoundary struct {
	DivisionID string                `json:"division_id"`
	Name       string                `json:"name"`
	Region     *DistrictRegion       `json:"region"`
	Chamber    string                `json:"chamber"`
	Shape      *DistrictShape        `json:"shape"`
	Abbr       string                `json:"abbr"`
	BoundaryID string                `json:"boundary_id"`
	NumSeats   int                   `json:"num_seats"`
	ID         string                `json:"id"`
	BBox       [2]DistrictBoundaryXY `json:"bbox"`
}

// DistrictBoundaryXY is an x/y
type DistrictBoundaryXY [2]float64

// DistrictRegion is a bounding region
type DistrictRegion struct {
	CenterLat float64 `json:"center_lat"`
	CenterLon float64 `json:"center_lon"`
	LatDelta  float64 `json:"lat_delta"`
	LonDelta  float64 `json:"lon_delta"`
}

// DistrictShape is a bounding shape
type DistrictShape [][][]DistrictBoundaryXY
