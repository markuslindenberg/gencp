package brandmeister

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Repeater is a Brandmeister repeater as fetched from
// https://api.brandmeister.network/v1.0/repeater/?action=LIST
type Repeater struct {
	Agl                 string `json:"agl"`
	Callsign            string `json:"callsign"`
	City                string `json:"city"`
	Colorcode           uint8  `json:"colorcode,string"`
	Description         string `json:"description"`
	Firmware            string `json:"firmware"`
	Gain                string `json:"gain"`
	Hardware            string `json:"hardware"`
	LastKnownMaster     string `json:"lastKnownMaster"`
	LastUpdated         string `json:"last_updated"`
	Lat                 string `json:"lat"`
	Lng                 string `json:"lng"`
	Pep                 string `json:"pep"`
	PriorityDescription string `json:"priorityDescription"`
	Repeaterid          string `json:"repeaterid"`
	Rx                  string `json:"rx"`
	Status              string `json:"status"`
	Sysops              []struct {
		UserRoles []string `json:"user_roles"`
		Username  string   `json:"username"`
	} `json:"sysops,omitempty"`
	Tx      string `json:"tx"`
	Website string `json:"website"`
}

// GetRepeater fetches metadata for given repeater from BM API
func GetRepeater(id string) (*Repeater, error) {
	u, err := url.Parse(fmt.Sprint(apiUri, "/repeater/"))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("action", "GET")
	q.Set("q", id)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var r Repeater
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
