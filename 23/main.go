package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	url string
	err error
	latency time.Duration
}

//channel
//<- rstricts to write only to channel
func get(url string, ch chan<- result){
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}




func main() {
	main_start := time.Now()
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	for _, url := range list {
		go get(url, results)
	}
	// arrow before channel name means do or read on 'results' channel

	//if range over channel, it will read till channel closses, but we aint clossing channel
	//to ensure we dont read more than write
	//only 1 can close channel
	for range list {
		r := <- results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		}else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}

	fmt.Println(time.Since(main_start).Round(time.Millisecond))

}
