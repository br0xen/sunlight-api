package openstates

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// SearchCommittees retrieves a list of committees based on a number of
// parameters.
// Valid Parameters:
// * committee
// * subcommittee
// * chamber
// * state
func (o *OpenStates) SearchCommittees(v url.Values) ([]Committee, error) {
	var ret []Committee
	var err error
	var getVal []byte
	if getVal, err = o.call("committees/", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		for i := range ret {
			if err = o.fixCommitteeTimes(&ret[i]); err != nil {
				return ret, err
			}
		}
	}
	return ret, err
}

// GetCommitteesForState gets all committees for a state
func (o *OpenStates) GetCommitteesForState(st string) ([]Committee, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return []Committee{}, err
	}
	v := url.Values{}
	v.Set("state", st)
	return o.SearchCommittees(v)
}

// GetCommitteeDetail gets the details for a committee
func (o *OpenStates) GetCommitteeDetail(id string) (*Committee, error) {
	var ret *Committee
	var err error
	var getVal []byte
	v := url.Values{}
	if getVal, err = o.call("committees/"+id+"/", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		if err = o.fixCommitteeTimes(ret); err != nil {
			return ret, err
		}
	}
	return ret, err
}

func (o *OpenStates) fixCommitteeTimes(c *Committee) error {
	if err := UnmarshalTimeString(c.CreatedAtStr, &c.CreatedAt); err != nil {
		return errors.New("No Created At Time")
	}
	if err := UnmarshalTimeString(c.UpdatedAtStr, &c.UpdatedAt); err != nil {
		return errors.New("No Updated At Time")
	}
	return nil
}
