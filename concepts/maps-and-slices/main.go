package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	myMap := make(map[string]User)

	me := User {
		FirstName: "Matthias",
		LastName: "Heylen",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)
}
