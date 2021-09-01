package main

import "C"

import (
	"fmt"
	"log"
	"time"

	"github.com/wzshiming/getch"
)

func mainLoop() {
	input := make(chan string)
	tick := time.Tick(time.Millisecond * 1)
	go func() {
		for {
			r, _, err := getch.Getch()
			if err != nil {
				log.Fatal(err)
			}
			input <- string(r)
		}
	}()
	for {
		select {
		case <-tick:
		case v := <-input:
			fmt.Println("input:", v)
		}
	}
}

func main() {
	mainLoop()
}
