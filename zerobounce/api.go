package zerobounce

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// URI -
const URI = "https://api.zerobounce.net/v1"

// ZeroBounce -
type ZeroBounce struct {
	Apikey string
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
	Credits int `json:"Credits"`
}

// Validate -
func (z *ZeroBounce) Validate(email string) ValidateResponse {
	url := fmt.Sprintf("%s/%s?apikey=%s&email=%s", URI, "validate", z.Apikey, email)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	var validate ValidateResponse

	if err := json.NewDecoder(response.Body).Decode(&validate); err != nil {
		fmt.Println(err)
	}

	return validate
}

// GetCredits -
func (z *ZeroBounce) GetCredits() {

}

// ValidateWithip -
func (z *ZeroBounce) ValidateWithip(email string, ipaddress ...string) ValidateWithipResponse {
	addr := "99.123.12.122"

	if len(ipaddress) > 0 {
		addr = ipaddress[0]
	}

	url := fmt.Sprintf(
		"%s/%s?apikey=%s&email=%s&ipaddress=%s", URI, "validatewithip", z.Apikey, email, addr)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	var validate ValidateWithipResponse

	if err := json.NewDecoder(response.Body).Decode(&validate); err != nil {
		fmt.Println(err)
	}

	return validate
}
