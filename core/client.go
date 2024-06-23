package core

import (
	"fmt"
	"io"
	"net/http"
)

const hostURL = "https://api.planningcenteronline.com/"

type PC_Client struct {
	Client   *http.Client
	Token    string
	AppID    string
	Endpoint string
}

func NewPCClient(appId, secretToken string) *PC_Client {
	fmt.Println("Returning a new PCClient")
	return &PC_Client{
		Client:   &http.Client{},
		AppID:    appId,
		Token:    secretToken,
		Endpoint: hostURL,
	}
}

func (c *PC_Client) DoRequest(req *http.Request) ([]byte, error) {

	req.SetBasicAuth(c.AppID, c.Token)

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error during Client.Do: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error during Client.Do: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("\nResponse Code : %d\nBody : %s\n", response.StatusCode, body)
	}

	return body, err
}
