package main

import "fmt"

type Student struct {
	name  string
	email string
}

func main() {
	fmt.Println("Structs session... ")

	var studentA Student = Student{"Mahesh", "test@email.com"}
	fmt.Println("Student A: ", studentA)

	studentb := Student{"Mahesh", "test@email.com"}
	fmt.Println("Student b: ", studentb)

	fmt.Printf("Student Name: %v", studentA)

	ptrsInStructs()

}

func ptrsInStructs() {

	type Employee struct {
		Firstname, Lastname string
		Age, Salary         float64
	}

	emp := Employee{Firstname: "Mahesh", Lastname: "Kulkarni", Age: 38, Salary: 10000}
	empPtr := &Employee{Firstname: "Mahesh", Lastname: "Kulkarni", Age: 38, Salary: 10000}

	fmt.Printf("Details with direct Ref \n :firstName: %+v, lastname: %+v, age: %+v, Salary: %+v \n\n", emp.Firstname, emp.Lastname, emp.Age, emp.Salary)
	fmt.Printf("Details with pointer Ref \n firstName: %+v, lastname: %+v, age: %+v, Salary: %+v \n\n", (*empPtr).Firstname, (*empPtr).Lastname, (*empPtr).Age, (*empPtr).Salary)

	fmt.Printf("Employee : %+v", emp)
	fmt.Printf("Employee with Ptr : %+v", *empPtr)

}
