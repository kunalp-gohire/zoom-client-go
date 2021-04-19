package zoom

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "https://api.zoom.us/v2/users"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}
type AuthResponse struct{

}


// NewClient -
func NewClient(host, token string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Token: token,
		HostURL: HostURL,
	}

	

	if (token != "")  {
		

		// authenticate
		req, err := http.NewRequest("GET", host, nil)
		

		body, err := c.doRequest(req)

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.Token = "Bearer "+token
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
