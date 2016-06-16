package openstates

import (
	"encoding/json"
	"net/url"
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
		//for i := range ret {
		//}
	}
	return ret, err
}
