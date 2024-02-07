package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"

	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 30})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 30})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 30})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
