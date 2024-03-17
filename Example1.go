package main

import (
	"fmt"
)

type Student struct {
	Roll  int
	Name  string
	Marks [5]int
}

func main() {
	var student Student

	fmt.Print("Enter Roll: ")
	fmt.Scanln(&student.Roll)

	fmt.Print("Enter Name: ")
	fmt.Scanln(&student.Name)

	fmt.Println("Enter Marks for 5 Subjects:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Subject %d: ", i+1)
		fmt.Scanln(&student.Marks[i])
	}

	fmt.Println("\nStudent Details:")
	fmt.Println("Roll:", student.Roll)
	fmt.Println("Name:", student.Name)
	fmt.Println("Marks:", student.Marks)
}
