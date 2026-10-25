package main

import "fmt"

func main() {
	form := map[string][]string{
		"name": []string{"Mazi"},
		"surname": []string{"Kinzy"},
		"occupation": []string{"Developer"},
	}

	var form2 = make(map[string]interface{})

	for k, v := range form {
		form2[k] = v[0]
			fmt.Println(form2)
	}
}