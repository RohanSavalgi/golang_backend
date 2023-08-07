package mapstore

import (
	"assigments/domain"

	"errors"
)

type ModuleStore struct {
	store map[int]domain.Customer
}

func NewModuleStore() *ModuleStore {
	return &ModuleStore{ make(map[int]domain.Customer) }
}

func (moduleStore ModuleStore) CreateCustomer(customer domain.Customer) error {
	if _, item := moduleStore.store[customer.Id]; item {
		return errors.New("Customer already present")
	}

	moduleStore.store[customer.Id] = customer
	return nil
}

func (moduleStore ModuleStore) DeleteCustomer(id int) error {
	if _, item := moduleStore.store[id]; item {
		return errors.New("These is no customer to delete")
	}
	
	delete(moduleStore.store, id)
	return nil
}

func (moduleStore ModuleStore) GetAllCustomer() ([]domain.Customer, error) {
	var allCustomer []domain.Customer
	for _,item := range moduleStore.store {
		allCustomer = append(allCustomer, item)
	}
	return allCustomer, nil
}

func (moduleStore ModuleStore) GetCustomerById(id int) (domain.Customer, error) {
	if _, item := moduleStore.store[id]; item {
		return domain.Customer{}, errors.New("No Customer is found")
	}
	return moduleStore.store[id], nil
}

func (moduleStore ModuleStore) UpdateCustomer(id int, updateData domain.Customer) error {
	if _, item := moduleStore.store[id]; item {
		return errors.New("There is customer at all")
	}
	moduleStore.store[id] = updateData
	return nil
}
























// package mapstore

// import (
// 	"assigments/domain"

// 	"errors"
// )

// type MapStore struct {
// 	store map[int]domain.Customer
// }


// func NewMapStore() *MapStore {
// 	return &MapStore{ store: make(map[int]domain.Customer) }
// }

// func (mapStore *MapStore) CreateCustomer(customer domain.Customer) error {
// 	if _, cust := mapStore.store[customer.Id]; cust {
// 		return errors.New("Customer already exists")
// 	}

// 	mapStore.store[customer.Id] = customer 	
// 	return nil
// }