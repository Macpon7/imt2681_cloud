package main

import "testing"

func Test_addStudent(t *testing.T) {

	// db *StudentDB
	db := &StudentsDB{}

	db.Add(Student{"Tom", 21})
	if db.Count() != 1 {
		t.Error("Wrong student count")
	}
	s := db.Get(0)
	if s.Name != "Tom" {
		t.Error("Student Tom was not added.")
	}
}

func Test_multipleStudents(t *testing.T) {
	testData := []Student{
		{"Bob", 21},
		{"Alice", 20},
		{"Alice", 24},
	}

	db := StudentsDB{}
	for _, s := range testData {
		db.Add(s)
	}

	if db.Count() != len(testData) {
		t.Error("Wrong number of students")
	}

	for i := range db.students {
		if db.Get(i).Name != testData[i].Name {
			t.Error("Wrong name")
		}
		if db.Get(i).Age != testData[i].Age {
			t.Error("Wrong name")
		}
	}

}