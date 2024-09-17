package usecase

import (
	"api/model"
	"api/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct{
	MongoCollection *mongo.Collection
}

type Response struct{
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}


func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return 
	}

	//assign new id
	emp.EmployeeId = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	//insert employee
	insertID, err := repo.InsertEmployee(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insert error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp.EmployeeId
	w.WriteHeader(http.StatusOK)

	log.Println("Employee inserted with id", insertID, emp)
}

func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get employee id
	empID := mux.Vars(r)["id"]
	log.Println("Employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Get employee error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEMployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Get employee error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get employee id
	empID := mux.Vars(r)["id"]
	log.Println("Employee id", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid employee id", empID)
		res.Error = "Invalid employee id"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return 
	}

	emp.EmployeeId = empID

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Update employee error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get employee id
	empID := mux.Vars(r)["id"]
	log.Println("Employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Delete employee error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)  
}

func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Delete employee error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK) 
}
