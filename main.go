package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	// printOneTime(links, c)

	checkForEver(c)
}

/**
 * Print status one time
 */
func checkOneTime(links []string, c chan string) {
	for i := 0; i < len(links); i++ {
		<-c
	}
}

/**
 * Check forever
 */
func checkForEver(c chan string) {
	for l := range c {
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
		c <- link
		return
	}

	fmt.Printf("%s is up!\n", link)
	c <- link

}
