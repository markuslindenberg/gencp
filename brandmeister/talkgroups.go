package brandmeister

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTalkgroups fetches all talkgroups from the BM API
func GetTalkgroups() (map[string]string, error) {
	resp, err := http.Get(fmt.Sprint(apiUri, "/groups/"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var g map[string]string
	err = dec.Decode(&g)
	if err != nil {
		return nil, err
	}

	return g, nil
}
