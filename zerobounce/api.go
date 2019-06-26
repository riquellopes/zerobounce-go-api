package zerobounce

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// URI -
const URI = "https://api.zerobounce.net/v2"

// BuildURL -
func (z *ZeroBounce) BuildURL(endpoint string, params ...url.Values) string {
	value := url.Values{}

	if len(params) == 1 {
		value = params[0]
	}

	value.Set("api_key", z.Apikey)
	return fmt.Sprintf("%s/%s?%s", URI, endpoint, value.Encode())
}

// Request -
func Request(url string, object ResponseBase) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("server error %d", response.StatusCode))
	}

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
	Address       string `json:"address"`
	Status        string `json:"status"`
	SubStatus     string `json:"sub_status"`
	Account       string `json:"account"`
	Domain        string `json:"domain"`
	DidYouMean    string `json:"did_you_mean"`
	DomainAgeDays string `json:"domain_age_days"`
	FreeEmail     bool   `json:"free_email"`
	MXFound       string `json:"mx_found"`
	MXRecord      string `json:"mx_record"`
	SMTPPRovider  string `json:"smtp_provider"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Gender        string `json:"gender"`
	City          string `json:"city"`
	Region        string `json:"region"`
	Zipcode       string `json:"zipcode"`
	Country       string `json:"country"`
	ProcessedAt   string `json:"processed_at"`
}

// GetCreditsResponse -
type GetCreditsResponse struct {
	Credits string `json:"Credits"`
}

// Validate -
func (z *ZeroBounce) Validate(email string) (*ValidateResponse, error) {
	return z.ValidateWithIP(email)
}

// GetCredits -
func (z *ZeroBounce) GetCredits() (int, error) {
	credit := &GetCreditsResponse{}
	err := Request(z.BuildURL("getcredits"), credit)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(credit.Credits)
}

// ValidateWithip -
func (z *ZeroBounce) ValidateWithIP(email string, ipaddress ...string) (*ValidateResponse, error) {
	validate := &ValidateResponse{}
	params := url.Values{}
	addr := ""

	if len(ipaddress) > 0 {
		addr = ipaddress[0]
	}

	params.Set("ip_address", addr)
	params.Set("email", email)

	err := Request(z.BuildURL("validate", params), validate)

	return validate, err
}
