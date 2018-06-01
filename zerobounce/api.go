package zerobounce

// URI -
const URI = "https://api.zerobounce.net/v1"

// ZeroBounce -
type ZeroBounce struct {
	Apikey string
}

// Validate -
func (z *ZeroBounce) Validate(email string) string {
	return ""
}

// GetCredits -
func (z *ZeroBounce) GetCredits() {

}

// ValidateWithip -
func (z *ZeroBounce) ValidateWithip(email, ipaddress string) {

}
