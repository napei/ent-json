package main

import (
	"encoding/json"
	"fmt"
)

type omitted struct {
	StringField string  `json:",omitempty"`
	IntField    int64   `json:",omitempty"`
	FloatField  float64 `json:",omitempty"`
	BoolField   bool    `json:",omitempty"`
}

type notOmitted struct {
	StringField string
	IntField    int64
	FloatField  float64
	BoolField   bool
}

func main() {
	o := omitted{
		StringField: "",
		IntField:    0,
		FloatField:  0,
		BoolField:   false,
	}

	// Prints {}
	jo, _ := json.Marshal(o)
	fmt.Println(string(jo))

	n := notOmitted{
		StringField: "",
		IntField:    0,
		FloatField:  0,
		BoolField:   false,
	}

	// Prints {"StringField":"","IntField":0,"FloatField":0,"BoolField":false}
	jn, _ := json.Marshal(n)
	fmt.Println(string(jn))
}
