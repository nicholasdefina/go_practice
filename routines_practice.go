package main

import (
	"fmt"
	"net/http"
	"time"
)

// Go Scheduler supervises Go Routines. If one finishes or makes a blocking call, Scheduler switches to another Routine.
// If only one core, concurrency is best outcome. Multiple cores allow for true parallelism.

func siteChecker() {

	// Routines need to be communicated by Channels. The main routine and all child routines comunicate via a channel.
	// receiving messages from a channel is a blocking line of code!
	c := make(chan string)

	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, site := range sites {
		go checkSite(site, c) // 'go' keyword launches new go routine. allows for concurrency. pass in channel.
	}

	for site := range c { // infinite loop
		go func(site string) { // function  literal like python lambda. delay calls with time.Sleep
			time.Sleep(time.Second * 5)
			checkSite(site, c)
		}(site)
	}

}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "might be down! Error: ", err)
		c <- site
		return
	}
	fmt.Println(site, "is working fine.")
	c <- site // send string to channel
}
