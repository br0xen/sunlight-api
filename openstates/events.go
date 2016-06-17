package openstates

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// SearchEvents retrieves a list of events based on some parameters
// Valid Parameters
// * state - Filter by state
// * type  - Filter by type
func (o *OpenStates) SearchEvents(v url.Values) ([]Event, error) {
	var ret []Event
	var err error
	var getVal []byte
	if getVal, err = o.call("events", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		for i := range ret {
			if err = o.fixEventTimes(&ret[i]); err != nil {
				return ret, err
			}
		}
	}
	return ret, err
}

// GetEventsForState returns all events for a state
func (o *OpenStates) GetEventsForState(st string) ([]Event, error) {
	st, err := states.ScrubToAbbr(st)
	if err != nil {
		return []Event{}, err
	}
	v := url.Values{}
	v.Set("state", st)
	return o.SearchEvents(v)
}

// GetEventDetail returns event detail for an event id
func (o *OpenStates) GetEventDetail(id string) (*Event, error) {
	var ret *Event
	var err error
	var getVal []byte
	v := url.Values{}
	if getVal, err = o.call("events/"+id+"/", v); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	if err == nil {
		err = o.fixEventTimes(ret)
	}
	return ret, err
}

func (o *OpenStates) fixEventTimes(e *Event) error {
	if err := UnmarshalTimeString(e.CreatedAtStr, &e.CreatedAt); err != nil {
		return errors.New("No Created At Time")
	}
	if err := UnmarshalTimeString(e.UpdatedAtStr, &e.UpdatedAt); err != nil {
		return errors.New("No Updated At Time")
	}
	UnmarshalTimeString(e.EndStr, &e.End)
	UnmarshalTimeString(e.WhenStr, &e.When)
	return nil
}
