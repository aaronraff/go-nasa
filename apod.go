package nasa

import (
	"net/http"
	"time"
	"strings"
)

const uri string = "https://api.nasa.gov/planetary/apod"

// Custom struct to implement Unmarshallar interface. This allows the decoder
// to decode the date in the response.
type NasaDate struct {
	time.Time
}

// Structure to hold the APOD API call's response data.
type ApodResult struct {
	Copyright string `json:"copyright"`
	Date NasaDate `json:"date"`
	Explanation string `json:"explanation"`
	HdUrl string `json:"hdurl"`
	MediaType string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title string `json:"title"`
	Url string `json:"url"`
}

type ApodOptions struct {
	Date time.Time
	Hd bool
}

// Call the Astronomy Picture of the Day (APOD) API with the default paramters.
func (c *Client) Apod() (*ApodResult, error) {
	return c.ApodOpts(nil)
}

// Call the Astronomy Picture of the Day (APOD) API with the given options.
func (c *Client) ApodOpts(opts *ApodOptions) (*ApodResult, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	if opts != nil {
		q.Add("date", opts.Date.Format("2006-01-02"))

		if opts.Hd == true {
			q.Add("hd", "True")
		} else {
			q.Add("hd", "False")
		}
	}

	req.URL.RawQuery = q.Encode()

	data := &ApodResult{}
	err = c.get(req, data)
	
	return data, err
}

// Implement the Unmarshallar interface for the NasaData struct.
func (d *NasaDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	*d = NasaDate { t }
	return err
}
