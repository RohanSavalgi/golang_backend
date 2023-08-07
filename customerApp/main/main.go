package main

import (
	"assigments/domain"
	"assigments/mapstore"

	"fmt"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (c CustomerController) Add(customer domain.Customer) {
	err := c.store.CreateCustomer(customer)
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println("Customer has been created successfully")
}

func (c CustomerController) Remove(id int) {
	err := c.store.DeleteCustomer(id) 
	if err != nil{
		fmt.Println("error : ", err)
		return
	}
	fmt.Println("Customer delete successfully")
}

func (c CustomerController) GetAll() {
	allCust, err := c.store.GetAllCustomer() 
	if err != nil {
		fmt.Println("error : ", err)
	}
	fmt.Println(allCust)
}

func main() {
	controller := CustomerController{ mapstore.NewModuleStore() }

	customer1 := domain.Customer{
		Id: 1,
		Name: "rohan",
		Email: "rohan@gmail.com",
	}

	controller.Add(customer1)
	controller.GetAll()

	controller.Remove(1)
}



