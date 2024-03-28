package database

import (
	"database/sql"
	"errors"
	"math/rand"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

type Employee = struct {
	EmployeeId int64  `json:"employee_id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Position   string `json:"position"`
}

type Attendance = struct {
	AttendanceId int64     `json:"attendance_id"`
	EmployeeId   int64     `json:"employee_id"`
	Time         time.Time `json:"time"`
	IsIn         bool      `json:"is_in"` // true -> start, false -> end
}
type WorkPeriod = struct {
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
	Duration float64   `json:"duration"`
}
type Database struct {
	db *sql.DB
}

func New() (*Database, error) {
	newDB := false
	if _, err := os.Stat("./employees.db"); errors.Is(err, os.ErrNotExist) {
		newDB = true
	}
	db, err := sql.Open("sqlite", "./employees.db")
	if err != nil {
		return nil, err
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS employees (
		employee_id INTEGER NOT NULL PRIMARY KEY, 
		name TEXT, 
		department TEXT, 
		position TEXT
	);`
	if _, err = db.Exec(sqlStmt); err != nil {
		return nil, err
	}
	// AttendanceId int       `json:"attendance_id"`
	// EmployeeId   int       `json:"employee_id"`
	// Time         time.Time `json:"time"`
	// IsIn         bool      `json:"is_in"` // true -> start, false -> end

	// Execute the SQL statement
	sqlStmt = `
	CREATE TABLE IF NOT EXISTS attendances (
		attendance_id INTEGER NOT NULL PRIMARY KEY, 
		employee_id INTEGER NOT NULL, 
		time DATETIME, 
		is_in BOOLEAN,
		CONSTRAINT fk_employee  
		FOREIGN KEY (employee_id)  
		REFERENCES employees (employee_id)  
		ON UPDATE CASCADE 
		ON DELETE CASCADE
	);
	`

	if _, err = db.Exec(sqlStmt); err != nil {
		return nil, err
	}
	resDB := &Database{db: db}
	if newDB {
		resDB.fillMokData()
	}
	return resDB, nil
}

func (db *Database) Disconnect() {
	db.db.Close()
}

func (db *Database) CreateEmployee(employee *Employee) error {
	sqlStmt := `
	INSERT INTO employees (name, department, position)
	VALUES (?, ?, ?);`

	stmt, err := db.db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(employee.Name, employee.Department, employee.Position)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	employee.EmployeeId = id
	return nil
}
func (db *Database) EmployeeExists(id int64) bool {
	sqlStmt := `
	select count(*) from employees where employee_id = ?;
	`

	stmt, err := db.db.Prepare(sqlStmt)
	if err != nil {
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return false
	}
	defer rows.Close()

	if rows.Next() {
		var cnt int
		err = rows.Scan(&cnt)
		if err != nil {
			return false
		}
		return cnt > 0
	}

	return false
}

func (db *Database) FindEmployeeById(id int64) (Employee, bool) {
	var employee Employee
	sqlStmt := `
	select employee_id, name, department, position from employees where employee_id = ?;
	`

	stmt, err := db.db.Prepare(sqlStmt)
	if err != nil {
		return employee, false
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return employee, false
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&employee.EmployeeId, &employee.Name, &employee.Department, &employee.Position)
		if err != nil {
			return employee, false
		}
		return employee, true
	}

	return employee, false
}

func (db *Database) GetAllEmployees() []Employee {
	sqlStmt := `
	select employee_id, name, department, position from employees;
	`

	stmt, err := db.db.Prepare(sqlStmt)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil
	}
	defer rows.Close()
	var res []Employee = make([]Employee, 0, 100)
	for rows.Next() {
		var employee Employee
		err = rows.Scan(&employee.EmployeeId, &employee.Name, &employee.Department, &employee.Position)
		if err == nil {
			res = append(res, employee)
		}
	}

	return res
}

