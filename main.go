package githubjobs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// endpoint is the Github Jobs API endpoint to query.
const endpoint = "https://jobs.github.com"

// Position defines a position returned by Github Jobs.
type Position struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	CompanyLogo string `json:"company_logo"`
	URL         string `json:"url"`
}

// Error defines an error received when making a request to the API.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Error returns a string representing the error, satisfying the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("githubjobs: %s (%d)", e.Message, e.Code)
}

// GetPositions gets positions from Github Jobs by description and location.
func GetPositions(description, location string, fullTime bool) ([]Position, error) {
	v := url.Values{}
	v.Set("description", description)
	v.Set("location", location)
	v.Set("full_time", "false")
	if fullTime {
		v.Set("full_time", "true")
	}

	url := fmt.Sprintf("%v", endpoint+"/positions.json?"+v.Encode())

	r, err := http.Get(url)
	if err != nil {
		return nil, Error{fmt.Sprintf("Could not create request: %s", err), -1}
	}
	defer r.Body.Close()

	var p *[]Position
	json.NewDecoder(r.Body).Decode(&p)

	return *p, nil
}

// GetPositionsByCoordinates gets positions from Github Jobs by coordinates (latitude and longitude) in decimal degrees.
func GetPositionsByCoordinates(latitude, longitude string) ([]Position, error) {
	v := url.Values{}
	v.Set("lat", latitude)
	v.Set("long", longitude)

	url := fmt.Sprintf("%v", endpoint+"/positions.json?"+v.Encode())

	r, err := http.Get(url)
	if err != nil {
		return nil, Error{fmt.Sprintf("Could not create request: %s", err), -1}
	}
	defer r.Body.Close()

	var p *[]Position
	json.NewDecoder(r.Body).Decode(&p)

	return *p, nil
}

// GetPositionByID gets a single job posting from Github Jobs by ID.
func GetPositionByID(ID string) (Position, error) {
	url := fmt.Sprintf("%v", endpoint+"/positions/"+ID+".json")

	r, err := http.Get(url)
	if err != nil {
		return Position{}, Error{fmt.Sprintf("Could not create request: %s", err), -1}
	}
	defer r.Body.Close()

	var p *Position
	json.NewDecoder(r.Body).Decode(&p)

	return *p, nil
}
