package peercast

import (
	"fmt"
	"net/http"

	"peercast-yayp/config"
)

func requestViewStatXML() (*http.Response, error) {
	cfg := config.GetConfig()
	viewStatXML := fmt.Sprintf("http://%s:%s/admin?cmd=viewxml",
		cfg.Peercast.Host, cfg.Peercast.Port)

	client := &http.Client{}
	req, err := http.NewRequest("GET", viewStatXML, nil)
	if err != nil {
		return nil, err
	}
	if cfg.Peercast.AuthType == "basic" {
		req.SetBasicAuth(cfg.Peercast.AuthPassword, cfg.Peercast.AuthPassword)
	}
	resp, err := client.Do(req)

	return resp, err
}
