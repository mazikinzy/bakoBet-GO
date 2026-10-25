package main

import (
	"fmt"
	"net/http"
)

func main() {


	port := ":8000"

	http.HandleFunc("/", Home)

	fmt.Println("Listening and serving on port", port)

	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to bakobet homepage\n", "\n", "Betting service coming soon")
}


type FamName struct {
	Name string
	Age int
	MaritalStatus string
	Tall bool
}

var family = map[string]FamName {
		"Ikenna": {Name: "Ikemefuna Orji", Age: 35, MaritalStatus: "Married"},
		"Ogochukwu": {Name: "Kingsley Orji", Age: 33, MaritalStatus: "Engaged", Tall: true},
		"Ifeanyi": {Name: "Desmond Orji", Age: 31, MaritalStatus: "Single", Tall: true},
		"BB": {Name: "Blessing Orji", Age: 27, MaritalStatus: "Engaged", Tall: false},
}