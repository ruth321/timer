package main

import (
	"fmt"
	"time"
)

func main() {
	//НЕ УДАЛЯТЬ!!!1!11
	var in string
	t := 5
	c := make(chan string, 2)
	d := make(chan string, 1)
	go timeLim(c, t)
	for i := 0; i < 3; i++ {
		go input(c)
		//go timer(c)
		fmt.Println("question: 5+5")
		fmt.Print("answer: ")
		go timer(d, t)
		in = <-c
		if in == "end" {
			fmt.Println("end of time")
			break
		} else {
			d <- "stop"
		}
		fmt.Println(in)
	}
	fmt.Print("asdjkasdl")
}

func timeLim(c chan string, t int) {
	time.Sleep(time.Second * time.Duration(t))
	if len(c) != 0 {
		return
	}
	c <- "end"
	close(c)
}

func input(c chan string) {
	var s string
	_, _ = fmt.Scan(&s)
	if len(c) != 0 {
		return
	}
	c <- s
}

func timer(d chan string, t int) {
	for i := t; i > 0; i-- {
		if s := <-d; s != "stop" {
			fmt.Printf("\ranswer(%d): ", i)
			time.Sleep(time.Second)
		} else {
			return
		}
	}
}
