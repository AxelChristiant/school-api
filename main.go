package main

import (
	"SCHOOL-API/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var data = map[string]models.Student{
	"E001": models.Student{"E001", "ethan", 21},
	"W001": models.Student{"W001", "wick", 22},
	"B001": models.Student{"B001", "bourne", 23},
	"B002": models.Student{"B002", "bond", 23},
}

func showAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result, err = json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
	return
}

func addNewStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent models.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data[newStudent.ID] = newStudent
	w.Write([]byte("Succesfully added the data!"))
	return
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	var editedStudent models.Student
	id := strings.TrimPrefix(r.URL.Path, "/edit/student/")
	if _, ok := data[id]; ok {
		err := json.NewDecoder(r.Body).Decode(&editedStudent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data[id] = editedStudent
		result := fmt.Sprintf("Succesfully edited the data with id %s!", id)
		w.Write([]byte(result))
		return
	}
	http.Error(w, fmt.Sprintf("id : %s, not found!", id), http.StatusBadRequest)

	return

}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/delete/student/")
	delete(data, id)
	result := fmt.Sprintf("Succesfully deleted a data with id %s!", id)
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Haii"))
	})
	http.HandleFunc("/student", showAllStudents)
	http.HandleFunc("/addstudent", addNewStudent)
	http.HandleFunc("/edit/student/", updateStudent)
	http.HandleFunc("/delete/student/", deleteStudent)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)

}
