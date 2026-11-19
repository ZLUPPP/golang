package main

import (
	"fmt"
	"strconv"
	"time"
)

func WriteToChannel(ch chan int) {
	ch <- 1
}

func ReadFromChannel(ch chan int) int {
	return <-ch
}

func CloseChannel(ch chan int) {
	close(ch)
}

func ChannelParams(ch chan int) string {
	return "Len: " + strconv.Itoa(len(ch)) + "\nCap: " + strconv.Itoa(cap(ch))
}

func main() {
	start := time.Now()
	since := func() time.Duration { return time.Since(start).Round(100 * time.Millisecond) }

	ch := make(chan string) // unbuffered

	go func() {
		fmt.Println("sender: about to send", since())
		ch <- "ping"
		fmt.Println("sender: send finished", since())
	}()

	fmt.Println("receiver: waiting... ", since())
	time.Sleep(500 * time.Millisecond)
	msg := <-ch
	fmt.Println("receiver: got", msg, " ", since())
	var ch1 chan int     // nil
	ch1 = make(chan int) // open
	//close(ch1)           // close

	go WriteToChannel(ch1)
	fmt.Println(ReadFromChannel(ch1))
	CloseChannel(ch1)
	fmt.Println(ChannelParams(ch1))

	ch2 := make(chan string)

	go func() {
		ch2 <- "goroutine"
	}()

	fmt.Println(<-ch2)

}
