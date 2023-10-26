package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string
	LastName  string
}

func main() {
	p1 := person{
		FirstName: "Alona",
		LastName:  "Maksymova",
	}
	p2 := person{
		FirstName: "Polina",
		LastName:  "Maksymova",
	}
	xp := []person{p1, p2}
	bc, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("JSON string", string(bc))

	xp2 := []person{}
	err = json.Unmarshal(bc, &xp2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Objects slice", xp2)
}
