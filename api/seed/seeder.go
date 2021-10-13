package seed

import (
	"github.com/fadlikadn/poc-ecommerce-api/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

var customers = []models.Customer {
	{
		Nickname: "Fadlika Dita Nurjanto",
		Email: "fadlikadn@gmail.com",
		Password: "password",
	},
	{
		Nickname: "Fauzan Ibnu Prihadiyono",
		Email: "fauzan@gmail.com",
		Password: "password",
	},
}

var supplierItems = []models.SupplierItem {
	{
		UPC: "123456",
		BrandName: "Nike",
		ModelName: "Basketball Shoes",
		Description: "Nike Basketball Shoes",
		Price: 1000000,
		Quantity: 20,
	},
	{
		UPC: "987654",
		BrandName: "Adidas",
		ModelName: "Football Shoes",
		Description: "Adidas Football Shoes",
		Price: 750000,
		Quantity: 15,
	},
}

var warehouseItems = []models.WarehouseItem {
	{
		SKU: "123456-NIKE-BLACK-42",
		BrandName: "Nike",
		ModelName: "Basketball Shoes Black Size 42",
		Description: "Nike Basketball Shoes Black Size 42",
		Price: 1000000,
		Quantity: 5,
	},
	{
		SKU: "123456-NIKE-BLACK-43",
		BrandName: "Nike",
		ModelName: "Basketball Shoes Black Size 43",
		Description: "Nike Basketball Shoes Black Size 43",
		Price: 1000000,
		Quantity: 5,
	},
	{
		SKU: "123456-NIKE-BLACK-44",
		BrandName: "Nike",
		ModelName: "Basketball Shoes Black Size 44",
		Description: "Nike Basketball Shoes Black Size 44",
		Price: 1000000,
		Quantity: 5,
	},
	{
		SKU: "123456-NIKE-BLACK-44",
		BrandName: "Nike",
		ModelName: "Basketball Shoes Black Size 44",
		Description: "Nike Basketball Shoes Black Size 44",
		Price: 1000000,
		Quantity: 5,
	},
}

var objectSchema = []models.ObjectSchema {
	{
		Description: "additional information for warehouse item based on SKU variations",
		Type: "size",
		Value: "42",
	},
	{
		Description: "additional information for warehouse item based on SKU variations",
		Type: "size",
		Value: "43",
	},
	{
		Description: "additional information for warehouse item based on SKU variations",
		Type: "size",
		Value: "44",
	},
	{
		Description: "additional information for warehouse item based on SKU variations",
		Type: "size",
		Value: "45",
	},
}

var transactions = []models.Transaction {
	{
		Description: "Transaction for Customer Fadli",
	},
	{
		Description: "Transaction for Customer Fauzan",
	},
}

var transactionDetail = []models.TransactionDetail {
	{
		Quantity: 2,
		Price: 1000000,
	},
	{
		Quantity: 3,
		Price: 950000,
	},
	{
		Quantity: 1,
		Price: 750000,
	},
	{
		Quantity: 5,
		Price: 500000,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Customer{}, &models.SupplierItem{}, &models.WarehouseItem{}, &models.ObjectSchema{}, &models.Transaction{}, &models.TransactionDetail{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Customer{}, &models.SupplierItem{}, &models.WarehouseItem{}, &models.ObjectSchema{}, &models.Transaction{}, &models.TransactionDetail{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// Adding foreign key
	err = db.Debug().Model(&models.WarehouseItem{}).AddForeignKey("supplier_item_id", "supplier_items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table WarehouseItem error: %v", err)
	}

	err = db.Debug().Model(&models.ObjectSchema{}).AddForeignKey("warehouse_item_id", "warehouse_items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table ObjectSchema error: %v", err)
	}

	err = db.Debug().Model(&models.Transaction{}).AddForeignKey("customer_id", "customers(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table Transaction error: %v", err)
	}

	err = db.Debug().Model(&models.TransactionDetail{}).AddForeignKey("transaction_id", "transactions(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table TransactionDetails transactionID error: %v", err)
	}

	err = db.Debug().Model(&models.TransactionDetail{}).AddForeignKey("warehouse_item_id", "warehouse_items(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key table TransactionDetails warehouseItemID error: %v", err)
	}
}