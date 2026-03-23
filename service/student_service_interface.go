package service

import "test/model"

type StudentService interface {
	AddStudentGrade(student model.Student, reportCard model.ReportCard) (string, error)
	CalculateStudentAverage(studentId int) (float64, error)
	FindTopStudent(subjectName string) (string, error)
	FilterStudents(id int) (string, error)
}
