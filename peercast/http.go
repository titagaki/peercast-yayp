package peercast

import (
	"net/http"
)

const viewStatXML = "http://localhost:7145/admin?cmd=viewxml"
const username = "root"
const password = "peer"

func requestViewStatXML() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", viewStatXML, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)

	return resp, err
}
