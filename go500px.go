/*
Package go500px provides a client for using the 500px API in Go.
*/
package go500px

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const (
	BaseURL = "https://api.500px.com/v1/"
)

//A Client manages communication with the 500px API.
type Client struct {
	// Consumer Key
	ConsumerKey string
	// Base URL for API requests.
	BaseURL *url.URL
	// HTTP client used to communicate with the API.
	client *http.Client

	// Services used for talking to different parts of the API.
	Photos *PhotosService
}

// NewClient returns a new 500px API client. if a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}
	c.Photos = &PhotosService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified
func (c *Client) NewRequest(method, urlStr string, body string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	if c.ConsumerKey != "" && q.Get("consumer_key") == "" {
		q.Set("consumer_key", c.ConsumerKey)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	//req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response.
// The API response is decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// CheckResponse checks the API response for error, and returns it
// if present. A response is considered an error if it has non StatusOK
// code.
func CheckResponse(r *http.Response) error {
	if r.StatusCode != http.StatusOK {
		return errors.New(r.Status)
	}
	return nil
}
