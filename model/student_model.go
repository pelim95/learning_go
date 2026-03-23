package model

import "fmt"

type Student struct {
	Id   int
	Name string
}

type ReportCard struct {
	Id        int
	Subject   string
	Grade     float64
	StudentId int
}

func (s Student) String() string {
	return fmt.Sprintf("Student name is %s with ID %d", s.Name, s.Id)
}

func (r ReportCard) String() string {
	return fmt.Sprintf("Their result in %s is %.2f", r.Subject, r.Grade)
}
