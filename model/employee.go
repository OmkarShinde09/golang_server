package model

type Employee struct{
	EmployeeId string `json:"employee_id,omitempty" bson:"employee_id"`
	Name string	`json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`//Because mongodb does not understand json hence bson tag is added.
}