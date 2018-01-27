package main

import (
	"fmt"
	"encoding/json"
	"log"
)

func main() {
	m := map[string]string{
		"name": "Carl", 
		"age": "23",
		"game": "Overwatch",
	}

	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(m)
	fmt.Println(string(bs))	
}