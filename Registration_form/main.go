package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Employee struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var (
	employees []Employee
	mu        sync.Mutex
	fileName  = "employees.json"
)

func main() {
	// Load existing employees from file
	loadEmployeesFromFile()

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/register", registerEmployee)
	http.HandleFunc("/employees", getEmployees)
	http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

func registerEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	mu.Lock()
	employees = append(employees, emp)
	saveEmployeesToFile()
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func loadEmployeesFromFile() {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &employees); err != nil {
		panic(err)
	}
}

func saveEmployeesToFile() {
	data, err := json.MarshalIndent(employees, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		panic(err)
	}
}
