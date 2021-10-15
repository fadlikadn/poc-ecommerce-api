# POC Ecommerce API

1. Describe what you think happened that caused those bad reviews during our 12.12 event and why it happened. 
2. Based on your analysis, propose a solution that will prevent the incidents from occurring again.
3. Build a POC to demo technical solution


## Analysis
The misreported can be happened because system doesn't implement SKU for the inventory management. Inventory seems only use UPC (Universal Product Code) that get from product.

## Solution
We should use SKU (Stock Keeping Unit) to track the product's variation.
For example shoes. A model of shoes can be split into several variation, such as color, size. Furthermore, we can include "tanggal masuk" barang di gudang and "warehouse code" if the system already been scaled to have multiple warehouses.

###Scope limit :
This POC only using **Transactions** & **TransactionDetails** table, which this may be different with the real ecommerce that has more complex _transaction's_ table to handle enter items in cart, checkout, and then doing payment. This POC only validate process on **addNewTransactionDetails** process, which will compare the total amount of current stock in table WarehouseItem with total warehouseItem data that already purchased with recorded in TransactionDetails tables.

### Entity (POC, this can be expanded further to add supplier data, warehouse entry process, purchase process, etc. But the focus here is demonstrating using SKU)
- **Customer**
  - ID (UUID)
  - Name
- **Supplier Item**
  - ID (UUID) 
  - UPC from supplier
  - Brand Name
  - Model Name
  - Item's Categorize
  - Description
  - Price (let's make simple for now, assume in IDR)
  - Quantity (total amount of product in various size/color)
- **Warehouse Item**
  - ID (UUID)
  - Supplier Item's ID
  - SKU
  - Brand Name
  - Model Name
  - Description
  - Price (IDR)
  - Quantity (total on specific size/color)
- **Object Schema - Warehouse (this table is used to connect Warehouse Item with Object Schema Detail - Color, size, etc)**
  - ID (UUID)
  - Warehouse Item's ID
  - Description
  - Type (ex: color, size)
  - Value (ex: blue, 43, 44, L, XL, etc, depends on the product)
- **Transaction**
  - ID (UUID)
  - Customer ID
- **Transaction Detail**
  - ID (UUID)
  - Transaction's ID
  - Warehouse Item's ID
  - Quantity
  - Price (lock current price that customer buy)

### Run Test
To run test command, please change the directory to `tests/controllers` first. The following command tests :
-  `go test -v --run TestAddTransactionDetails`