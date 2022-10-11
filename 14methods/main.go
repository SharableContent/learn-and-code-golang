package main

import "fmt"

type Student struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (s *Student) GetStatus() bool {
	fmt.Println("Is User Active: ", s.Status)
	return s.Status
}

func (s *Student) NewEmail() {
	s.Email = "test@test.com"
	fmt.Println("Is User Active: ", s.Status)
}

func main() {
	fmt.Println("Structs session... ")
	var studentA Student = Student{"Mahesh", "test@email.com", true, 35}
	fmt.Printf("Student studentA: %+v \n", studentA)
	fmt.Println("Student Status: ", studentA.GetStatus())
	fmt.Println("Setting new email: ")
	studentA.NewEmail()
	fmt.Printf("Student studentA: %+v \n", studentA)

}
