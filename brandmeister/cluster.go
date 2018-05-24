package brandmeister

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Cluster is a cluster of repeaters as fetched from
// https://api.brandmeister.network/v1.0/repeater/?action=LIST
type Cluster struct {
	Clusterid   string `json:"clusterid"`
	Clustername string `json:"clustername"`
	Masterid    string `json:"masterid"`
	Talkgroup   string `json:"talkgroup"`
}

// GetCluster fetches a cluster from the BM API
func GetCluster(id string) (*Cluster, error) {
	u, err := url.Parse(fmt.Sprint(apiUri, "/cluster/"))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("action", "GET")
	q.Set("id", id)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var c Cluster
	err = dec.Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
