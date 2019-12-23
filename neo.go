package nasa

import (
	"net/http"
	"time"
)

const neoUri string = "https://api.nasa.gov/neo/rest/v1/"

type NasaPaginationLinks struct {
	Next string
	Prev string
	Self string
}

type NasaAsteroidLink struct {
	Self string `json:"self"`
}

type NasaDiameter struct {
	EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
	EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
}

type NasaEstimatedDiameter struct {
	Kilometers NasaDiameter `json:"kilometers"`
	Meters     NasaDiameter `json:"meters"`
	Miles      NasaDiameter `json:"miles"`
	Feet       NasaDiameter `json:"feet"`
}

type NasaRelativeVelocity struct {
	KilometersPerSecond string `json:"kilometers_per_second"`
	KilometersPerHour   string `json:"kilometers_per_hour"`
	MilesPerHour        string `json:"miles_per_second"`
}

type NasaMissDistance struct {
	Astronomical string `json:"astronomical"`
	Lunar        string `json:"lunar"`
	Kilometers   string `json:"kilometers"`
	Miles        string `json:"miles"`
}

type NasaCloseApproachData struct {
	CloseApproachDate      string               `json:"close_approach_date"`
	CloseApproachDateFull  string               `json:"close_approach_date_full"`
	EpochDateCloseApproach int                  `json:"epoch_date_close_approach"`
	RelativeVelocity       NasaRelativeVelocity `json:"relative_velocity"`
	MissDistance           NasaMissDistance     `json:"miss_distance"`
	OrbitingBody           string               `json:"orbiting_body"`
}

type NasaAsteroid struct {
	Links                          NasaAsteroidLink        `json:"links"`
	Id                             string                  `json:"id"`
	NeoReferenceId                 string                  `json:"neo_reference_id"`
	Name                           string                  `json:"name"`
	NasaJplUrl                     string                  `json:"nasa_jpl_url"`
	AbsoluteMagnitudeH             float64                 `json:"absolute_magnitude_h"`
	EstimatedDiameter              NasaEstimatedDiameter   `json:"estimated_diameter"`
	IsPotentiallyHazardousAsteroid bool                    `json:"is_potentially_hazardous_asteroid"`
	CloseApproachData              []NasaCloseApproachData `json:"close_approach_data"`
	IsSentryObject                 bool                    `json:"is_sentry_object"`
}

type NeoFeedResult struct {
	Links            NasaPaginationLinks       `json:"links"`
	ElementCount     int                       `json:"element_count"`
	NearEarthObjects map[string][]NasaAsteroid `json:"near_earth_objects"`
}

// NeoFeedOptions is a struct that represents the options available for NeoFeed API
// calls.
type NeoFeedOptions struct {
	StartDate time.Time
	EndDate   time.Time
}

// NeoFeed calls the Near Earth Object Web Service (Neo) Feed API with the default
// paramters.
func (c *Client) NeoFeed() (*NeoFeedResult, error) {
	return c.NeoFeedOpts(nil)
}

// NeoFeed calls the Near Earth Object Web Service (Neo) Feed API with the given
// options as paramters.
func (c *Client) NeoFeedOpts(opts *NeoFeedOptions) (*NeoFeedResult, error) {
	req, err := http.NewRequest("GET", neoUri+"feed", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	if opts != nil {
		q.Add("start_date", opts.StartDate.Format("2006-01-02"))
		q.Add("end_date", opts.EndDate.Format("2006-01-02"))
	}

	req.URL.RawQuery = q.Encode()

	data := &NeoFeedResult{}
	err = c.get(req, data)

	return data, err
}
