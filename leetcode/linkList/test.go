package main

import "fmt"

func main() {
	done := make(chan string, 1)
	c1, c2 := make(chan int, 1), make(chan int, 1)
	i, j := 0, 10
	go func() {
		for i < 10 {
			<-c1
			i += 1
			fmt.Println("thread1 ", i)
			c2 <- 0
		}
		done <- "done"
	}()
	go func() {
		for j > 0 {
			<-c2
			j -= 1
			fmt.Println(" thread2 ", j)
			c1 <- 0
		}
	}()
	c1 <- 1
	fmt.Println("main ", <-done)
}
