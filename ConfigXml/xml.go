package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

func main() {
	file, _ := os.Open("config.xml")
	var person Person
	decoder := xml.NewDecoder(file)
	decoder.Decode(&person)
	fmt.Println(person.Age)
	fmt.Println(person.Name)

	var any string
	fmt.Scanln(&any)
}
