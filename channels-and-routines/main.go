package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://yahoo.com",
		"http://bing.com",
		"http://amazon.com",
	}

	//var c chan string
	c := make(chan string)

	for _, link := range links {
		go getStatus(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	//for {
	//	go getStatus(<-c, c)
	//}

	//for l := range c {
	//	go getStatus(l, c)
	//}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			getStatus(link, c)
		}(l)
	}

}

func getStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error while accessing link "+link+" : ", err)
		//c <- "Link is down"
		c <- link
		return
	}
	fmt.Println(link + " link is up and working")
	//c <- "Link is up"
	c <- link
}
