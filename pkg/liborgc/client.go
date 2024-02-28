package liborgc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseUrl string
}

func NewClient(host string, port int) *Client {
	return &Client{
		BaseUrl: fmt.Sprintf("%s:%d", host, port),
	}
}

func (c *Client) GetMembershipHistoryByEhidOrderByStartDateDesc(
	ehid string,
) (*GetMembershipHistoryResponseDto, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/members/%s/history?sort=desc", c.BaseUrl, ehid),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := GetMembershipHistoryResponseDto{}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
