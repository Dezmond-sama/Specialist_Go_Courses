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

	"github.com/gorilla/mux"
)

const (
	port = "8080"
)

var (
	employees        []Employee
	attendances      []Attendance
	nextEmployeeId   = 1
	nextAttendanceId = 1
)

type ErrorMessage = struct {
	Message string `json:"message"`
}

type Duration = struct {
	Hours float64 `json:"hours"`
}

type Employee = struct {
	EmployeeId int    `json:"employee_id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Position   string `json:"position"`
}

type Attendance = struct {
	AttendanceId int       `json:"attendance_id"`
	EmployeeId   int       `json:"employee_id"`
	Time         time.Time `json:"time"`
	IsIn         bool      `json:"is_in"` // true -> start, false -> end
}

func FindEmployeeById(id int) (Employee, bool) {
	for _, p := range employees {
		if p.EmployeeId == id {
			return p, true
		}
	}
	return Employee{}, false
}
func GetIdHandler(w http.ResponseWriter, r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Invalid parameter: ", err)
		msg := ErrorMessage{Message: "ID property should be the int"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return 0, false
	}
	return id, true
}
func FindEmployeeByIdHandler(w http.ResponseWriter, r *http.Request, id int) (Employee, bool) {
	employee, found := FindEmployeeById(id)
	if !found {
		log.Printf("Employee with id %d not found.", id)
		msg := ErrorMessage{Message: fmt.Sprintf("Employee with id %d not found.", id)}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return Employee{}, false
	}
	return employee, true
}
func GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
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
func GetEmployeeSchduleHandler(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	id, found := GetIdHandler(w, r)
	if !found {
		return
	}
	if _, found = FindEmployeeByIdHandler(w, r, id); !found {
		return
	}
	duration := 0.0
	started := false
	for _, att := range attendances {
		if att.EmployeeId != id {
			continue
		}
		if att.Time.Before(fromDate) {
			continue
		}
		if att.IsIn {
			started = true
		} else if started {
			duration = duration + att.Time.Sub(fromDate).Seconds()
			started = false
		}
		fromDate = att.Time
	}
	if started {
		duration = duration + time.Since(fromDate).Seconds()
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Duration{Hours: duration / 60 / 60})
}
func EmploeeWorkHandler(isIn bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, found := GetIdHandler(w, r)
		if !found {
			return
		}
		if _, found = FindEmployeeByIdHandler(w, r, id); !found {
			return
		}
		attendance := Attendance{AttendanceId: nextAttendanceId, EmployeeId: id, Time: time.Now(), IsIn: isIn}
		nextAttendanceId++
		attendances = append(attendances, attendance)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(attendance)
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
	var employee Employee
	json.Unmarshal(body, &employee)
	employee.EmployeeId = nextEmployeeId
	nextEmployeeId++
	employees = append(employees, employee)

	json.NewEncoder(w).Encode(employee)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", GetAllEmployeesHandler).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", GetEmployeeByIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/employees", CreateEmployeeHandler).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}/start", EmploeeWorkHandler(true)).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}/end", EmploeeWorkHandler(false)).Methods(http.MethodPost)
	router.HandleFunc("/employees/{id}/schedule", GetEmployeeSchduleHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
