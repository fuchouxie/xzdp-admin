package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
}

func (u User) fuck() {
	fmt.Println("fuck")
}

func testType() {
	var user User
	v := reflect.ValueOf(user)
	t := reflect.TypeOf(user)
	fmt.Println(v.Type())
	fmt.Println(v.NumField())
	fmt.Println(t.Name())
	fmt.Println(t.NumMethod())
	t.Method(0).Func.Call(nil)

}

func main() {
	testType()
}
