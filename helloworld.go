package main

import "fmt"

func main() {
	r()

	ch := make(chan *struct{})
	go func() {
		fmt.Println("Hello world!")
		ch <- nil
	}()
	<-ch
}
