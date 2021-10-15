package controllers

import (
	"encoding/json"
	"github.com/fadlikadn/poc-ecommerce-api/api/models"
	"github.com/fadlikadn/poc-ecommerce-api/api/responses"
	"github.com/fadlikadn/poc-ecommerce-api/api/utils/formaterror"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	transaction := models.Transaction{}
	err = json.Unmarshal(body, &transaction)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	transaction.Prepare()
	err = transaction.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	transactionCreated, err := transaction.AddNewTransaction(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusCreated, transactionCreated)
}

func (server *Server) GetTransactionByCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	transaction := models.Transaction{}

	transactions, err := transaction.FindAllTransactionsByCustomer(server.DB, cid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactions)
}

func (server *Server) AddTransactionDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	transaction := models.Transaction{}
	transactionDetail := models.TransactionDetail{}
	warehouseItem := models.WarehouseItem{}

	// Check if the transaction exist
	err = server.DB.Debug().Model(models.Transaction{}).Where("id = ?", tid).Take(&transaction).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, formaterror.FormatError("Transaction not found"))
		return
	}

	// Read from data posted (transaction detail)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//start processing the request data
	err = json.Unmarshal(body, &transactionDetail)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	totalPurchased, err := transactionDetail.GetWarehouseItemTotalPurchased(server.DB, transactionDetail.WarehouseItemID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, formaterror.FormatError("Error fetch warehouse item total purchased"))
		return
	}
	warehouseFound, err := warehouseItem.FindWarehouseByID(server.DB, transactionDetail.WarehouseItemID)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, formaterror.FormatError("Error get warehouse by id"))
		return
	}

	if warehouseFound.Quantity < (totalPurchased + transactionDetail.Quantity) {
		// TODO: validation to prevent out of stock happened when purchase new item
		responses.ERROR(w, http.StatusUnprocessableEntity, formaterror.FormatError("Item out of stock"))
		return
	}

	transactionDetail.Prepare()
	transactionDetailCreated, err := transactionDetail.AddNewTransactionDetail(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, formaterror.FormatError(err.Error()))
		return
	}
	responses.JSON(w, http.StatusCreated, transactionDetailCreated)
}
