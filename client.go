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

func NewPCClient(appId, secretToken, endpoint string) *PC_Client {
	fmt.Println("Returning a new PCClient")
	return &PC_Client{
		Client:   &http.Client{},
		AppID:    appId,
		Token:    secretToken,
		Endpoint: endpoint,
	}
}

func (c *PC_Client) doRequest(req *http.Request, secretToken, appId string) ([]byte, error) {

	req.SetBasicAuth(appId, secretToken)

	response, err := c.Client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("status : %d\n, body : %s", response.StatusCode, body)
	}

	return body, err
}
