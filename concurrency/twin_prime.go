// This code based on the Sieve of Eratosthenes Pipeline

package main

import "fmt"

func generate(ch chan <- int) {
	for i:= 2; ;i++ {
		ch <- i
	}
}

// filter primes from the input channel and send them to the output channel
func filter(in <- chan int, out chan <- int, prime int) {
	for {
		num := <-in
		if num%prime != 0 {
			out <- num
		}
	}
}

//// The sieve: generate primes using pipeline of filters
func sieve() {
	ch := make(chan int)
	go generate(ch)

	var prev int = -1 // save previous prime number

	for {
		prime := <-ch
		if prev != -1 && prime-prev == 2 {
			fmt.Printf("Twin primes: (%d, %d)\n", prev, prime)
		}
		prev = prime
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
    sieve()
}
