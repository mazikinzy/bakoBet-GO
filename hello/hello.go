package main

import (
	"fmt"
	"github.com/mazikinzy/bakobet"
)

func main() {
	msg := bakobet.Hello("Mazi")
	fmt.Println("%s, How are you?", msg)
}