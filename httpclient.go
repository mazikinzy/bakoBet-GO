package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//Create a http client
	client := &http.Client{}
	resp, err := client.Get("https://bit.ly/kingsleyorji")
	if err != nil {
		fmt.Println("Error facilitating Get", err)
		return
	}
	defer resp.Body.Close()

	//Read and print the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error printing Body", err)
		return
	}
	fmt.Println(body)
}