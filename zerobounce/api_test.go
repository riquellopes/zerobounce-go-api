package zerobounce

import (
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
