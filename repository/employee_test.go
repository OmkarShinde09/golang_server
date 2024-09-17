package repository

import (
	"api/model"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://admin:OGS100902hp@cluster0.wal6d.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		panic(err)
	}
	log.Println("Connected to MongoDB!")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	
	if err != nil {
		log.Fatal("Ping failed", err)
	}
	
	log.Println("Ping success")


	return mongoTestClient
} 

func TestMongoOperations(t *testing.T){
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	//dummy data
	emp1 := uuid.New().String()
	emp2 := uuid.New().String()

	//connect to collection
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{
		MongoCollection: coll,
	}

	//Insert One Employee
	t.Run("Insert Employee 1", func(t *testing.T){
		emp := model.Employee{
			Name: "John",
			EmployeeId: emp1,
			Department: "Engineering",
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("Insert 1 successfully", result)
	})

	t.Run("Insert Employee 1", func(t *testing.T){
		emp := model.Employee{
			Name: "Steve",
			EmployeeId: emp2,
			Department: "Physics",
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("Insert 1 successfully", result)
	})

	//Get Employee 1 Data
	t.Run("Get Employee 1", func(t *testing.T){
		result, err := empRepo.FindEmployeeByID(emp1)

		if err != nil {
			t.Fatal(err)
		}
		t.Log("Get Employee 1 successfully", result.Name)
	})

	//Get all Employees
	t.Run("Get All Employee", func(t *testing.T){
		result, err := empRepo.FindAllEMployee()
		if err != nil {
			t.Fatal("Got operation falied",err)
		}
		t.Log("Get All Employee successfully", result)
	})

	//Update Employee 1 data
	t.Run("Update Employee 1", func(t *testing.T){
		emp := model.Employee{
			Name: "Tony",
			EmployeeId: emp1,
			Department: "Physics",
		}

		result, err := empRepo.UpdateEmployeeByID(emp1, &emp)

		if err != nil {
			t.Fatal("Update operation failed", err)
		}
		t.Log("update count", result)

		//Get Employee 1 Data
		t.Run("Get Employee 1", func(t *testing.T){
			result, err := empRepo.FindEmployeeByID(emp1)

			if err != nil {
				t.Fatal(err)
			}
			t.Log("Get Employee 1 successfully", result.Name)
		})

		//Delete Employee 1 data
		t.Run("Delete Employee 1", func(t *testing.T){
			result, err := empRepo.DeleteEmployeeByID(emp1)

			if err != nil {
				t.Fatal("Delete operation failed", err)
			}
			t.Log("delete count", result)
		})

		//Get all Employee Data after delete
		t.Run("Get all employee after delete", func(t *testing.T) {
			result, err := empRepo.FindAllEMployee()

			if err != nil {
				t.Fatal("Got operation falied",err)
			}
			t.Log("Employees", result)
		})

		//Delete all employees
		t.Run("Delete all employees for cleanup", func(t *testing.T) {
			result, err := empRepo.DeleteAllEmployee()

			if err != nil {
				t.Fatal("Delete operation failed", err)
			}
			t.Log("delete count", result) 
		})
	})
}
