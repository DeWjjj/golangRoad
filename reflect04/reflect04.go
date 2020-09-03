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
	t := reflect.TypeOf(stu1)           //t赋予变量stu1结构体的信息
	fmt.Println(t.Name(), t.Kind())     //student struct
	for i := 0; i < t.NumField(); i++ { //NumField()函数可以得到所有结构体中的字段信息
		field := t.Field(i) //Field()函数可以得到结构体中的字段信息
		fmt.Printf("name: %s index: %d type: %v json_tag: %v.\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name: %s index: %d type: %v json tag:%v.\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
