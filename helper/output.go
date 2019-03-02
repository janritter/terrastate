package helper

import (
	"fmt"
	"reflect"
)

func PrintStateValues(in interface{}) {
	v := reflect.ValueOf(in)
	fmt.Println("")
	fmt.Println("------- Using the following values -------")
	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		fmt.Printf("%s = %s \n", name, v.Field(i).Interface())
	}
	fmt.Println("------------------------------------------")
}
