package main

import (
	"log"

	"github.com/MatthiasHeylenTM2021/packages/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	RandomNumber := helpers.RandomNumber(numPool)
	intChan <- RandomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)
}
