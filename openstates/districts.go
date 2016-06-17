package openstates

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// SearchDistricts retrieves a list of districts based on a number of
// parameters
// Valid
func (o *OpenStates) SearchDistricts(v url.Values) ([]District, error) {
	var ret []District
	var err error
	if v.Get("state") == "" {
		return ret, errors.New("SearchDistricts: State is required")
	}
	var getVal []byte
	ep := "districts/" + v.Get("state") + "/"
	if v.Get("chamber") != "" {
		ep += v.Get("chamber") + "/"
		v.Del("chamber")
	}
	v.Del("state")
	if getVal, err = o.call(ep, v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	return ret, err
}

// GetDistrictsForState returns all districts for state st
func (o *OpenStates) GetDistrictsForState(st string) ([]District, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return []District{}, err
	}
	vals := url.Values{}
	vals.Set("state", st)
	return o.SearchDistricts(vals)
}

// GetDistrictBoundary returns the boundary object for a district
func (o *OpenStates) GetDistrictBoundary(bid string) (*DistrictBoundary, error) {
	var ret *DistrictBoundary
	var err error
	var getVal []byte
	v := url.Values{}
	if getVal, err = o.call("/districts/boundary/"+bid, v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	return ret, err
}
