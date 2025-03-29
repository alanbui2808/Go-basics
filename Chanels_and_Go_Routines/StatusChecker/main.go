package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
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

	/*
		receive the output from c chan string
		however this actually terminates when ONE response is spit out of channel and main will terminate
		leaving the rest of routines in the void.
	*/
	// fmt.Println(<-c)

	for i := 0; i < len(links); i++ {
		// Note: Println() is still a blocking call
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	// make a request to link and spit out the status
	_, err := http.Get(link)
	var msg string

	if err != nil {
		fmt.Println("Error:", err, "with:", link)
		msg = "Might be down I think"
	} else {
		msg = link + " is up!"
	}
	// send the msg into the channel
	c <- msg
}
