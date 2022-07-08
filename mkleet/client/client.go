package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://leetcode.com/graphql"

func New() *Client {
	return &Client{
		client: &http.Client{
			Transport: http.DefaultTransport,
		},
	}
}

type Client struct {
	client *http.Client
}

type RequestPlaceholder struct {
	OperationName string            `json:"operationName"`
	Query         string            `json:"query"`
	Variables     map[string]string `json:"variables"`
}

func (c *Client) request(query *RequestPlaceholder) ([]byte, error) {
	reqData, err := json.Marshal(query)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36`)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code %d. expected 200", resp.StatusCode)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) GetQuestionData(slug string) (*GetQuestionDataResponse, error) {
	reqData := &RequestPlaceholder{
		OperationName: "questionData",
		Variables: map[string]string{
			"titleSlug": slug,
		},
		Query: questionDataQuery,
	}

	var response *GetQuestionDataResponse
	res, err := c.request(reqData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &response)

	return response, err
}
