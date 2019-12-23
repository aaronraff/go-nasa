package nasa

import (
	"testing"
	"time"
)

func TestNeoFeed(t *testing.T) {
	_, err := nasa.NeoFeed()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestNeoFeedOpts(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2015-09-07")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	opts := &NeoFeedOptions{
		StartDate: date,
		EndDate:   date.AddDate(0, 0, 1),
	}

	_, err = nasa.NeoFeedOpts(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestNeoLookup(t *testing.T) {
	_, err := nasa.NeoLookup(3542519)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestNeoLookupOpts(t *testing.T) {
	opts := &NeoLookupOptions{
		AsteroidId: 2000434,
	}

	_, err := nasa.NeoLookupOpts(3542519, opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
