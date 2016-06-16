package openstates

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// SearchLegislators retrieves a list of legislators based on a number of
// parameters, the results do not include the `roles` or `old_roles` items
// by default
// Valid Parameters:
// * state			- Filter by state
// * first_name - Filter by first name
// * last_name	- Filter by last name
// * chamber		- Only legislators with a role in the specified chamber
// * active			- 'true' (default) to only include current legislators
// * term				- Only legislators that have a role in a certain term
// * district		- Only legislators that have represented the specified district
// * party			- Only legislators that have been associated with a specified
//									party
func (o *OpenStates) SearchLegislators(v url.Values) ([]Legislator, error) {
	var ret []Legislator
	var err error
	var getVal []byte
	if getVal, err = o.call("legislators", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		for i := range ret {
			o.fixLegislatorTimes(&ret[i])
		}
	}
	return ret, err
}

// GetLegislatorsForState returns all legislators for a specific state
func (o *OpenStates) GetLegislatorsForState(st string) ([]Legislator, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return []Legislator{}, err
	}
	v := url.Values{}
	v.Set("state", st)
	return o.SearchLegislators(v)
}

// GetLegislatorDetail returns the detail for a legislator
func (o *OpenStates) GetLegislatorDetail(id string) (*Legislator, error) {
	var ret *Legislator
	var err error
	var getVal []byte
	v := url.Values{}
	if getVal, err = o.call("legislators/"+id, v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	o.fixLegislatorTimes(ret)
	return ret, err
}

// GetLegislatorsForGeo return all legislators for a latitude & longitude
func (o *OpenStates) GetLegislatorsForGeo(lat, lng float64) ([]Legislator, error) {
	var ret []Legislator
	var err error
	var getVal []byte
	v := url.Values{}
	v.Set("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	v.Set("long", strconv.FormatFloat(lng, 'f', -1, 64))
	if getVal, err = o.call("legislators/geo/", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		for i := range ret {
			o.fixLegislatorTimes(&ret[i])
		}
	}
	return ret, err
}

func (o *OpenStates) fixLegislatorTimes(l *Legislator) error {
	if err := UnmarshalTimeString(l.CreatedAtStr, &l.CreatedAt); err != nil {
		return errors.New("No Created At Time")
	}
	if err := UnmarshalTimeString(l.UpdatedAtStr, &l.UpdatedAt); err != nil {
		return errors.New("No Updated At Time")
	}
	for i := range l.Roles {
		UnmarshalTimeString(l.Roles[i].EndDateStr, &l.Roles[i].EndDate)
		UnmarshalTimeString(l.Roles[i].StartDateStr, &l.Roles[i].StartDate)
	}
	for k := range l.OldRoles {
		for i := range l.OldRoles[k] {
			UnmarshalTimeString(l.OldRoles[k][i].EndDateStr, &l.OldRoles[k][i].EndDate)
			UnmarshalTimeString(l.OldRoles[k][i].StartDateStr, &l.OldRoles[k][i].StartDate)
		}
	}
	return nil
}
