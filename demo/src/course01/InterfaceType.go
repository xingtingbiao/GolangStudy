package main

import "fmt"

func main() {
	cat := Cat{"Little C", 2, "In the house"}
	cat.Grow()
	animal, ok := interface{}(&cat).(Animal)
	fmt.Printf("%v, %v", ok, animal)
}

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name    string
	age     uint8
	address string
}

func(cat *Cat) Grow() {
	cat.age += 1
}

func(cat *Cat) Move(new string) string {
	old := cat.address
	cat.address = new
	return old
}

