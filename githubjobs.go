package githubjobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/google/go-querystring/query"
)

// endpoint is the Github Jobs API endpoint to query.
const endpoint = "https://jobs.github.com"

type searchQueries struct {
	Description string `url:"description"`
	Location    string `url:"location"`
	FullTime    bool   `url:"full_time"`
}
type cooridinateQueries struct {
	Latitude  string `url:"lat"`
	Longitude string `url:"long"`
}

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

func (p Position) String() string {
	return stringify(p)
}

// GetPositions gets positions from Github Jobs by description and location.
func GetPositions(description, location string, fullTime bool) ([]*Position, error) {
	sq := searchQueries{
		Description: description,
		Location:    location,
		FullTime:    fullTime,
	}
	v, _ := query.Values(sq)

	url := fmt.Sprintf("%v", endpoint+"/positions.json")

	positions := new([]*Position)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{fmt.Sprintf("Could not create request: %s", err), -1}
	}
	req.URL.RawQuery = v.Encode()

	err = clientRequest(req, positions)
	if err != nil {
		return nil, err
	}

	return *positions, nil
}

// GetPositionsByCoordinates gets positions from Github Jobs by coordinates (latitude and longitude) in decimal degrees.
func GetPositionsByCoordinates(latitude, longitude string) ([]*Position, error) {
	cq := cooridinateQueries{
		Latitude:  latitude,
		Longitude: longitude,
	}
	v, _ := query.Values(cq)

	url := fmt.Sprintf("%v", endpoint+"/positions.json")

	positions := new([]*Position)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{fmt.Sprintf("Could not create request: %s", err), -1}
	}
	req.URL.RawQuery = v.Encode()

	err = clientRequest(req, positions)
	if err != nil {
		return nil, err
	}

	return *positions, nil
}

// GetPositionByID gets a single job posting from Github Jobs by ID.
func GetPositionByID(ID string) (*Position, error) {
	url := fmt.Sprintf("%v", endpoint+"/positions/"+ID+".json")

	position := new(Position)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{fmt.Sprintf("Could not create request: %s", err), -1}
	}

	err = clientRequest(req, position)
	if err != nil {
		return nil, err
	}

	return position, nil
}

func clientRequest(req *http.Request, v interface{}) error {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return Error{fmt.Sprintf("Failed to make request: %s", err), -1}
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return Error{fmt.Sprintf("Could not read JSON response: %s", err), -1}
	}
	if err == io.EOF {
		err = nil
	}
	return err
}

// Stringify attempts to create a reasonable string representation of types in.
// It does things like resolve pointers to their values and omits struct fields with nil values.
func stringify(message interface{}) string {
	var buf bytes.Buffer
	v := reflect.ValueOf(message)
	stringifyValue(&buf, v)
	return buf.String()
}

// stringifyValue was adopted from go-github library https://github.com/google/go-github/blob/master/github/strings.go.
func stringifyValue(w io.Writer, val reflect.Value) {
	if val.Kind() == reflect.Ptr && val.IsNil() {
		w.Write([]byte("<nil>"))
		return
	}

	v := reflect.Indirect(val)

	switch v.Kind() {
	case reflect.String:
		fmt.Fprintf(w, `"%s"`, v)
	case reflect.Slice:
		w.Write([]byte{'['})
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				w.Write([]byte{' '})
			}

			stringifyValue(w, v.Index(i))
		}

		w.Write([]byte{']'})
		return
	case reflect.Struct:
		if v.Type().Name() != "" {
			w.Write([]byte(v.Type().String()))
		}

		w.Write([]byte{'{'})

		var sep bool
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}
			if fv.Kind() == reflect.Slice && fv.IsNil() {
				continue
			}

			if sep {
				w.Write([]byte(", "))
			} else {
				sep = true
			}

			w.Write([]byte(v.Type().Field(i).Name))
			w.Write([]byte{':'})
			stringifyValue(w, fv)
		}

		w.Write([]byte{'}'})
	default:
		if v.CanInterface() {
			fmt.Fprint(w, v.Interface())
		}
	}
}
