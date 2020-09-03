package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "Dew",
		Score: 100,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name: %s index: %d type: %v json_tag: %v.\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name: %s index: %d type: %v json tag:%v.\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
