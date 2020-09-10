package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//MysqlConfig type >> exported
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//RedisConfig type >> exported
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
	Test     bool   `ini:"test"`
}

//Config type >> exported
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind()) //*main.MysqlConfig ptr
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") //创建一个非指针err类型报错。
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct") //创建一个非结构体err类型报错。
		return err
	}
	//0.Read line.
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n", lineSlice)
	var structName string
	for idx, line := range lineSlice {
		//1.0 去掉首尾的空格
		line = strings.TrimSpace(line)
		//1.1 如果行内无内容直接跳过此次循环
		if len(line) == 0 {
			continue
		}
		//fmt.Println(line)
		//2.0 判断字符串是否含有；或者#如果是则跳出循环
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//3.0 判断循环出来的line字符串内是否有[，有则进入内部详细判断
		if strings.HasPrefix(line, "[") {
			//3.1 判断是否前后都有[]
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d synatx error", idx+1)
				return
			}
			//3.2 字符串去掉前后[]，判断里面的字符是否存在空的可能
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d synatx error", idx+1)
				return
			}
			//4.0 根据sectionName去data里面根据反射找到对应结构体
			reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			//5.0 如果不是以[开头
			//5.1 以等号分割这一行等号左边是key右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//5.2 根据structName去data里面对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type() //拿到嵌套结构体的类型信息
			structObj := v.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			//5.3 遍历嵌套结构体的每一个字段判断，tag是否为取出的key
			var fieldName string
			var fileType reflect.StructField
			for i := 0; i < structObj.NumField(); i++ {
				filed := sType.Field(i) //tag信息存储在类型信息中
				fileType = filed
				if filed.Tag.Get("ini") == key {
					//5.4 找到对应的字段
					fieldName = filed.Name
					break
				}
			}

			if len(fileName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%dd value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}

		}
	}
	return err
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err: %v.\n", err)
		return
	}
	fmt.Printf("%#v\n", cfg) // 0
}
