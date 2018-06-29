[![Build Status](https://travis-ci.org/riquellopes/zerobounce-go-api.svg?branch=master)](https://travis-ci.org/riquellopes/zerobounce-go-api)
[![Coverage Status](https://coveralls.io/repos/github/riquellopes/zerobounce-go-api/badge.svg?branch=master)](https://coveralls.io/github/riquellopes/zerobounce-go-api?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/riquellopes/zerobounce-go-api)](https://goreportcard.com/report/github.com/riquellopes/zerobounce-go-api)

Zerobounce Go Api
-----------------
[Go](http://golang.org) integration for [ZeroBounce](https://www.zerobounce.net/) service.

# Installation #
```sh
go get github.com/riquellopes/zerobounce-go-api
```

# Usage #
First you will need an api token. Can you create an account at [here](https://www.zerobounce.net/members/register/) and get you token.

### Quick start ###
```go
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
```
