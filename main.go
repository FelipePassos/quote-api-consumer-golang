package main

import (
	"cotacao-go/apis"
	"cotacao-go/generators"
	"fmt"
)

func main() {

	response, err := apis.InitializeApi()

	if err != nil {
		fmt.Println("No Data")
	}

	generators.GenereatorDocFile(response)

}
