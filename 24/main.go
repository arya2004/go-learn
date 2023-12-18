package main

import (
	"log"
	"time"
)

type result struct {
	url string
	err error
	latency time.Duration
}

//channel
//<- rstricts to write only to channel




func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan <- int){
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i + 1, chans[i])
	}
//select selects the channel which is ready, thus not blocking other channel

	for i := 0; i < 12; i++ {

		select {
		case i1 := <-chans[0]:
			log.Println("ch1", i1)
		case i2 := <-chans[1]:
			log.Println("ch2", i2) 
		}
		
	}
}
