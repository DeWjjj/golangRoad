package main

import (
	"fmt"
	"os"
)

func findNumber(number int) bool {
	if allStudentInfor[number] == nil {
		return true
	}
	return false
}

func newStudent(stuName string, stuNumber int) *stuInfor {
	return &stuInfor{
		name:   stuName,
		number: stuNumber,
	}

}

func addNewStudent() {
	var (
		name   string
		number int
	)
	fmt.Println("Enter student name.")
	fmt.Scanln(&name)
	fmt.Println("Enter student number.")
	fmt.Scanln(&number)
	newStu := newStudent(name, number)
	allStudentInfor[number] = newStu
}

func delStudent() {
	var number int
	var quit string
	fmt.Println("Enter student number!")
	fmt.Scanln(&number)
	if findNumber(number) == true {
		fmt.Println("This number is empty!")
		fmt.Println("Enter quit to main!")
		fmt.Scanln(&quit)
		if quit == "quit" {
			main()
		} else {
			delStudent()
		}
	} else {
		fmt.Println("Del Sucessful!")
	}

}

func showStudent() {
	// fmt.Println(allStudentInfor)
	for _, val := range allStudentInfor {
		fmt.Printf("StuNumber:%v StuName:%v\n", val.number, val.name)
	}
	fmt.Println("")
	fmt.Println("")

}

func exit() {
	os.Exit(0) //直接中断退出程序
}

type stuInfor struct {
	name   string
	number int
}

var (
	choose          int
	allStudentInfor map[int]*stuInfor
)

func main() {
	allStudentInfor = make(map[int]*stuInfor, 10)
	for {
		fmt.Println("Welcome to student system!")
		fmt.Println("Add student enter 1")
		fmt.Println("Del student enter 2")
		fmt.Println("Show all student enter 3")
		fmt.Println("Exit enter 4")
		fmt.Scanln(&choose)
		fmt.Printf("Your choose is %v!\n", choose)
		switch choose {
		case 1:
			addNewStudent()
		case 2:
			delStudent()
		case 3:
			showStudent()
		case 4:
			exit()
		default:
			fmt.Println("Please a right number!")
			fmt.Println("")
			fmt.Println("")
		}
	}
}
