package controllers

import (
	"fmt"
	"github.com/fadlikadn/poc-ecommerce-api/api/controllers"
	"github.com/fadlikadn/poc-ecommerce-api/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var server = controllers.Server{}
var customerInstance = models.Customer{}
var supplierItemInstance = models.SupplierItem{}
var warehouseItemInstance = models.WarehouseItem{}
var objectSchemaInstance = models.ObjectSchema{}
var transactionInstance = models.Transaction{}
var transactionDetailInstance = models.TransactionDetail{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error
	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshTable() error {
	err := server.DB.DropTableIfExists(&models.TransactionDetail{}, &models.Transaction{}, &models.ObjectSchema{}, &models.WarehouseItem{}, &models.SupplierItem{}, &models.Customer{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.Customer{}, &models.SupplierItem{}, &models.WarehouseItem{}, &models.ObjectSchema{}, &models.Transaction{}, &models.TransactionDetail{}).Error
	if err != nil {
		return err
	}
	fmt.Printf("Successfully refreshed tables")
	return nil
}

func SeedDataTest() ([]models.Customer, []models.SupplierItem, []models.WarehouseItem, []models.ObjectSchema, []models.Transaction, []models.TransactionDetail, error) {
	var err error
	//err = refreshTable()
	if err != nil {
		return []models.Customer{}, []models.SupplierItem{}, []models.WarehouseItem{}, []models.ObjectSchema{}, []models.Transaction{}, []models.TransactionDetail{}, err
	}

	var customers = []models.Customer{
		{
			Nickname: "Fadlika Dita Nurjanto",
			Email:    "fadlikadn@gmail.com",
			Password: "password",
		},
		{
			Nickname: "Fauzan Ibnu Prihadiyono",
			Email:    "fauzan@gmail.com",
			Password: "password",
		},
	}

	var supplierItems = []models.SupplierItem{
		{
			UPC:         "123456",
			BrandName:   "Nike",
			ModelName:   "Basketball Shoes",
			Description: "Nike Basketball Shoes",
			Price:       1000000,
			Quantity:    20,
		},
		{
			UPC:         "987654",
			BrandName:   "Adidas",
			ModelName:   "Football Shoes",
			Description: "Adidas Football Shoes",
			Price:       750000,
			Quantity:    15,
		},
	}

	var warehouseItems = []models.WarehouseItem{
		{
			SKU:         "123456-NIKE-BLACK-42",
			BrandName:   "Nike",
			ModelName:   "Basketball Shoes Black Size 42",
			Description: "Nike Basketball Shoes Black Size 42",
			Price:       1000000,
			Quantity:    5,
		},
		{
			SKU:         "123456-NIKE-BLACK-43",
			BrandName:   "Nike",
			ModelName:   "Basketball Shoes Black Size 43",
			Description: "Nike Basketball Shoes Black Size 43",
			Price:       1000000,
			Quantity:    5,
		},
		{
			SKU:         "123456-NIKE-BLACK-44",
			BrandName:   "Nike",
			ModelName:   "Basketball Shoes Black Size 44",
			Description: "Nike Basketball Shoes Black Size 44",
			Price:       1000000,
			Quantity:    5,
		},
		{
			SKU:         "123456-NIKE-BLACK-45",
			BrandName:   "Nike",
			ModelName:   "Basketball Shoes Black Size 45",
			Description: "Nike Basketball Shoes Black Size 45",
			Price:       1000000,
			Quantity:    5,
		},
	}

	var objectSchemas = []models.ObjectSchema{
		{
			Description: "additional information for warehouse item based on SKU variations",
			Type:        "size",
			Value:       "42",
		},
		{
			Description: "additional information for warehouse item based on SKU variations",
			Type:        "size",
			Value:       "43",
		},
		{
			Description: "additional information for warehouse item based on SKU variations",
			Type:        "size",
			Value:       "44",
		},
		{
			Description: "additional information for warehouse item based on SKU variations",
			Type:        "size",
			Value:       "45",
		},
	}

	var transactions = []models.Transaction{
		{
			Description: "Transaction for Customer Fadli",
		},
		{
			Description: "Transaction for Customer Fauzan",
		},
	}

	var transactionDetails = []models.TransactionDetail{
		{
			Quantity: 2,
			Price:    1000000,
		},
		{
			Quantity: 3,
			Price:    950000,
		},
		{
			Quantity: 1,
			Price:    750000,
		},
		{
			Quantity: 5,
			Price:    500000,
		},
	}

	// Adding foreign key
	err = server.DB.Model(&models.WarehouseItem{}).AddForeignKey("supplier_item_id", "supplier_items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table WarehouseItem error: %v", err)
	}

	err = server.DB.Model(&models.ObjectSchema{}).AddForeignKey("warehouse_item_id", "warehouse_items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table ObjectSchema error: %v", err)
	}

	err = server.DB.Model(&models.Transaction{}).AddForeignKey("customer_id", "customers(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table Transaction error: %v", err)
	}

	err = server.DB.Model(&models.TransactionDetail{}).AddForeignKey("transaction_id", "transactions(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table TransactionDetails transactionID error: %v", err)
	}

	err = server.DB.Model(&models.TransactionDetail{}).AddForeignKey("warehouse_item_id", "warehouse_items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table TransactionDetails warehouseItemID error: %v", err)
	}

	for i, _ := range customers {
		err = server.DB.Model(&models.Customer{}).Create(&customers[i]).Error
		if err != nil {
			log.Fatalf("cannot seed customers table: %v", err)
		}
	}

	for i, _ := range supplierItems {
		err = server.DB.Model(&models.SupplierItem{}).Create(&supplierItems[i]).Error
		if err != nil {
			log.Fatalf("cannot seed supplierItems table: %v", err)
		}
	}

	for i, _ := range warehouseItems {
		warehouseItems[i].SupplierItemID = supplierItems[0].ID
		err = server.DB.Model(&models.WarehouseItem{}).Create(&warehouseItems[i]).Error
		if err != nil {
			log.Fatalf("cannot seed warehouseItems table: %v", err)
		}
	}

	for i, _ := range objectSchemas {
		objectSchemas[i].WarehouseItemID = warehouseItems[i].ID
		err = server.DB.Model(&models.ObjectSchema{}).Create(&objectSchemas[i]).Error
		if err != nil {
			log.Fatalf("cannot seed objectSchemas table: %v", err)
		}
	}

	for i, _ := range transactions {
		transactions[i].CustomerID = customers[i].ID
		err = server.DB.Model(&models.Transaction{}).Create(&transactions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed transactions table: %v", err)
		}
	}

	for i, _ := range transactionDetails {
		transactionDetails[i].TransactionID = transactions[0].ID
		transactionDetails[i].WarehouseItemID = warehouseItems[i].ID
		err = server.DB.Model(&models.TransactionDetail{}).Create(&transactionDetails[i]).Error
		if err != nil {
			log.Fatalf("cannot seed transactionDetails table: %v", err)
		}
	}

	return customers, supplierItems, warehouseItems, objectSchemas, transactions, transactionDetails, nil
}
