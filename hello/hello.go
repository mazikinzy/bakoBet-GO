package main

import (
	"fmt"
	"github.com/mazikinzy/bakoBet-GO/experiment"
)

func main() {
	msg := experiment.Hello("Mazi")
	fmt.Println("%s, How are you?", msg)
}