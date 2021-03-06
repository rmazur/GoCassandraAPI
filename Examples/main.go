package main

// dogfood our own app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getPayload(url string) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	bytes, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r.StatusCode, ": ", string(bytes))

	return err
}

func postPayload(url string, payload url.Values) error {

	r, err := myClient.PostForm(
		url,
		payload,
	)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	bytes, _ := ioutil.ReadAll(r.Body)
	fmt.Println(r.StatusCode, ": ", string(bytes))

	return err
}

func main() {
	var err error

	fmt.Println("fetching heartbeat")
	err = getPayload("http://localhost:8080/")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("fetching all users (should be emtpy)")
	err = getPayload("http://localhost:8080/users")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("create a new user")
	err = postPayload("http://localhost:8080/users/new", url.Values{
		"firstname": {"John"},
		"lastname":  {"Doe"},
		"city":      {"Boston"},
		"email":     {"john.doe@example.com"},
		"age":       {"42"},
	})
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("fetching John Doe")
	err = getPayload("http://localhost:8080/users/75a243c1-e356-11e6-81c2-c4b301bb0fa9")
	if err != nil {
		fmt.Println("error:", err)
	}
}
