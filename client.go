package client

import (
	"fmt"
	"io"
	"net/http"
)

const HostURL = "https://api.planningcenteronline.com/"

type PC_Client struct {
	Client   *http.Client
	Token    string
	AppID    string
	Endpoint string
}

func NewPCClient(appID, secretToken, endpoint string) *PC_Client {
	fmt.Println("Returning a new PCClient")
	return &PC_Client{
		Client:   &http.Client{},
		AppID:    appID,
		Token:    secretToken,
		Endpoint: endpoint,
	}
}

func (c *PC_Client) doRequest(req *http.Request, secretToken, appID string) ([]byte, error) {

	req.SetBasicAuth(appID, secretToken)

	response, err := c.Client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer response.Body.Close()

	return body, err
}
