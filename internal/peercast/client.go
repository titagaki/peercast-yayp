package peercast

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"peercast-yayp/internal/config"
)

// Client はPeerCastサーバーとの通信を担うHTTPクライアント。
type Client struct {
	cfg        config.PeercastConfig
	httpClient *http.Client
}

// NewClient は新しいPeerCastクライアントを生成する。
func NewClient(cfg config.PeercastConfig) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: &http.Client{},
	}
}

// FetchChannels はPeerCastからチャンネル情報XMLを取得・パースして返す。
func (c *Client) FetchChannels(ctx context.Context) (*StatXML, error) {
	url := fmt.Sprintf("http://%s:%s/admin?cmd=viewxml", c.cfg.Host, c.cfg.Port)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if c.cfg.AuthType == "basic" {
		req.SetBasicAuth(c.cfg.AuthUser, c.cfg.AuthPassword)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch stat xml: %w", err)
	}
	defer resp.Body.Close()

	var data StatXML
	if err := xml.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return &data, nil
}
