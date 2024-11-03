package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	GetScoreboard(ctx context.Context, sport, league string) (*ScoreboardResponse, error)
}

type client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *client {
	return &client{httpClient: httpClient}
}

func (c client) GetScoreboard(ctx context.Context, sport, league string) (*ScoreboardResponse, error) {
	uri := fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/%s/%s/scoreboard", sport, league)

	resp, err := c.httpClient.Get(uri)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var scoreboard ScoreboardResponse
	err = json.Unmarshal(body, &scoreboard)
	if err != nil {
		return nil, err
	}
	return &scoreboard, nil
}
