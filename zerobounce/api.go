package zerobounce

import (
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
func (z *ZeroBounce) Validate(email string) string {

	resp, err := http.Get(fmt.Sprintf("%s/%s", URI, "validate"))

	if err != nil {
		fmt.Println(err)
	}

	return ""
}

// GetCredits -
func (z *ZeroBounce) GetCredits() {

}

// ValidateWithip -
func (z *ZeroBounce) ValidateWithip(email, ipaddress string) {

}
