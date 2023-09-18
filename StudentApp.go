package main

 

import (

"fmt"

"strings"

)

 

// Student struct represents a student with the specified attributes.

type Student struct {

FirstName            string

LastName             string

FullName             string

DOB                  string

Age                  int

SemesterCGPAs        []float64

FinalCGPA            float64

SemesterGrades       []string

FinalGrade           string

YearOfEnrollment     int

YearOfPassing        int

NumberOfYearsToGraduate int

}

 

// CalculateFinalCGPA calculates the final CGPA based on semester CGPAs.

func (s *Student) CalculateFinalCGPA() {

sum := 0.0

for _, cgpa := range s.SemesterCGPAs {

sum += cgpa

}

s.FinalCGPA = sum / float64(len(s.SemesterCGPAs))

}

 

// CalculateFinalGrade calculates the final grade based on semester grades.

func (s *Student) CalculateFinalGrade() {

grades := []string{"AA", "AB", "BB", "BC", "CC", "DD", "FF"}

gradeCount := make(map[string]int)

for _, sg := range s.SemesterGrades {

gradeCount[sg]++

}

 

finalGrade := ""

for _, grade := range grades {

if count, ok := gradeCount[grade]; ok {

finalGrade = grade

if grade == "FF" {

break

}

if count == len(s.SemesterGrades) {

break

}

}

}

s.FinalGrade = finalGrade

}

 

func main() {

students := make(map[string]*Student)

var choice int

 

for {

fmt.Println("Student Record Management")

fmt.Println("1. Create Student")

fmt.Println("2. Read Student")

fmt.Println("3. Update Student")

fmt.Println("4. Delete Student")

fmt.Println("5. Exit")

fmt.Print("Enter your choice: ")

fmt.Scanln(&choice)

 

switch choice {

case 1:

var key string

fmt.Print("Enter student's first name: ")

fmt.Scanln(&key)

if _, exists := students[key]; exists {

fmt.Println("Student already exists.")

break

}

 

var s Student

fmt.Print("Enter student's last name: ")

fmt.Scanln(&s.LastName)

fmt.Print("Enter student's DOB: ")

fmt.Scanln(&s.DOB)

fmt.Print("Enter student's age: ")

fmt.Scanln(&s.Age)

fmt.Print("Enter student's year of enrollment: ")

fmt.Scanln(&s.YearOfEnrollment)

fmt.Print("Enter student's year of passing: ")

fmt.Scanln(&s.YearOfPassing)

fmt.Print("Enter the number of years to graduate: ")

fmt.Scanln(&s.NumberOfYearsToGraduate)

 

fmt.Print("Enter semester CGPAs (comma separated): ")

cgpaInput := ""

fmt.Scanln(&cgpaInput)

cgpaStrings := strings.Split(cgpaInput, ",")

s.SemesterCGPAs = make([]float64, len(cgpaStrings))

for i, cgpaStr := range cgpaStrings {

fmt.Sscanf(cgpaStr, "%f", &s.SemesterCGPAs[i])

}

 

fmt.Print("Enter semester grades (comma separated): ")

gradeInput := ""

fmt.Scanln(&gradeInput)

s.SemesterGrades = strings.Split(gradeInput, ",")

 

s.CalculateFinalCGPA()

s.CalculateFinalGrade()

 

s.FullName = s.FirstName + " " + s.LastName

students[key] = &s

fmt.Println("Student created successfully.")

case 2:

var key string

fmt.Print("Enter student's first name to read: ")

fmt.Scanln(&key)

if student, exists := students[key]; exists {

fmt.Printf("Student Name: %s %s\n", student.FirstName, student.LastName)

fmt.Printf("DOB: %s\n", student.DOB)

fmt.Printf("Age: %d\n", student.Age)

fmt.Printf("Year of Enrollment: %d\n", student.YearOfEnrollment)

fmt.Printf("Year of Passing: %d\n", student.YearOfPassing)

fmt.Printf("Number of Years to Graduate: %d\n", student.NumberOfYearsToGraduate)

fmt.Printf("Semester CGPAs: %v\n", student.SemesterCGPAs)

fmt.Printf("Final CGPA: %.2f\n", student.FinalCGPA)

fmt.Printf("Semester Grades: %v\n", student.SemesterGrades)

fmt.Printf("Final Grade: %s\n", student.FinalGrade)

} else {

fmt.Println("Student not found.")

}

case 3:

var key string

fmt.Print("Enter student's first name to update: ")

fmt.Scanln(&key)

if student, exists := students[key]; exists {

fmt.Print("Enter student's new age: ")

fmt.Scanln(&student.Age)

 

fmt.Print("Enter new semester CGPAs (comma separated): ")

cgpaInput := ""

fmt.Scanln(&cgpaInput)

cgpaStrings := strings.Split(cgpaInput, ",")

student.SemesterCGPAs = make([]float64, len(cgpaStrings))

for i, cgpaStr := range cgpaStrings {

fmt.Sscanf(cgpaStr, "%f", &student.SemesterCGPAs[i])

}

 

fmt.Print("Enter new semester grades (comma separated): ")

gradeInput := ""

fmt.Scanln(&gradeInput)

student.SemesterGrades = strings.Split(gradeInput, ",")

 

student.CalculateFinalCGPA()

student.CalculateFinalGrade()

 

fmt.Println("Student updated successfully.")

} else {

fmt.Println("Student not found.")

}

case 4:

var key string

fmt.Print("Enter student's first name to delete: ")

fmt.Scanln(&key)

if _, exists := students[key]; exists {

delete(students, key)

fmt.Println("Student deleted successfully.")

} else {

fmt.Println("Student not found.")

}

case 5:

fmt.Println("Exiting the application.")

return

default:

fmt.Println("Invalid choice. Please try again.")

}

}

}