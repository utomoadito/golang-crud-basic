package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"crud_tutor/config"
	"crud_tutor/models"
)

// AllEmployees export functions
func AllEmployees(w http.ResponseWriter, r *http.Request) {
	var employees models.Employees
	var arrEmployee []models.Employees
	var response models.EmployeeResponse

	db, err := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id,name,city,phone FROM employees")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&employees.ID, &employees.Name, &employees.City, &employees.Phone); err != nil {
			log.Fatal(err.Error())

		} else {
			arrEmployee = append(arrEmployee, employees)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrEmployee

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddEmployees export functions
func AddEmployees(w http.ResponseWriter, r *http.Request) {
	var response models.EmployeeResponse
	var arrEmployee []models.Employees

	//Untuk ambil data dari json body request
	decoder := json.NewDecoder(r.Body)
	var payload models.Employees
	/* payload := struct {
		Name  string `json:"name"`
		City  string `json:"city"`
		Phone string `json:"phone"`
	}{} */
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	// Mulai koneksi ke DB dan INSERT data
	db, err := config.Connect()
	defer db.Close()

	rows, err := db.Exec("INSERT INTO employees values (?, ?, ?, ?)", "", payload.Name, payload.City, payload.Phone)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	id, err := rows.LastInsertId()
	payload.ID = id
	fmt.Println("insert success!")

	response.Status = 1
	response.Message = "Success"
	response.Data = append(arrEmployee, payload)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
