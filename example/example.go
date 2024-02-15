package example

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Student represents a student entity.
type Student struct {
	ID   int // Assuming integer for simplicity, you can use UUID here as well
	Name string
	Age  int
	// Add other fields if needed
}

// Create inserts a new student into the database.
func (r *StudentRepository) Create(s *Student) error {
	_, err := r.db.Exec("INSERT INTO students (name, age) VALUES ($1, $2)", s.Name, s.Age)
	return err
}

// Update updates an existing student in the database.
func (r *StudentRepository) Update(id int, s *Student) error {
	_, err := r.db.Exec("UPDATE students SET name=$1, age=$2 WHERE id=$3", s.Name, s.Age, id)
	return err
}

// Delete removes a student from the database by ID.
func (r *StudentRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM students WHERE id=$1", id)
	return err
}

// FindByID retrieves a student from the database by ID.
func (r *StudentRepository) FindByID(id int) (*Student, error) {
	var s Student
	err := r.db.QueryRow("SELECT name, age FROM students WHERE id=$1", id).Scan(&s.Name, &s.Age)
	if err != nil {
		return nil, err
	}
	s.ID = id
	return &s, nil
}

// StudentRepository handles database operations for students.
type StudentRepository struct {
	db *sql.DB
}

// NewStudentRepository creates a new StudentRepository instance.
func NewStudentRepository(dataSourceName string) (*StudentRepository, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &StudentRepository{db: db}, nil
}

// Close closes the underlying database connection.
func (r *StudentRepository) Close() error {
	return r.db.Close()
}

// FindByName retrieves a list of students from the database by name.
func (r *StudentRepository) FindByName(name string) ([]*Student, error) {
	rows, err := r.db.Query("SELECT id, age FROM students WHERE name=$1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []*Student{}
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Age); err != nil {
			return nil, err
		}
		s.Name = name
		students = append(students, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

// FindByAge retrieves a list of students from the database by age.
func (r *StudentRepository) FindByAge(age int) ([]*Student, error) {
	rows, err := r.db.Query("SELECT id, name FROM students WHERE age=$1", age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []*Student{}
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, err
		}
		s.Age = age
		students = append(students, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

func main2() {
	// Replace the dataSourceName with your PostgreSQL connection string
	dataSourceName := "user=yourusername password=yourpassword dbname=yourdbname sslmode=disable"
	repo, err := NewStudentRepository(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	// Example usage
	student := &Student{Name: "John Doe", Age: 20}
	if err := repo.Create(student); err != nil {
		log.Fatal(err)
	}

	// Querying data
	studentsByName, err := repo.FindByName("John Doe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Students with name John Doe:", studentsByName)

	studentsByAge, err := repo.FindByAge(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Students with age 20:", studentsByAge)
}
