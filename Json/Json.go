package main

import (
	"encoding/json"
	"fmt"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

type IImage struct {
	//	Raw     []byte `json:"raw"`
	Channel uint32 `json:"channel"`
	Width   uint32 `json:"width"`
	Height  uint32 `json:"height"`
}

type IImageArry struct {
	Parents []IImage
}

func main() {
	/*
		b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
		var m FamilyMember
		// json.Unmarshal 用于解码 json 串
		err := json.Unmarshal(b, &m)
		if err == nil {
			fmt.Printf("name: %s\nAge: %d\nParents: %v\n", m.Name, m.Age, m.Parents[0])
		}
	*/

	c := []byte(`{"Parents":[{"Channel": 1, "Width":100, "Height": 600},
							 {"Channel": 2, "Width":100, "Height": 600}]}`)
	var Image IImageArry
	err := json.Unmarshal(c, &Image)
	if err == nil {
		//fmt.Printf("%s\n", Image.Parents[0].Raw)
		fmt.Printf("%s\n", Image.Parents[0].Channel)
		fmt.Printf("%s\n", Image.Parents[0].Width)
		fmt.Printf("%s\n", Image.Parents[0].Height)
	} else {
		fmt.Println(err.Error())
	}
	var any string
	fmt.Scanln(&any)
}
