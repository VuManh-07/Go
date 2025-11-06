// xml
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:"name"`
}

func main() {
	person := Person{
		ID:   1,
		Name: "Manh",
	}

	per, err := xml.MarshalIndent(person, " ", " ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(xml.Header + string(per))
}
