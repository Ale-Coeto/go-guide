package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func get(url string) {
	res, err := http.Get(url)
	
	if err != nil {
		fmt.Printf("Error making get request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Client got response. Status code: %d\n", res.StatusCode)
}

func getRequest(url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Client couldn't create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Client error making request: %s", req)
		os.Exit(1)
	}

	fmt.Printf("Client got response. Status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Client couldn't read response body: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Response body: %s\n", resBody)
}

func main() {
	serverPort := 3333
	requestUrl := fmt.Sprintf("http://localhost:%d", serverPort)

	getRequest(requestUrl)

}