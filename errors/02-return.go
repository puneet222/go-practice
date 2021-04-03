package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person2 struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person2{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bs)

}

// toJSON needs to return an error also
func toJSON(a interface{}) (string,error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error while converting data to JSON %v", err))
	}
	return string(bs), nil
}
