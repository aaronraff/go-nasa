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
		panic(err)
	}

	opts := &ApodOptions {
		Date: date,
	}

	data, err := nasa.ApodOpts(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	
	expectedCopyright := "Mark HansonMartin PughSSROPROMPTCTIONSF"
	if data.Copyright != expectedCopyright {
		t.Errorf("Copyright mismatch. Got: %s, Want: %s", data.Copyright, expectedCopyright)
	}

	expectedDate := date
	if !data.Date.Equal(expectedDate) {
		t.Errorf("Date mismatch. Got: %s, Want: %s", data.Date, expectedDate)
	}

	expectedExplanation := "Sculpted by stellar winds and radiation, a magnificent " +
		"interstellar dust cloud by chance has assumed this recognizable shape.  " +
		"Fittingly named the Horsehead Nebula, it is some 1,500 light-years distant, " +
		"embedded in the vast Orion cloud complex. About five light-years \"tall\", the " +
		"dark cloud is cataloged as Barnard 33 and is visible only because its obscuring " +
		"dust is silhouetted against the glowing red emission nebula IC 434.  Stars are " +
		"forming within the dark cloud. Contrasting blue reflection nebula NGC 2023, " +
		"surrounding a hot, young star, is at the lower left of the full image.  The " +
		"featured gorgeous color image combines both narrowband and broadband images " +
		"recorded using several different telescopes."
	if data.Explanation != expectedExplanation {
		t.Errorf("Explanation mismatch. Got: %s, Want: %s", data.Explanation, expectedExplanation)
	}

	expectedHdUrl := "https://apod.nasa.gov/apod/image/1912/Horsehead_Hanson_2604.jpg"
	if data.HdUrl != expectedHdUrl {
		t.Errorf("HdUrl mismatch. Got: %s, Want: %s", data.HdUrl, expectedHdUrl)
	}

	expectedMediaType := "image"
	if data.MediaType != expectedMediaType {
		t.Errorf("MediaType mismatch. Got: %s, Want: %s", data.MediaType, expectedMediaType)
	}

	expectedServiceVersion := "v1"
	if data.ServiceVersion != expectedServiceVersion {
		t.Errorf("ServiceVersion mismatch. Got: %s, Want: %s", data.ServiceVersion, expectedServiceVersion)
	}

	expectedTitle := "The Horsehead Nebula"
	if data.Title != expectedTitle {
		t.Errorf("Title mismatch. Got: %s, Want: %s", data.Title, expectedTitle)
	}

	expectedUrl := "https://apod.nasa.gov/apod/image/1912/Horsehead_Hanson_960.jpg"
	if data.Url != expectedUrl {
		t.Errorf("Url mismatch. Got: %s, Want: %s", data.Url, expectedUrl)
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
