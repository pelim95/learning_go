## Assignment 1 — Student Grade Tracker

---

### Overview
Build a grade tracking system for a school that manages students and their grades across different subjects. The system should be built using TDD — tests written before implementation.

---

### Domain
You need two models — a **Student** and a **Grade**. Think about what fields make sense for each. Both must have a `String()` method that displays their information in a readable format. Student names longer than 20 characters must be truncated safely using runes.

---

### What to Build

**Repository layer**
An interface and an in-memory implementation that can store and retrieve students. Should support adding a student, finding by ID, and retrieving all students. Internally use a map for storage.

**Service layer**
An interface and implementation that handles all the business logic. Should support adding a grade to a student, calculating a student's average, finding the top student, and filtering students by a predicate function. The service must depend on the repository through constructor injection — never instantiate the repository directly inside the service.

---

### Business Rules
- Student name cannot be empty
- Grade score must be between 0 and 100
- Grade subject cannot be empty
- Average returns 0.0 if student has no grades
- Top student returns an error if no students exist
- Finding a student by an ID that doesn't exist returns an error

---

### TDD Requirements
- Tests must be written before implementation — Red, Green, Refactor
- Use table-driven tests with `t.Run` for every function
- Use testify for assertions
- Use `require` when subsequent test lines would break on nil, `assert` otherwise
- Service tests must use a mock repository — never the real in-memory one
- Every business rule must have at least one dedicated test case

---

### Concepts Expected
Your implementation must naturally use all of the following — pointers, structs, methods, interfaces, slices, maps, loops, range, error handling, first class functions, and stringers.

---

### Project Structure
Organise your code into at least three packages — domain, repository, and service.

---

### Submission
Paste all your files when done. I will review as a code reviewer — correctness, structure, idiomatic Go, test quality, and whether TDD was genuinely followed.