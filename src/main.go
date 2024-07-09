package main

import (
	"time"
)

func Solution(d time.Duration, message string, ch ...chan string) (numberOfAccesses int) {
	time.Sleep(time.Second * d)
	numberOfAccesses = 0
	for _, channel := range ch {
		select {
		case channel <- message:
			numberOfAccesses++
		default:
		}
	}
	return numberOfAccesses
}
