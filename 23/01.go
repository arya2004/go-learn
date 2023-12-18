package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID = make(chan int)

//unsafee
//var nextID int 


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", <-nextID)

	//nextID ++ // unsafe. inc is read, modify, write
}

func counter(){
	//keep generating nos and put in channel
	for i := 0;; i++ {
		nextID <- i
	}
}

//why channel cant go on forever?


func main() {
	go counter()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3333", nil))
}
