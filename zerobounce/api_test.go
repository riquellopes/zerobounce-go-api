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
	result, err := zero.GetCredits()

	assert.NoError(t, err)
	assert.Equal(t, 95, result)
}

func Test_when_the_status_is_500_method_should_get_error(t *testing.T) {
	defer gock.Off()

	gock.New(URI).
		Get("/getcredits").
		Reply(500)

	zero := ZeroBounce{Apikey: "xxxxxxx"}
	_, error := zero.GetCredits()

	assert.NotNil(t, error)
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

	assert.Equal(t, zero.BuildURL("score", params), "https://api.zerobounce.net/v2/score?api_key=xxxxxxx&email=jonas%40example.com")
}

func Test_should_get_url_without_query_string_params(t *testing.T) {
	zero := ZeroBounce{Apikey: "xxxxxxx"}
	assert.Equal(t, zero.BuildURL("score"), "https://api.zerobounce.net/v2/score?api_key=xxxxxxx")
}

func Test_status_should_be_valid(t *testing.T) {
	defer gock.Off()

	response := ValidateResponse{
		Address: "flowerjill@aol.com",
		Status:  "Valid",
	}

	gock.New(URI).
		Get("/validate").
		Reply(200).
		JSON(response)

	zero := ZeroBounce{Apikey: "xxxxxxx"}
	result, _ := zero.ValidateWithIP("contato@heriquelopes.com.br", "99.123.12.122")

	assert.Equal(t, result.Address, response.Address)
	assert.Equal(t, result.Status, response.Status)
}

func Test_status_should_be_valid_when_called_validate_service(t *testing.T) {
	defer gock.Off()

	response := ValidateResponse{
		Address: "flowerjill@aol.com",
		Status:  "Valid",
	}

	gock.New(URI).
		Get("/validate").
		Reply(200).
		JSON(response)

	zero := ZeroBounce{Apikey: "xxxxxxx"}
	result, _ := zero.Validate("contato@heriquelopes.com.br")

	assert.Equal(t, result.Address, response.Address)
	assert.Equal(t, result.Status, response.Status)
}
