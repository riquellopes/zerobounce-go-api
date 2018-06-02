package main

import (
	"fmt"
	"os"

	. "github.com/riquellopes/zerobounce-go-api/zerobounce"
)

func main() {
	zero := ZeroBounce{Apikey: os.Getenv("API_KEY_ZERO")}

	fmt.Println(zero.Validate("contato@henriquelopes.com.br"))
	fmt.Println(zero.ValidateWithip("contato@henriquelopes.com.br"))
	fmt.Println(zero.GetCredits())
}
