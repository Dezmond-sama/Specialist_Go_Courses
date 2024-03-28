// Нужно реализовать сервис для учета времени сотрудников на рабочем месте.
// ID, FIO, DEPARTMENT, POSITION

// In(Пришел)  -> Datetime
// Out(Ушел)   -> Datetime

// GET EMPLOYEE
// POST EMPLOYEE

// GET ALLTIME(day or month) выдать сколько времени провел на рабочем месте за указанный период времени.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Dezmond-sama/Specialist_Go_Courses/GO2/Lesson2/attendanceScheduleAPI/database"
	"github.com/gorilla/mux"
)

const (
	port = "8080"
)

var (
	db *database.Database
)

type ErrorMessage = struct {
	Message string `json:"message"`
}

type Duration = struct {
	Hours float64 `json:"hours"`
}

func GetIdHandler(w http.ResponseWriter, r *http.Request) (int64, bool) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid parameter: ", err)
		msg := ErrorMessage{Message: "ID property should be the int"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return 0, false
	}
	return int64(id), true
}
func FindEmployeeByIdHandler(w http.ResponseWriter, r *http.Request, id int64) (database.Employee, bool) {
	employee, found := db.FindEmployeeById(id)
	if !found {
		log.Printf("Employee with id %d not found.", id)
		msg := ErrorMessage{Message: fmt.Sprintf("Employee with id %d not found.", id)}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return employee, false
	}
	return employee, true
}

func GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {

	log.Println(db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.GetAllEmployees())
}

func GetEmployeeByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, found := GetIdHandler(w, r)
	if !found {
		return
	}
	employee, found := FindEmployeeByIdHandler(w, r, id)
	if !found {
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func GetEmployeeScheduleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	daysRaw := r.URL.Query()["days"]
	fromDate := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	if len(daysRaw) == 1 {
		days, err := strconv.Atoi(daysRaw[0])
		if err != nil {
			log.Println("Invalid query parameter: ", err)
			msg := ErrorMessage{Message: "days parameter should be the int"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(msg)
			return
		}
		fromDate = time.Now().AddDate(0, 0, -days)
	}
	id, found := GetIdHandler(w, r)
	if !found {
		return
	}
	employee, found := FindEmployeeByIdHandler(w, r, id)
	if !found {
		return
	}
	schedule, duration, err := db.EmployeeSchedule(id, fromDate, time.Now())
	if err != nil {
		log.Printf("Error while getting schedule data for employee with id = %d", id)
		msg := ErrorMessage{Message: fmt.Sprint(err.Error())}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
	type Schedule = struct {
		Schedule []database.WorkPeriod `json:"schedule"`
		Duration float64               `json:"total_duration"`
		Employee database.Employee     `json:"employee"`
	}
	json.NewEncoder(w).Encode(Schedule{Schedule: schedule, Duration: duration, Employee: employee})
}

func EmploeeWorkHandler(isIn bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, found := GetIdHandler(w, r)
		if !found {
			return
		}
		if !db.EmployeeExists(id) {
			return
		}
		var err error
		if isIn {
			_, err = db.EmployeeEnter(id)
		} else {
			_, err = db.EmployeeExit(id)
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := ErrorMessage{Message: err.Error()}
			json.NewEncoder(w).Encode(msg)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Bad request: ", err)
		msg := ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	var employee database.Employee
	json.Unmarshal(body, &employee)
	err = db.CreateEmployee(&employee)
	if err != nil {
		log.Println("DB Error: ", err)
		msg := ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(msg)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<style>table,th,td{border: 1px solid black; border-collapse: collapse; padding: 5px}</style>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")

	fmt.Fprintln(w, "<h1>API Employees</h1>")
	fmt.Fprintln(w, "<h2>Routes:</h2>")
	fmt.Fprintln(w, "<table>")
	fmt.Fprintln(w, "<tr><th>Method</th><th>Path</th><th>Description</th></tr>")
	fmt.Fprintln(w, "<tr><td>GET</td><td>/</td><td>Info</td></tr>")
	fmt.Fprintln(w, "<tr><td>GET</td><td>/employees</td><td>Employees list</td></tr>")
	fmt.Fprintln(w, "<tr><td>GET</td><td>/employees/{id}</td><td>Get employee by ID</td></tr>")
	fmt.Fprintln(w, "<tr><td>POST</td><td>/employees</td><td>Create new employee</td></tr>")
	fmt.Fprintln(w, "<tr><td>POST</td><td>/employees/{id}/start</td><td>Employee with ID starts working</td></tr>")
	fmt.Fprintln(w, "<tr><td>POST</td><td>/employees/{id}/end</td><td>Employee with ID finishes working</td></tr>")
	fmt.Fprintln(w, "<tr><td>GET</td><td>/employees/{id}/schedule?days=x</td><td>Get employee attendance for last x days. Total without x.</td></tr>")
	fmt.Fprintln(w, "</table>")
	fmt.Fprintln(w, "</body>")
}

func main() {
	var err error
	db, err = database.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Disconnect()
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/employees", GetAllEmployeesHandler).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", GetEmployeeByIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/employees", CreateEmployeeHandler).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}/start", EmploeeWorkHandler(true)).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}/end", EmploeeWorkHandler(false)).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}/schedule", GetEmployeeScheduleHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
