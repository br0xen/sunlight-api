package openstates

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/br0xen/sunlight-api/openstates/states"
)

// Create returns a new OpenStates object
func Create(key string) *OpenStates {
	o := OpenStates{APIVersion: 1, APIKey: key}
	return &o
}

// SetAPIVersion changes the api version
func (o *OpenStates) SetAPIVersion(i int) {
	o.APIVersion = i
}

func (o *OpenStates) call(endpoint string, values url.Values) ([]byte, error) {
	var ret []byte
	url := "http://openstates.org/api/v" + strconv.Itoa(o.APIVersion) + "/" + endpoint + "?apikey=" + o.APIKey
	if len(values) > 0 {
		url += "&" + values.Encode()
	}
	resp, err := http.Get(url)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

// AllMetadata returns all state's metadata
func (o *OpenStates) AllMetadata() ([]StateMeta, error) {
	var ret []StateMeta
	var getVal []byte
	var err error
	vals := url.Values{}
	if getVal, err = o.call("metadata", vals); err != nil {
		return ret, err
	}
	err = json.Unmarshal(getVal, &ret)
	return ret, err
}

// StateMetadata pulls metadata for a specific state
func (o *OpenStates) StateMetadata(st string) (*StateMeta, error) {
	var ret *StateMeta
	var err error
	orig := st
	if states.IsValidName(st) {
		st, err = states.NameToAbbr(st)
	}
	if !states.IsValidAbbr(st) {
		return ret, errors.New("Invalid State: " + orig)
	}
	var getVal []byte
	vals := url.Values{}
	if getVal, err = o.call("metadata/"+st, vals); err != nil {
		return ret, err
	}
	fmt.Println(string(getVal))
	err = json.Unmarshal(getVal, &ret)
	return ret, err
}

// UnmarshalTimeString Takes a time string and a pointer to a time object
// and populates the time object with the value from the string
func UnmarshalTimeString(s string, t *time.Time) error {
	var err error
	// Check if 's' is empty
	if s == "" {
		return errors.New("Invalid Time Value")
	}
	*t, err = time.Parse("2006-01-02 15:04:05", s)
	return err
}

/*
	resp, err := http.Get("http://openstates.org/api/v1/legislators/geo/?apikey=" + site.OpenStatesKey + "&lat=" + lat + "&long=" + lng)
	if err != nil {
		fmt.Fprint(w, "{\"status\":\"error\"}")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(body))
	// http://openstates.org/api/v1/legislators/geo/?lat=35.79&long=-78.78
	// https://sunlightlabs.github.io/openstates-api/legislators.html#examples/geo-lookup
*/
