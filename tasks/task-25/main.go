package main

import (
	"fmt"
	"time"
)

func MySleep(d time.Duration) {
	ch := make(chan struct{})

	// Deferred start of function, closing channel after d time
	time.AfterFunc(d, func() {
		close(ch)
	})

	<-ch // Block goroutine until channel close
}

func main() {
	fmt.Println("Start:", time.Now())
	MySleep(3 * time.Second)
	fmt.Println("End:", time.Now())
}
