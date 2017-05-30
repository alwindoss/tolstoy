package server

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get Courses!")
}
