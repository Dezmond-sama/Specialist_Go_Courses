package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	calculator "github.com/Dezmond-sama/Specialist_Go_Courses/GO2/Lesson1/calculatorAPI/internal"
)

type NumberData struct {
	Value int `json:"value"`
}
type ResultData[ResultType int | float64] struct {
	First  int        `json:"first"`
	Second int        `json:"second"`
	Result ResultType `json:"result"`
}
type calculatorServer struct {
	calculator *calculator.Calculator
}

func NewCalculatorServer() *calculatorServer {
	return &calculatorServer{
		calculator: calculator.New(),
	}
}
func (cs *calculatorServer) infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path, r.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, cs.calculator.Info())
}
func (cs *calculatorServer) firstHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	res := NumberData{Value: cs.calculator.First()}

	json.NewEncoder(w).Encode(res)
}

func (cs *calculatorServer) secondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	res := NumberData{Value: cs.calculator.Second()}

	json.NewEncoder(w).Encode(res)
}

func (cs *calculatorServer) addHandler(w http.ResponseWriter, r *http.Request) {
	commonActionHandler(cs, w, r, cs.calculator.Add)
}
func (cs *calculatorServer) subHandler(w http.ResponseWriter, r *http.Request) {
	commonActionHandler(cs, w, r, cs.calculator.Sub)
}
func (cs *calculatorServer) mulHandler(w http.ResponseWriter, r *http.Request) {
	commonActionHandler(cs, w, r, cs.calculator.Mul)
}
func (cs *calculatorServer) divHandler(w http.ResponseWriter, r *http.Request) {
	commonActionHandler(cs, w, r, cs.calculator.Div)
}
func commonActionHandler[FuncResult int | float64](cs *calculatorServer, w http.ResponseWriter, r *http.Request, callback func(int, int) FuncResult) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	first := cs.calculator.First()
	second := cs.calculator.Second()
	result := callback(first, second)
	res := ResultData[FuncResult]{
		First:  first,
		Second: second,
		Result: result,
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	mux := http.NewServeMux()
	server := NewCalculatorServer()
	mux.HandleFunc("/info", server.infoHandler)
	mux.HandleFunc("/first", server.firstHandler)
	mux.HandleFunc("/second", server.secondHandler)
	mux.HandleFunc("/add", server.addHandler)
	mux.HandleFunc("/sub", server.subHandler)
	mux.HandleFunc("/mul", server.mulHandler)
	mux.HandleFunc("/div", server.divHandler)
	log.Fatal(http.ListenAndServe(":1234", mux))
}
