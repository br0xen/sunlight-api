package openstates

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// GetBillsWithParams returns all bills with parameters 'v'
func (o *OpenStates) GetBillsWithParams(v url.Values) ([]Bill, error) {
	var ret []Bill
	var err error
	var getVal []byte
	if getVal, err = o.call("bills", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		for i := range ret {
			ret[i].CreatedAt, err = time.Parse("2006-01-02 15:04:05", ret[i].CreatedAtStr)
			if err != nil {
				fmt.Println("Error on idx: " + strconv.Itoa(i))
				return ret, err
			}
			ret[i].UpdatedAt, err = time.Parse("2006-01-02 15:04:05", ret[i].UpdatedAtStr)
			if err != nil {
				fmt.Println("Error on idx: " + strconv.Itoa(i))
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
	return o.GetBillsWithParams(vals)
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
	return o.GetBillsWithParams(vals)
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
		if err = UnmarshalTimeString(ret.CreatedAtStr, &ret.CreatedAt); err != nil {
			return ret, err
		}
		if err = UnmarshalTimeString(ret.UpdatedAtStr, &ret.UpdatedAt); err != nil {
			return ret, err
		}
		UnmarshalTimeString(ret.ActionDates.PassedUpperStr, &ret.ActionDates.PassedUpper)
		UnmarshalTimeString(ret.ActionDates.PassedLowerStr, &ret.ActionDates.PassedLower)
		UnmarshalTimeString(ret.ActionDates.LastStr, &ret.ActionDates.Last)
		UnmarshalTimeString(ret.ActionDates.SignedStr, &ret.ActionDates.Signed)
		UnmarshalTimeString(ret.ActionDates.FirstStr, &ret.ActionDates.First)
		for i := range ret.Actions {
			UnmarshalTimeString(ret.Actions[i].DateStr, &ret.Actions[i].Date)
		}
		for i := range ret.Votes {
			UnmarshalTimeString(ret.Votes[i].DateStr, &ret.Votes[i].Date)
		}
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
