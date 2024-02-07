package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("MyString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("After func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
