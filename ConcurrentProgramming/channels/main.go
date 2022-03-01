package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("unbuffered:")
	unbufferedChannels()

	fmt.Println("buffered:")
	bufferedChannels()

	fmt.Println("channel types:")
	channelTypes()

	fmt.Println("closing channels:")
	closingChannels()

	fmt.Println("control flow:")
	ifControlFlows()
	forControlFlows()

	// fmt.Println("select statements:") see course demo...

}

func unbufferedChannels() {
	wg := &sync.WaitGroup{}
	ch := make(chan int) // unbuffered channel

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) { // this is a birectional channnel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) { // send only channel
		ch <- 42
		// ch <- 27 // this results in a deadlock
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func bufferedChannels() {

	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // buffered channel

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27 // only works with buffered channel (never prints but doesn't hang app)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func channelTypes() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // buffered channel

	wg.Add(2)
	// go func(ch chan int, wg *sync.WaitGroup) { // this is a birectional channnel
	go func(ch <-chan int, wg *sync.WaitGroup) { // recieve only channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) { // send only channel
		ch <- 42
		ch <- 27 // only works with buffered channel, otherwise deadlock
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func closingChannels() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // buffered channel

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		close(ch)         // can't closer on a recieve only channel
		fmt.Println(<-ch) // recieving on a close channel returns 0 value
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		// close(ch)
		// ch <- 27 // this panics..no way to check if channel is closed
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func ifControlFlows() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		// msg, ok := <-ch // if channel is closed wil get 0 an false
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		// ch <- 0
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func forControlFlows() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		} // if channel isn't closed at end will result in deadlock
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
