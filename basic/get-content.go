package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func readContent(u string) (string, error) {
	fmt.Printf("Attempting to read content from %s\n", u)
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
