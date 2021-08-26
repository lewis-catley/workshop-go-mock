package main

import (
	"fmt"
)

func main() {
	u := "https://google.com"

	resp, err := readContent(u)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the content.", resp)
}
