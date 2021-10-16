package tests

import (
	"github.com/fadlikadn/poc-ecommerce-api/api/models"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
)

func TestFindAllCustomers(t *testing.T) {
	err := RefreshTables()
	if err != nil {
		log.Fatal(err)
	}

	_, err = SeedCustomer()
	if err != nil {
		log.Fatal(err)
	}

	users, err := customerInstance.FindAllCustomers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {
	err := RefreshTables()
	if err != nil {
		log.Fatal(err)
	}
	newCustomer := models.Customer{
		Email: "test@gmail.com",
		Nickname: "test",
		Password: "password",
	}
	savedCustomer, err := newCustomer.AddCustomer(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newCustomer.ID, savedCustomer.ID)
	assert.Equal(t, newCustomer.Email, savedCustomer.Email)
	assert.Equal(t, newCustomer.Nickname, savedCustomer.Nickname)
}

func TestGetUserByID(t *testing.T) {
	err := RefreshTables()
	if err != nil {
		log.Fatal(err)
	}

	customers, err := SeedCustomer()
	if err != nil {
		log.Fatalf("cannot seed customers table: %v", err)
	}
	foundCustomer, err := customerInstance.FindCustomerByID(server.DB, customers[0].ID)
	if err != nil {
		t.Errorf("this is the error getting one customer: %v\n", err)
		return
	}
	assert.Equal(t, foundCustomer.ID, customers[0].ID)
	assert.Equal(t, foundCustomer.Email, customers[0].Email)
	assert.Equal(t, foundCustomer.Nickname, customers[0].Nickname)
}

func TestUpdateCustomer(t *testing.T) {
	err := RefreshTables()
	if err != nil {
		log.Fatal(err)
	}

	customers, err := SeedCustomer()
	if err != nil {
		log.Fatalf("Cannot seed customer: %v\n", err)
	}

	customerUpdate := models.Customer{
		Nickname: "modiUpdate",
		Email: "modiupdate@gmail.com",
		Password: "password",
	}
	updatedCustomers, err := customerUpdate.UpdateCustomer(server.DB, customers[0].ID)
	if err != nil {
		t.Errorf("this is the error updating the customer: %v\n", err)
		return
	}
	assert.Equal(t, updatedCustomers.ID, customerUpdate.ID)
	assert.Equal(t, updatedCustomers.Email, customerUpdate.Email)
	assert.Equal(t, updatedCustomers.Nickname, customerUpdate.Nickname)
}

func TestDeleteCustomer(t *testing.T) {
	err := RefreshTables()
	if err != nil {
		log.Fatal(err)
	}

	customers, err := SeedCustomer()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := customerInstance.DeleteCustomer(server.DB, customers[0].ID)
	if err != nil {
		t.Errorf("this is the error updating the customer: %v\n", err)
		return
	}
	// one shows that the record has been deleted or:
	// assert.Equal(t, int(isDeleted), 1)

	// Can be done this way too
	assert.Equal(t, isDeleted, int64(1))
}

