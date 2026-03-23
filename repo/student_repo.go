package repo

import (
	"database/sql"
	"fmt"
	"log"
	"test/model"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connectionString := "postgresql://postgres:postgres@localhost:5432/student"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected to database")

	return db
}

type StudentRepoImpl struct {
	studentDb *sql.DB
}

func NewStudentRepoImpl(studentDb *sql.DB) *StudentRepoImpl {
	studentDb = Connect()
	return &StudentRepoImpl{studentDb}
}

func (r *StudentRepoImpl) AddStudent(student model.Student) (bool, error) {
	query := "INSERT INTO students (name) VALUES ($1) RETURNING id"
	err := r.studentDb.QueryRow(query, student.Name).Scan(&student.Id)
	if err != nil {
		log.Fatalf("Error creating student: %v", err)
		return false, err
	}

	return true, nil
}

func (r *StudentRepoImpl) AddGrade(reportCard model.ReportCard) (bool, error) {
	query := "INSERT INTO reports (subject, grade, student_id) VALUES ($1, $2, $3) RETURNING id"
	err := r.studentDb.QueryRow(query, reportCard.Subject, reportCard.Grade, reportCard.StudentId).Scan(&reportCard.Id)
	if err != nil {
		log.Fatalf("Error creating report: %v", err)
		return false, err
	}

	return true, nil
}

func (r *StudentRepoImpl) FindStudentById(id int) (model.Student, error) {
	var student model.Student

	query := "SELECT * FROM students WHERE id = $1"
	err := r.studentDb.QueryRow(query, id).Scan(&student.Id, &student.Name)
	if err != nil {
		log.Fatalf("Error finding student: %v", student)
		return model.Student{}, err
	}

	return student, nil
}

func (r *StudentRepoImpl) FindGradesByStudentId(id int) ([]model.ReportCard, error) {
	var reportCards []model.ReportCard

	rows, err := r.studentDb.Query("SELECT * FROM reports WHERE student_id = $1", id)
	if err != nil {
		log.Fatalf("Error finding student reports: %v", id)
		return reportCards, err
	}
	defer rows.Close()

	for rows.Next() {
		var reportCard model.ReportCard
		if err := rows.Scan(&reportCard.Subject, &reportCard.Grade, &reportCard.StudentId); err != nil {
			return reportCards, err
		}
		reportCards = append(reportCards, reportCard)
	}

	return reportCards, nil
}

func (r *StudentRepoImpl) FindGradesBySubject(subject string) ([]model.ReportCard, error) {
	var reportCards []model.ReportCard

	rows, err := r.studentDb.Query("SELECT * FROM reports WHERE subject = $1", subject)
	if err != nil {
		log.Fatalf("Error finding student reports: %v", subject)
		return reportCards, err
	}
	defer rows.Close()

	for rows.Next() {
		var reportCard model.ReportCard
		if err := rows.Scan(&reportCard.Subject, &reportCard.Grade, &reportCard.StudentId); err != nil {
			return reportCards, err
		}
		reportCards = append(reportCards, reportCard)
	}

	return reportCards, nil
}
