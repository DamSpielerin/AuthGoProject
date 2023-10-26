package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	FirstName string
	LastName  string
}

func main() {
	//p1 := person{
	//	FirstName: "Alona",
	//	LastName:  "Maksymova",
	//}
	//p2 := person{
	//	FirstName: "Polina",
	//	LastName:  "Maksymova",
	//}
	//xp := []person{p1, p2}
	//bc, err := json.Marshal(xp)
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//fmt.Println("JSON string", string(bc))
	//
	//xp2 := []person{}
	//err = json.Unmarshal(bc, &xp2)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println("Objects slice", xp2)
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/encode2", foo2)
	http.HandleFunc("/decode", bar)
	http.HandleFunc("/decode2", bar2)
	http.ListenAndServe(":8089", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		FirstName: "Alona",
		LastName:  "Maksymova",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Issue with encoding!!!")
	}
}
func foo2(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		FirstName: "Alona",
		LastName:  "Maksymova",
	}
	p2 := person{
		FirstName: "Polina",
		LastName:  "Maksymova",
	}
	xp := []person{p1, p2}
	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		log.Println("Issue with slice encoding!!!", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Issue with decoding!!!", err)
	}
	log.Println("Person", p1)

}
func bar2(w http.ResponseWriter, r *http.Request) {
	xp := []person{}
	err := json.NewDecoder(r.Body).Decode(&xp)
	if err != nil {
		log.Println("Issue with slice decoding!!!", err)
	}
	log.Println("Person", xp)

}
