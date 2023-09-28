package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://yahoo.com",
		"http://bing.com",
		"http://amazon.com",
	}

	for _, link := range links {
		getStatus(link)
	}
}

func getStatus(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error while accessing link "+link+" : ", err)
		return
	}
	fmt.Println(link + " link is up and working")

}
