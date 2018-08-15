package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"github.com/connelevalsam/GoWebDev/GoLearning/SchoolWork/src/views"

	"github.com/connelevalsam/GoWebDev/GoLearning/SchoolWork/src/db"
)

var (
	err        error
	file       io.Writer
)



func init() {
	views.InitFunc()
	//database connection
	db.DBConnect()
}

func main() {
	http.HandleFunc("/", views.IndexHandle)
	http.HandleFunc("/about", views.AboutHandle)
	http.HandleFunc("/login", views.HandleLogin)
	http.HandleFunc("/logout", views.Logout)
	http.HandleFunc("/admin", views.AdminHome)
	http.HandleFunc("/lecturer", views.AdminLecturer)
	http.HandleFunc("/student", views.AdminStudent)
	http.HandleFunc("/transfer", views.AdminTransfer)
	//log
	file, err = os.OpenFile("logfile.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		e := err.Error()
		file.Write([]byte(e))
	}

	log.SetOutput(file)

	http.Handle("/assets/", http.FileServer(http.Dir("."))) //serve other files in assets dir
	fmt.Println("server running on port :9002")
	http.ListenAndServe(":9002", nil)
}



