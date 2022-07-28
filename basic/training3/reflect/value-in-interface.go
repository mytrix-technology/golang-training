package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe  variable   :", reflectValue.Type())
	fmt.Println("nilai variable   :", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)
	fmt.Println("nilai asli (int) :", nilai)
}
