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

	data, err := nasa.NeoFeedOpts(opts)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	expectedLinksNext := "http://www.neowsapp.com/rest/v1/feed?start_date=2015-09-08&end_date=2015-09-09&detailed=false&api_key=DEMO_KEY"
	if data.Links.Next != expectedLinksNext {
		t.Errorf("Links Next mismatch. Got: %s, Want: %s", data.Links.Next, expectedLinksNext)
	}

	expectedElementCount := 23
	if data.ElementCount != expectedElementCount {
		t.Errorf("Element Count mismatch. Got: %d, Want: %d", data.ElementCount, expectedElementCount)
	}

	firstNeo := data.NearEarthObjects["2015-09-08"][0]
	expectedSelfLink := "http://www.neowsapp.com/rest/v1/neo/3726710?api_key=DEMO_KEY"
	if firstNeo.Links.Self != expectedSelfLink {
		t.Errorf("FirstNeo Links Self mismatch. Got: %s, Want: %s", firstNeo.Links.Self, expectedSelfLink)
	}

	expectedId := "3726710"
	if firstNeo.Id != expectedId {
		t.Errorf("FirstNeo Id mismatch. Got: %s, Want: %s", firstNeo.Id, expectedId)
	}

	expectedAbsoluteMagnitudeH := 24.3
	if firstNeo.AbsoluteMagnitudeH != expectedAbsoluteMagnitudeH {
		t.Errorf("FirstNeo Absolute Magnitude mismatch. Got: %f, Want: %f", firstNeo.AbsoluteMagnitudeH, expectedAbsoluteMagnitudeH)
	}

	expectedKilometersEstimatedDiameterMin := 0.0366906138
	if firstNeo.EstimatedDiameter.Kilometers.EstimatedDiameterMin != expectedKilometersEstimatedDiameterMin {
		t.Errorf("Kilometers Estimated Diameter Min mismatch. Got: %f, Want: %f",
			firstNeo.EstimatedDiameter.Kilometers.EstimatedDiameterMin, expectedKilometersEstimatedDiameterMin)
	}

	expectedIsPotentiallyHazardousAsteroid := false
	if firstNeo.IsPotentiallyHazardousAsteroid != expectedIsPotentiallyHazardousAsteroid {
		t.Errorf("Is Potentially Hazardous Asteroid mismatch. Got: %t, Want: %t",
			firstNeo.IsPotentiallyHazardousAsteroid, expectedIsPotentiallyHazardousAsteroid)
	}

	firstCloseApproachData := firstNeo.CloseApproachData[0]
	expectedEpochDateCloseApproach := 1441705500000
	if firstCloseApproachData.EpochDateCloseApproach != expectedEpochDateCloseApproach {
		t.Errorf("Epoch Date Close Approach mismatch. Got: %d, Want: %d",
			firstCloseApproachData.EpochDateCloseApproach, expectedEpochDateCloseApproach)
	}

	expectedKmPerSec := "19.4850295284"
	if firstCloseApproachData.RelativeVelocity.KilometersPerSecond != expectedKmPerSec {
		t.Errorf("Kilometers Per Second mismatch. Got: %s, Want: %s",
			firstCloseApproachData.RelativeVelocity.KilometersPerSecond, expectedKmPerSec)
	}

	expectedAstronomicalMiss := "0.0269230459"
	if firstCloseApproachData.MissDistance.Astronomical != expectedAstronomicalMiss {
		t.Errorf("Astonomical Miss mismatch. Got: %s, Want: %s",
			firstCloseApproachData.MissDistance.Astronomical, expectedAstronomicalMiss)
	}
}
