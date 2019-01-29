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
		Name string
	}
	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	// 修改结构体的属性值需要用指针
	valueOfCat := reflect.ValueOf(&cat{})
	fmt.Println(typeOfCat.NumField())
	fmt.Println(typeOfCat.Field(0).Name)
	name := valueOfCat.Elem().FieldByName("Name")
	// 赋值需要结构体的属性可导出, 即属性首字母为大写, 类似java的public 和 private
	name.SetString("豆芽")
	fmt.Println(name.String())

	typeOfZero := reflect.TypeOf(Zero)
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
}
