package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	channel := make(chan string)
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://amazon.com",
	}

	for _, link := range links {
		go checkLink(link, channel)
	}

	for link := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	}

	fmt.Println(link, "is up!")
	channel <- link
}
