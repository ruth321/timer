package main

import (
	"fmt"
	"time"
)

func main() {
	//НЕ УДАЛЯТЬ!!!
	var in string
	for i := 0; i < 3; i++ {
		c := make(chan string)
		go timeLim(c)
		go input(c)
		//go timer(c)
		fmt.Println("question: 5+5")
		fmt.Print("answer:")
		in = <-c
		if in == "end" {
			fmt.Println("end of time")
		}
		fmt.Println(in)
	}

	/*
		fmt.Print("Time: ")
		for i := 5; i > 0; i-- {
			fmt.Print(i)
			time.Sleep(time.Second)
			fmt.Printf("\r\b\r")
		}
	*/

}

func timeLim(c chan string) {
	time.Sleep(time.Second * 3)
	if len(c) != 0 {
		return
	}
	c <- "end"
}

func input(c chan string) {
	var s string
	if len(c) != 0 {
		return
	}
	_, _ = fmt.Scan(&s)
	c <- s
}

func timer(c chan string) {
	if len(c) == 0 {
		for i := 3; i > 0; i-- {
			fmt.Print(i, "): ")
			time.Sleep(time.Second)
			fmt.Printf("\b\b\b\b")
		}
	}
}
