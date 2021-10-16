package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for corses - file
type Course struct {
	CourseId    string  `json: courseid`
	CourseName  string  `json: coursename`
	CoursePrice float64 `json: price`
	Author      *Author `json: author`
}

type Author struct {
	Fullname string `json: "fullname"`
	Website  string `json: "website`
}

// fake db
var courses []Course

// middleware
func IsEmpty(c *Course) bool {
	// return c.CourseId == "" && c.CourseName == ""

	return c.CourseName == ""
}

func main() {

}

// controllers -file
// serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to APIs in go lang</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with ", params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return

}
