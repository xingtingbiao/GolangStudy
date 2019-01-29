package main

import (
	"fmt"
	"reflect"
)

/*
Go 的反射练习
*/
func main() {
	//testReflect01()
	testReflect02()
}

func testReflect01() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}

type Enum int

const Zero Enum = 0

func testReflect02() {
	type cat struct {
	}
	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	typeOfZero := reflect.TypeOf(Zero)
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
}
