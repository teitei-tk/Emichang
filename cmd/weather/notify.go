package main

import (
	"fmt"

	"github.com/teitei-tk/emichang/weather"
)

func main() {
	client, err := weather.NewClient()
	if err != nil {
		panic(err)
	}

	res, err := weather.Request(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Description)
}
