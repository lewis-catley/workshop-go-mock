package main

import (
	"fmt"
)

type application struct {
	web IWebClient
}

func new() application {
	return application{
		web: NewWebClient(),
	}
}

func main() {
	new().printURLContent("https://google.com")
}

func (a application) printURLContent(u string) {
	resp, err := a.web.GetContent(u)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is the content.", resp)
}