func (db *Database) fillEmployeeAttendance(id int64, tm time.Time, isIn bool) (int64, error) {
	sqlStmt := `
	INSERT INTO attendances (employee_id, time, is_in)
	VALUES (?, ?, ?);`

	stmt, err := db.db.Prepare(sqlStmt)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(id, tm, isIn)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (db *Database) EmployeeEnter(id int64) (int64, error) {
	return db.fillEmployeeAttendance(id, time.Now(), true)
}

func (db *Database) EmployeeExit(id int64) (int64, error) {
	return db.fillEmployeeAttendance(id, time.Now(), false)
}

func (db *Database) EmployeeSchedule(id int64, fromDate time.Time, toDate time.Time) ([]WorkPeriod, float64, error) {
	var workPeriods []WorkPeriod = make([]WorkPeriod, 0, 100)
	sqlStmt := `
	select time, is_in from attendances 
	where employee_id = ? and time > ? and time < ?
	order by time;
	`

	stmt, err := db.db.Prepare(sqlStmt)
	if err != nil {
		return workPeriods, 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id, fromDate, toDate)
	if err != nil {
		return workPeriods, 0, err
	}
	defer rows.Close()

	started := false
	var totalDuration float64
	var startedTime time.Time
	for rows.Next() {
		var tm time.Time
		var isIn bool
		err = rows.Scan(&tm, &isIn)
		if err == nil {
			if started && !isIn {
				dur := tm.Sub(startedTime).Hours()
				totalDuration = totalDuration + dur
				workPeriods = append(workPeriods, WorkPeriod{From: startedTime, To: tm, Duration: dur})
				started = false
			} else if !started && isIn {
				started = true
				startedTime = tm
			}
		}
	}
	if started {
		dur := toDate.Sub(startedTime).Hours()
		totalDuration = totalDuration + dur
		workPeriods = append(workPeriods, WorkPeriod{From: startedTime, To: toDate, Duration: dur})
	}
	return workPeriods, totalDuration, rows.Err()
}

// Data for testing
func (db *Database) fillMokData() {
	employees := []*Employee{
		{Name: "John Smith", Department: "Company", Position: "CEO"},
		{Name: "Adam Ministrator", Department: "Human Resources", Position: "CHRO"},
		{Name: "John Doe", Department: "Human Resources", Position: "HR Manager"},
		{Name: "Yolanda Bishop", Department: "Human Resources", Position: "HR Manager"},
		{Name: "Allen Hansen", Department: "Human Resources", Position: "HR Manager"},
		{Name: "Kay Walker", Department: "Accounting", Position: "General Manager"},
		{Name: "Stella Ruiz", Department: "Accounting", Position: "Accountant"},
		{Name: "Soham Flores", Department: "Accounting", Position: "Accountant"},
		{Name: "Leslie Bates", Department: "Accounting", Position: "Accountant"},
		{Name: "Jayden Jackson", Department: "Accounting", Position: "Accountant"},
		{Name: "Leon Gutierrez", Department: "Information Technology", Position: "CTO"},
		{Name: "Ryan Kuhn", Department: "Information Technology", Position: "Team Leader"},
		{Name: "Brent King", Department: "Information Technology", Position: "System Architector"},
		{Name: "Leona Bell", Department: "Information Technology", Position: "Senior Developer"},
		{Name: "Caleb Roberts", Department: "Information Technology", Position: "Senior Developer"},
		{Name: "William Ferguson", Department: "Information Technology", Position: "Middle Developer"},
		{Name: "Saginaw Delaware", Department: "Information Technology", Position: "Middle Developer"},
		{Name: "Adrian Gibson", Department: "Information Technology", Position: "Middle Developer"},
		{Name: "Danny Garza", Department: "Information Technology", Position: "Junior Developer"},
		{Name: "Frank King", Department: "Information Technology", Position: "Junior Developer"},
		{Name: "Yolanda Neal", Department: "Information Technology", Position: "Junior Developer"},
		{Name: "Johnny West", Department: "Information Technology", Position: "Junior Developer"},
		{Name: "Teresa Hernandez", Department: "Information Technology", Position: "QA Manager"},
		{Name: "Evelyn Wallace", Department: "Information Technology", Position: "QA"},
		{Name: "Ethel Hopkins", Department: "Information Technology", Position: "QA"},
		{Name: "Chad Ferguson", Department: "Information Technology", Position: "QA"},
		{Name: "Mario Griffin", Department: "Information Technology", Position: "QA"},
		{Name: "Alan Perez", Department: "Information Technology", Position: "Pen Tester"},
		{Name: "Dolores Evans", Department: "Information Technology", Position: "Pen Tester"},
		{Name: "Anita Bradley", Department: "Information Technology", Position: "Pen Tester"},
	}
	for _, employee := range employees {
		db.CreateEmployee(employee)
	}
	for i := 1; int64(i) <= employees[len(employees)-1].EmployeeId; i++ {
		db.fillMokEmployeeScheduleData(int64(i), time.Now().AddDate(0, -rand.Intn(6), -rand.Intn(30)), time.Now())
	}
}

func (db *Database) fillMokEmployeeScheduleData(id int64, fromDate time.Time, toDate time.Time) {
	for {
		if rand.Intn(30) > 25 {
			continue
		}
		start := time.Date(fromDate.Year(), fromDate.Month(), fromDate.Day(), 8+rand.Intn(3), rand.Intn(60), rand.Intn(60), 0, time.Local)
		end := start.Add(time.Hour * time.Duration(6+rand.Intn(3))).Add(time.Minute * time.Duration(rand.Intn(60)))
		db.fillEmployeeAttendance(id, start, true)
		db.fillEmployeeAttendance(id, end, false)
		fromDate = fromDate.Add(time.Hour * 24)
		if fromDate.After(toDate) {
			break
		}
	}

}
