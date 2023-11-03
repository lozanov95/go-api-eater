package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Requestor interface {
	Do(*http.Request) (*http.Response, error)
}

func main() {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := get[MoviesResponse](&client, "http://maimunda.vloz.website/movies", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Value[0])
}

func get[T any](client Requestor, url string, headers http.Header) (T, error) {
	var result T

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header = headers
	if err != nil {
		return result, err
	}

	response, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
