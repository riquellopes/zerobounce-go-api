package zerobounce

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	gock "gopkg.in/h2non/gock.v1"
)

func Test_should_get_95_credits(t *testing.T) {
	defer gock.Off()

	gock.New(URI).
		Get("/getcredits").
		Reply(200).
		JSON(map[string]string{"Credits": "95"})

	zero := ZeroBounce{Apikey: "xxxxxxx"}
	result, _ := zero.GetCredits()

	assert.Equal(t, result, "95")
}

func Test_should_get_error_when_server_response_status_500(t *testing.T) {
	defer gock.Off()

	gock.New(URI).
		Get("/getcredits").
		Reply(500)

	zero := ZeroBounce{Apikey: "xxxxxxx"}
	_, error := zero.GetCredits()

	assert.NotNil(t, error, nil)
}

func Test_should_build_url_with_email(t *testing.T) {
	params := url.Values{}
	params.Set("email", "jonas@example.com")
	zero := ZeroBounce{Apikey: "xxxxxxx"}

	assert.Equal(t, zero.BuildURL("score", params), "https://api.zerobounce.net/v1/score?apikey=xxxxxxx&email=jonas%40example.com")
}

func Test_should_get_url_without_query_string_params(t *testing.T) {
	zero := ZeroBounce{Apikey: "xxxxxxx"}
	assert.Equal(t, zero.BuildURL("score"), "https://api.zerobounce.net/v1/score?apikey=xxxxxxx")
}

func Test_status_should_be_valid(t *testing.T) {
	defer gock.Off()

	response := ValidateWithipResponse{
		Address: "flowerjill@aol.com",
		Status:  "Valid",
	}

	gock.New(URI).
		Get("/validatewithip").
		Reply(200).
		JSON(response)

	zero := ZeroBounce{Apikey: "xxxxxxx"}
	result, _ := zero.ValidateWithip("contato@heriquelopes.com.br")

	assert.Equal(t, result.Address, response.Address)
	assert.Equal(t, result.Status, response.Status)
}
