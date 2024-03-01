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

func (c *Client) GetMemberHistoryByEhidOrderByStartDateDesc(
	ehid string,
) (*GetMemberHistoryResponseDto, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/members/%s/history?sort=desc", c.BaseUrl, ehid),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := GetMemberHistoryResponseDto{}
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

func (c *Client) GetCurrentMembershipByEhid(
	ehid string,
) (*GetMemberNodeResponseDto, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/members/%s/nodes", c.BaseUrl, ehid),
		nil,
	)
	if err != nil {
		return nil, err
	}

	data := GetMemberNodeResponseDto{}
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
