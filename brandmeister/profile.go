package brandmeister

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Profile struct {
	BlockedGroups []interface{} `json:"blockedGroups"`
	Clusters      []struct {
		Clusterid    int64 `json:"clusterid"`
		ExtTalkgroup int64 `json:"extTalkgroup"`
		Repeaterid   int64 `json:"repeaterid"`
		Slot         uint8 `json:"slot"`
		Talkgroup    int64 `json:"talkgroup"`
	} `json:"clusters"`
	DynamicSubscriptions []interface{} `json:"dynamicSubscriptions"`
	Reflector            struct {
		Active    int64 `json:"active"`
		Interval  int64 `json:"interval"`
		Reflector int64 `json:"reflector"`
	} `json:"reflector"`
	StaticSubscriptions []struct {
		Slot      uint8 `json:"slot"`
		Talkgroup int64 `json:"talkgroup"`
	} `json:"staticSubscriptions"`
	TimedSubscriptions []struct {
		Data struct {
			EndDate   int64 `json:"endDate"`
			Friday    int64 `json:"friday"`
			Monday    int64 `json:"monday"`
			Saturday  int64 `json:"saturday"`
			Start     int64 `json:"start"`
			StartDate int64 `json:"startDate"`
			Stop      int64 `json:"stop"`
			Sunday    int64 `json:"sunday"`
			Thursday  int64 `json:"thursday"`
			Tuesday   int64 `json:"tuesday"`
			Wednesday int64 `json:"wednesday"`
		} `json:"data"`
		Recordid   int64 `json:"recordid"`
		Repeaterid int64 `json:"repeaterid"`
		Slot       uint8 `json:"slot"`
		Talkgroup  int64 `json:"talkgroup"`
	} `json:"timedSubscriptions"`
}

// GetProfile fetches the profile of given repeater from BM API
func GetProfile(id string) (*Profile, error) {
	u, err := url.Parse(fmt.Sprint(apiUri, "/repeater/"))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("action", "PROFILE")
	q.Set("q", id)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var p Profile
	err = dec.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
