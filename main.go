package main

import (
	"fmt"

	. "github.com/riquellopes/zerobounce-go-api/zerobounce"
)

func main() {
	zero := ZeroBounce{}

	fmt.Printf("%s ", zero.Validate("contato@henriquelopes.com.br"))
}
