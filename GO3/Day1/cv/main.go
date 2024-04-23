package main

import (
	"html/template"
	"log"
	"net/http"
)

const PORT = "3000"

var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templates["index"] = template.Must(
		template.ParseFiles(
			"templates/layout.html",
			"templates/header.html",
			"templates/footer.html",
			"templates/index.html",
			"templates/jobs.html",
			"templates/skills.html",
		),
	)
	templates["jobs"] = template.Must(
		template.ParseFiles(
			"templates/layout.html",
			"templates/header.html",
			"templates/footer.html",
			"templates/jobs-body.html",
			"templates/jobs.html",
		),
	)
	templates["skills"] = template.Must(
		template.ParseFiles(
			"templates/layout.html",
			"templates/header.html",
			"templates/footer.html",
			"templates/skills-body.html",
			"templates/skills.html",
		),
	)
	templates["contacts"] = template.Must(
		template.ParseFiles(
			"templates/layout.html",
			"templates/header.html",
			"templates/footer.html",
			"templates/contacts.html",
		),
	)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates["index"].Execute(w,
		struct {
			About  About
			Jobs   []Job
			Skills []string
			Active string
		}{about, jobs, skills, "/"})
}
func skillsHandler(w http.ResponseWriter, r *http.Request) {
	templates["skills"].Execute(w,
		struct {
			About  About
			Skills []string
			Active string
		}{about, skills, "/skills"})
}
func jobsHandler(w http.ResponseWriter, r *http.Request) {
	templates["jobs"].Execute(w,
		struct {
			Jobs   []Job
			Active string
		}{jobs, "/jobs"})
}
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	templates["contacts"].Execute(w,
		struct {
			About  About
			Active string
		}{about, "/contacts"})
}
func main() {
	loadTemplates()

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/skills", skillsHandler)
	http.HandleFunc("/jobs", jobsHandler)
	http.HandleFunc("/contacts", contactsHandler)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
