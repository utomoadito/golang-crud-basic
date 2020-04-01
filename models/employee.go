package models

// Employees export functions
type Employees struct {
	ID    int64  `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	City  string `form:"city" json:"city"`
	Phone string `form:"phone" json:"phone"`
}

// EmployeeResponse export functions
type EmployeeResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Employees
}
