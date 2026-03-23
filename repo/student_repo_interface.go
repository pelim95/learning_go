package repo

import "test/model"

type StudentRepository interface {
	AddStudent(student model.Student) (bool, error)
	AddGrade(report model.ReportCard) (bool, error)
	FindStudentById(id int) (model.Student, error)
	FindGradesByStudentId(id int) ([]model.ReportCard, error)
	FindGradesBySubject(subject string) ([]model.ReportCard, error)
}
