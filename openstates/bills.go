package openstates

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// SearchBills returns all bills with parameters 'v'
// Valid Parameters:
// * state					- Filter by state
// * chamber				- Filter by chamber
// * bill_id				- Only bills with id bill_id
// * q							- Bills matching provided full text query
// * search_window	- By default all bills are searched, but you
//											can limit the search window with these:
//		* all							- Default, include all sessions
//		* term						- Only bills from the current term
//		* session					- Only bills from the current session
//		* session:2009		- Only bills from the session named 2009
//		* term:2009-2011	- Only bills from the term named 2009-2011
// * updated_since	- Only bills updated since the provided date
// * sponsor_id			- Bills sponsored by a given legislator id
// * subject				- Only bills categorized by OpenStates as
//											belonging to this subject
// * type						- Only bills of a given type
//										('bill', 'resolution', etc.)
// Additionally, 'sort' is a valid parameter, defaults to 'last'
// Other sort options:
// * first				- Oldest first
// * last					- Newest first (default)
// * signed				- Signed first
// * passed_lower - Passed Lower first
// * passed_upper - Passed Upper first
// * updated_at		- Sort by updated_at time
// * created_at		- Sort by created_at time
//
// The API will not return exceedingly large responses, so it may
// be necessary to use 'page' and 'per_page' to control the # of
// of results returned:
// page			- Page of results, each of size per_page (defaults to 1)
// per_page	- Number of results per page, is unlimited unless page is
//							set, in which case it defaults to 50.
func (o *OpenStates) SearchBills(v url.Values) ([]Bill, error) {
	var ret []Bill
	var err error
	var getVal []byte
	if getVal, err = o.call("bills", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		for i := range ret {
			if err = o.fixBillTimes(&ret[i]); err != nil {
				return ret, err
			}
		}
	}
	return ret, err
}

// GetBillsForState gets all bills for state st
func (o *OpenStates) GetBillsForState(st string) ([]Bill, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return []Bill{}, err
	}
	vals := url.Values{}
	vals.Set("state", st)
	return o.SearchBills(vals)
}

// QueryBillsForState Does a search for bills with text 'q' in state 'st
func (o *OpenStates) QueryBillsForState(st string, q string) ([]Bill, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return []Bill{}, err
	}
	vals := url.Values{}
	vals.Set("state", st)
	vals.Set("q", q)
	return o.SearchBills(vals)
}

func (o *OpenStates) getBillDetailForEndpoint(ep string) (*Bill, error) {
	var ret *Bill
	var err error
	var getVal []byte
	v := url.Values{}
	if getVal, err = o.call(ep, v); err != nil {
		return ret, err
	}
	fmt.Println(string(getVal))
	err = json.Unmarshal(getVal, &ret)
	ret.HasDetail = true
	if err == nil {
		err = o.fixBillTimes(ret)
	}
	return ret, err
}

// GetBillDetail Gets bill detail from state, session, and name
func (o *OpenStates) GetBillDetail(st, sess, name string) (*Bill, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return nil, err
	}
	return o.getBillDetailForEndpoint("bills/" + st + "/" + sess + "/" + name)
}

// GetBillDetailFromID ...
func (o *OpenStates) GetBillDetailFromID(id string) (*Bill, error) {
	return o.getBillDetailForEndpoint("bills/" + id)
}

func (o *OpenStates) fixBillTimes(b *Bill) error {
	var err error
	if err = UnmarshalTimeString(b.CreatedAtStr, &b.CreatedAt); err != nil {
		return errors.New("No Created At Time")
	}
	if err = UnmarshalTimeString(b.UpdatedAtStr, &b.UpdatedAt); err != nil {
		return errors.New("No Updated At Time")
	}
	UnmarshalTimeString(b.ActionDates.PassedUpperStr, &b.ActionDates.PassedUpper)
	UnmarshalTimeString(b.ActionDates.PassedLowerStr, &b.ActionDates.PassedLower)
	UnmarshalTimeString(b.ActionDates.LastStr, &b.ActionDates.Last)
	UnmarshalTimeString(b.ActionDates.SignedStr, &b.ActionDates.Signed)
	UnmarshalTimeString(b.ActionDates.FirstStr, &b.ActionDates.First)
	for i := range b.Actions {
		UnmarshalTimeString(b.Actions[i].DateStr, &b.Actions[i].Date)
	}
	for i := range b.Votes {
		UnmarshalTimeString(b.Votes[i].DateStr, &b.Votes[i].Date)
	}
	return nil
}
