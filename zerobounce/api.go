package zerobounce

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// URI -
const URI = "https://api.zerobounce.net/v1"

// BuildURL -
func (z *ZeroBounce) BuildURL(endpoint string, params ...url.Values) string {
	value := url.Values{}

	if len(params) == 1 {
		value = params[0]
	}

	value.Set("apikey", z.Apikey)
	return fmt.Sprintf("%s/%s?%s", URI, endpoint, value.Encode())
}

// Request -
func Request(url string, object ResponseBase) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("server error")
	}

	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(&object)
}

// ZeroBounce -
type ZeroBounce struct {
	Apikey string
}

// ResponseBase -
type ResponseBase interface {
}

// ValidateResponse -
type ValidateResponse struct {
	Address      string      `json:"address"`
	Status       string      `json:"status"`
	SubStatus    string      `json:"sub_status"`
	Account      string      `json:"account"`
	Domain       string      `json:"domain"`
	Disposable   bool        `json:"disposable"`
	Toxic        bool        `json:"toxic"`
	Firstname    string      `json:"firstname"`
	Lastname     string      `json:"lastname"`
	Gender       string      `json:"gender"`
	Location     interface{} `json:"location"`
	Creationdate interface{} `json:"creationdate"`
	Processedat  string      `json:"processedat"`
}

// ValidateWithipResponse -
type ValidateWithipResponse struct {
	Address      string      `json:"address"`
	Status       string      `json:"status"`
	SubStatus    string      `json:"sub_status"`
	Account      string      `json:"account"`
	Domain       string      `json:"domain"`
	Disposable   bool        `json:"disposable"`
	Toxic        bool        `json:"toxic"`
	Firstname    string      `json:"firstname"`
	Lastname     string      `json:"lastname"`
	Gender       string      `json:"gender"`
	Location     interface{} `json:"location"`
	Country      string      `json:"country"`
	Region       string      `json:"region"`
	City         string      `json:"city"`
	Zipcode      string      `json:"zipcode"`
	Creationdate interface{} `json:"creationdate"`
	Processedat  string      `json:"processedat"`
}

// GetCreditsResponse -
type GetCreditsResponse struct {
	Credits string `json:"Credits"`
}

// Validate -
func (z *ZeroBounce) Validate(email string) (*ValidateResponse, error) {
	params := url.Values{}
	params.Set("email", email)

	validate := &ValidateResponse{}

	err := Request(z.BuildURL("validate", params), validate)
	return validate, err
}

// GetCredits -
func (z *ZeroBounce) GetCredits() (string, error) {
	credit := &GetCreditsResponse{}
	err := Request(z.BuildURL("getcredits"), credit)

	return credit.Credits, err
}

// ValidateWithip -
func (z *ZeroBounce) ValidateWithip(email string, ipaddress ...string) (*ValidateWithipResponse, error) {
	validate := &ValidateWithipResponse{}
	params := url.Values{}
	addr := "99.123.12.122"

	if len(ipaddress) > 0 {
		addr = ipaddress[0]
	}

	params.Set("ipaddress", addr)
	params.Set("email", email)

	err := Request(z.BuildURL("validatewithip", params), validate)

	return validate, err
}
