package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// list of websites to check
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
	}

	// creating a channel
	c := make(chan string)

	// looping through links 
	for _, link := range links {
		go checkStatus(link, c)
	}

	// Run status checker
	checkForEver(c)
}

/**
 * Print status one time
 */
func checkOneTime(links []string, c chan string) {
	for i := 0; i < len(links); i++ {
		// wait for the message from goroutine
		<-c
	}
}

/**
 * Check forever
 */
func checkForEver(c chan string) {
	for l := range c {
		// start goroutine
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	resp, err := http.Head(link)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("%s seems to be down\n", link)
		// passing the message to channel
		c <- link
		return
	}

	fmt.Printf("%s is up!\n", link)
	// passing the message to channel
	c <- link

}
