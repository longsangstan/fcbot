package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeTextMessage(title string, link string) string {
	title = "*🎥" + title + "*"
	return title + "\n" + link
}

func getRedditJSONString(endPoint string) (string, error) {
	url := endPoint

	// Create a request and add the proper headers.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	req.Header.Set("User-Agent", "Golang Reddit Bot (by /u/longsangstan)")

	// Handle the request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}

	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}

	return string(contents), err
}
