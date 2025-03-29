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
		// "http://golang.org",
		// "http://amazon.com",
	}

	/*
		When a go routine is fired, we need to make sure it's finished executing otherwise the main routine will finish
		first and exit without waiting for the child routine.

		Thus we need channel: Channel manages child routines, once all routines are done then main can exit
		Channel has a type, all routines must have the same type as channel.
		Channel = managing device between child routines and main routine.
	*/
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// launch new routine that waits 5s before checkLink() again
		go func(link string) {
			time.Sleep(5 * time.Second)
			// receives link from channel and continue to spawn another routine
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// make a request to link and spit out the status
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
