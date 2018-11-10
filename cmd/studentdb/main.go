package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Macpon7/imt2681_cloud/studentdb"
	"os"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 {
		status := 400
		http.Error(w, http.StatusText(status), status)
		return
	}
	name := parts[2]
	fmt.Fprintln(w, parts)
	fmt.Fprintf(w, "Hello %s %s!\n", name, parts[3])
}

// -----------------

func main() {
	// Using in-memory storage
	studentdb.Global_db = &studentdb.StudentsDB{}

	// Using MongoDB based storage
	/* studentdb.Global_db = &studentdb.StudentsMongoDB{
		"mongodb://localhost",
		"studentsDB",
		"students",
	}*/

	studentdb.Global_db.Init()

	port := os.Getenv("PORT")
	http.HandleFunc("/hello/", handlerHello)
	http.HandleFunc("/student/", studentdb.HandlerStudent)
	http.ListenAndServe(":"+port, nil)
}
