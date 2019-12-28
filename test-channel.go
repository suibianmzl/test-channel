package main

import (
	"fmt"
	"time"
)

func producer(queue chan <- int)  {
	for i :=0; i < 200; i++  {
		queue <- i
		fmt.Println("create :", i)
	}
}

func consumer(queue <-chan int)  {
	for i :=0; i < 200; i++ {
		v := <-queue
		fmt.Println("receive :", v)
	}
}

func main()  {
	queue := make(chan int, 88)
	go producer(queue)
	go consumer(queue)
	time.Sleep(1 * time.Second)
}