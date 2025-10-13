package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Age  int
	Name string
}

type StudentFunc interface {
	AddStudent(name string, age int)
	ListStudents()
	UpdateStudent(index int, newName string, newAge int)
	DeleteStudent(index int)
}

var students []*Todo
var manager = &Todo{}

func (t *Todo) AddStudent(name string, age int) {
	students = append(students, t)
}
func (t *Todo) DeleteStudent(index int) {
	if index < 0 || index >= len(students) {
		fmt.Println("Invalid index")
		return
	}
	students = append(students[:index], students[index+1:]...)
}

func (t *Todo) ListStudents() {
	fmt.Printf("No.\t| Name\t\t| Age\n")
	fmt.Printf("----\t| ----\t\t| ---\n")
	for i, val := range students {
		fmt.Printf("%d\t| %s\t\t| %d\n", i+1, val.Name, val.Age)
	}
}

func (t *Todo) UpdateStudent(index int, newName string, newAge int){
if index < 0 || index >= len(students) {
		fmt.Println("Invalid index")
		return
	}
 students[index].Name = newName
	students[index].Age = newAge


}

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("---------------")
		fmt.Println("Choose Operation (lowercase or uppercase)")
		fmt.Println("1.ADD")
		fmt.Println("2.DELETE")
		fmt.Println("3.EDIT")
		fmt.Println("4.VIEW")
		fmt.Println("5.EXIT")
		fmt.Println("--------")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.EqualFold(input, "exit") {
			fmt.Println("exiting the program")
			break
		}

		if strings.EqualFold(input, "add") {
			fmt.Println("Enter your name or enter exit")
			input1, _ := reader.ReadString('\n')
			input1 = strings.TrimSpace(input1)


			fmt.Println("Enter your age")
			input2, _ := reader.ReadString('\n')
			input2 = strings.TrimSpace(input2)
			convInt, _ := strconv.Atoi(input2)
			s1 := &Todo{Name: input1, Age: convInt}
			s1.AddStudent(input1, convInt)

		}

		if strings.EqualFold(input, "delete") {
			if len(students) == 0 {
				fmt.Println("No data to delete")
			} else {
				fmt.Println("Choose the index you want to delete !")
				fmt.Printf("No.\t| Name\t\t| Age\n")
				fmt.Printf("----\t| ----\t\t| ---\n")
				for i, val := range students {
					fmt.Printf("%d\t| %s\t\t| %d\n", i, val.Name, val.Age)
				}
				input3, _ := reader.ReadString('\n')
				input3 = strings.TrimSpace(input3)
				convToInt, _ := strconv.Atoi(input3)
				manager.DeleteStudent(convToInt)
			}

		}

		if strings.EqualFold(input, "view") {
			manager.ListStudents()
		}

		if strings.EqualFold(input,"edit"){
			fmt.Println("Choose the index of the student you want to update !")
			fmt.Printf("No.\t| Name\t\t| Age\n")
				fmt.Printf("----\t| ----\t\t| ---\n")
				for i, val := range students {
					fmt.Printf("%d\t| %s\t\t| %d\n", i, val.Name, val.Age)
				}
				input4, _ := reader.ReadString('\n')
				input4 = strings.TrimSpace(input4)
				convToInt, _ := strconv.Atoi(input4)

				fmt.Println("Enter the new name")
				input5, _ := reader.ReadString('\n')
				input5 = strings.TrimSpace(input5)
    
				fmt.Println("Enter the new age")
				input6, _ := reader.ReadString('\n')
				input6 = strings.TrimSpace(input6)
				ageInt, _ := strconv.Atoi(input6)

				manager.UpdateStudent(convToInt,input5,ageInt)

		}


	}

}
