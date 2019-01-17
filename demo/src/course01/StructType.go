package main

import "fmt"

func main() {
	p := Person{"Robert", "Male", 33, "beijing"}
	oldAddr := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddr, p.Address)
}

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (person *Person) Move(newAddr string) string {
	oldAddr := person.Address
	person.Address = newAddr
	return oldAddr
}


