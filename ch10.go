package main

import (
	"fmt"
	"time"
)

// Goroutines
// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 		amt := time.Duration(rand.Intn(250))
// 		time.Sleep(time.Millisecond * amt)
// 	}
// }
// func main() {
// 	for i := 0; i < 10; i++ {
// 		go f(i)
// 	}
// 	var input string
// 	fmt.Scanln(&input)
// }

// Channels
// func pinger(c chan<- string) {
// 	for i := 0; ; i++ {
// 		c <- "ping"
// 	}
// }
// func ponger(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "pong"
// 	}
// }
// func printer(c <-chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }
// func main() {
// 	var c chan string = make(chan string)
// 	go pinger(c)
// 	go ponger(c)
// 	go printer(c)
// 	var input string
// 	fmt.Scanln(&input)
// }

// Select
// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
// 	go func() {
// 		for {
// 			c1 <- "from 1"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()
// 	go func() {
// 		for {
// 			c2 <- "from 2"
// 			time.Sleep(time.Second * 3)
// 		}
// 	}()
// 	go func() {
// 		for {
// 			select {
// 			case msg1 := <-c1:
// 				fmt.Println("Message 1", msg1)
// 			case msg2 := <-c2:
// 				fmt.Println("Message 2", msg2)
// 			case <-time.After(time.Second):
// 				fmt.Println("timeout")
// 			}
// 		}
// 	}()
// 	var input string
// 	fmt.Scanln(&input)
// }

// Buffered channels
// type HomePageSize struct {
// 	URL  string
// 	Size int
// }
// func main() {
// 	urls := []string{
// 		"http://www.apple.com",
// 		"http://www.amazon.com",
// 		"http://www.google.com",
// 		"http://www.microsoft.com",
// 		"http://www.dacom.nl",
// 	}
// 	results := make(chan HomePageSize)
// 	for _, url := range urls {
// 		go func(url string) {
// 			res, err := http.Get(url)
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer res.Body.Close()
// 			bs, err := ioutil.ReadAll(res.Body)
// 			if err != nil {
// 				panic(err)
// 			}
// 			results <- HomePageSize{
// 				URL:  url,
// 				Size: len(bs),
// 			}
// 		}(url)
// 	}
// 	var biggest HomePageSize
// 	for range urls {
// 		result := <-results
// 		fmt.Println("Home page:", result.URL, result.Size)

// 		if result.Size > biggest.Size {
// 			biggest = result
// 		}
// 	}
// 	fmt.Println("The biggest home page:", biggest.URL, biggest.Size)
// }

// exercise 2
func Sleep(w time.Duration) {
	<-time.After(w * time.Second)
}
func main() {
	fmt.Println("First message")
	Sleep(2)
	fmt.Println("Second message")
}

/* Exercises
1. <- operator: send-only like s chan <- string / receive-only like s <- chan string
2. see main()
3. Bufferend channels will only block new writes when the results queue is full,
but let send operations to succeed regardless if the receiver is ready.
var c chan int = make(chan int, 20)
*/
