package nasa

import (
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
)

// Used to make all API calls. Keeps a hold of the API key as well as the http
// client that should be used for making requests.
type Client struct {
	apiKey string
	httpClient *http.Client
}

// Creates a client using the API key: DEMO_KEY.
// This is useful for testing, but you will experience stricter rate-limits 
// (https://api.nasa.gov/#authentication).
func NewClient() *Client {
	return &Client { 
		apiKey: "DEMO_KEY",
		httpClient: &http.Client{},
	}
}

// Sets the API key for all future requests.
// You can sign up for an API key here (https://api.nasa.gov/#signUp).
func (c *Client) SetApiKey(key string) {
	c.apiKey = key
}

func (c *Client) get(req *http.Request, data interface{}) error {
	q := req.URL.Query()
	q.Add("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		errString := fmt.Sprintf("Received status code: %d", res.StatusCode)
		return errors.New(errString)
	}

	return json.NewDecoder(res.Body).Decode(data)
}
