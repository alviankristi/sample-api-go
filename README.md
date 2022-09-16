# catalyst-backend-task

## Specification
* Golang 1.18
* Mysql

## Asumptions
* Since the stock is not managed, the transaction does not handle the product stock
* The product price is fixed and cannot be edited
* The database use mysql driver

## How To:
1. Migrate the database using the below command
```
    go run main.go migrate
```
2. Run the api using the below command
```
    go run main.go serve
```
3. For postman user, please import the collection [here](/Catalyst-BackendTask.postman_collection.json)