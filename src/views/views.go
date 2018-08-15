package views

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
	"log"
	"fmt"
	"github.com/connelevalsam/GoWebDev/GoLearning/SchoolWork/src/db"
)

var (
	err        error
	templ      *template.Template
	adtempl    *template.Template
	ustempl    *template.Template
	uName      string
	ID         int
	isLoggedin bool
	file       io.Writer
)

type PageData struct {
	Title    string
	IsAuth   int
	Id       int
	Username string
	Password string
	AdminD   bool
}

func InitFunc() {
	isLoggedin = false
	//admin landing page
	adtempl, err = templ.ParseGlob("templates/admin/*.html")
	if err != nil {
		checkErr(err)
	}
	//user landing page
	ustempl, err = templ.ParseGlob("templates/user/*.html")
	if err != nil {
		checkErr(err)
	}
	//landing page
	templ, err = templ.ParseGlob("templates/*.html")
	if err != nil {
		checkErr(err)
	}
}

func IndexHandle(res http.ResponseWriter, req *http.Request) {
	/*
	*			first check
	*			section
	 */
	cookie, err := req.Cookie("bangalee")
	title := "Welcome || CB Uni"
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "bangalee",
			Value: "0",
		}
	}

	if cookie.Value == strconv.Itoa(1) {
		if isLoggedin == true {
			var pd = PageData{Title: title, Username: uName}
			err = adtempl.ExecuteTemplate(res, "index.html", pd)
			if err != nil {
				log.SetOutput(file)
				http.Error(res, "there was an error", http.StatusInternalServerError)
				return
			}
			fmt.Println("Admin successfully logged in")
		} else {
			http.Redirect(res, req, "/", 301)
		}
	} else if cookie.Value == strconv.Itoa(2) {
		fmt.Println("user successfully logged in")
	} else {
		isLoggedin = false
		pd := PageData{Title: title}
		err = templ.ExecuteTemplate(res, "index.html", pd)
		if err != nil {
			log.SetOutput(file)
		}
	}

}

//
func HandleLogin(res http.ResponseWriter, req *http.Request) {
	/*
	*			login
	*			section
	 */
	title := "Login"
	if req.Method != "POST" {
		var pd = PageData{Title: title}
		err = templ.ExecuteTemplate(res, "login.html", pd)
		if err != nil {
			log.SetOutput(file)
		}
		return
	}
	//get username and password
	username := req.FormValue("username")
	fpassword := req.FormValue("password")
	var pword string
	var isAdmin bool
	var cookie *http.Cookie
	//
	con := db.Conn
	// query the db and check if the username entered matches any in the db
	rows, err := con.Query("SELECT password,isAdmin FROM users WHERE username = $1", username)
	if err != nil {
		log.SetOutput(file)
		http.Error(res, fmt.Sprintf("error in login or username: %v", err), http.StatusInternalServerError)
		return
	}
	//
	//loop through the db to check required details
	for rows.Next() {
		err = rows.Scan(&pword, &isAdmin)
		if err != nil {
			log.SetOutput(file)
			http.Error(res, "there was an error in login other fields", http.StatusInternalServerError)
			return
		}
		//is it is admin
		if isAdmin == true {
			if fpassword != pword {
				log.SetOutput(file)
				fmt.Println("invalid")
				http.Redirect(res, req, "/login", 301)
				return
			}
			//
			uName = username
			cookie = &http.Cookie{
				Name:  "bangalee",
				Value: "1",
			}
			isLoggedin = true
			http.SetCookie(res, cookie)
			http.Redirect(res, req, "/", 302)
			return
		} else {
			if fpassword != pword {
				log.SetOutput(file)
				fmt.Println("invalid")
				http.Redirect(res, req, "/login", 301)
				return
			}
			//
			uName = username
			cookie = &http.Cookie{
				Name:  "bangalee",
				Value: "2",
			}
			isLoggedin = true

		}
	}
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", 302)
	return
}

//handle logout
func Logout(res http.ResponseWriter, req *http.Request) {
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	http.SetCookie(res, &http.Cookie{
		Name:   "bangalee",
		MaxAge: -1,
	})
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

//about handler
func AboutHandle(res http.ResponseWriter, req *http.Request) {
	title := "About || CB Uni"
	pd := PageData{Title: title}
	templ.ExecuteTemplate(res, "about.html", pd)
}

//check error
func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

//================================ admin ====================
func AdminHome(res http.ResponseWriter, req *http.Request) {
	title := "Admin || Home"
	var pd = PageData{Title: title}
	adtempl.ExecuteTemplate(res, "index.html", pd)
}
func AdminLecturer(res http.ResponseWriter, req *http.Request) {
	title := "Admin || Lecturers"
	var pd = PageData{Title: title}
	adtempl.ExecuteTemplate(res, "lecturer.html", pd)
}
func AdminStudent(res http.ResponseWriter, req *http.Request) {
	title := "Admin || Students"
	var pd = PageData{Title: title}
	adtempl.ExecuteTemplate(res, "student.html", pd)
}
func AdminTransfer(res http.ResponseWriter, req *http.Request) {
	title := "Admin || Transfer"
	var pd = PageData{Title: title}
	adtempl.ExecuteTemplate(res, "transfer.html", pd)
}

