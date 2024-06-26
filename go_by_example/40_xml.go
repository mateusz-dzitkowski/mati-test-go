package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v name=%v origin=%v", p.Id, p.Name, p.Origin)
}

type Nesting struct {
	XMLName xml.Name `xml:"nesting"`
	Plants  []*Plant `xml:"parent>child>plant"`
}

func main() {
	coffee := &Plant{
		Id:     12,
		Name:   "Coffee",
		Origin: []string{"Ethiopia", "Brazil"},
	}
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{
		Id:     42,
		Name:   "Tomato",
		Origin: []string{"Mexico", "California"},
	}

	nesting := &Nesting{Plants: []*Plant{coffee, tomato}}
	out2, _ := xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out2))
}
