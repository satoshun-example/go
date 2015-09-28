package main

import "fmt"

func main() {
	ch := make(chan *struct{})
	go func() {
		fmt.Println("Hello world!")
		ch <- nil
	}()
	<-ch
}
