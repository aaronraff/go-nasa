package nasa

import (
	"testing"
	"time"
)

var nasa *Client

func init() {
	nasa = NewClient()
}

func TestApod(t *testing.T) {
	_, err := nasa.Apod()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestApodOpts(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-12-17")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	opts := &ApodOptions{
		Date: date,
	}

	_, err = nasa.ApodOpts(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestApodBadApiKey(t *testing.T) {
	nasa.SetApiKey("bad_key")
	_, err := nasa.Apod()

	// Reset the key
	nasa.SetApiKey("DEMO_KEY")

	if err == nil {
		t.Errorf("Expected an error.")
	}
}
