package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			Telephone{
				Mobile: true,
				Number: "1234-576",
			},
			Telephone{
				Mobile: false,
				Number: "abcc-567",
			},
		},
	}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(rec))

	var unRec Record
	err = json.Unmarshal(rec, &unRec)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(unRec)
}
