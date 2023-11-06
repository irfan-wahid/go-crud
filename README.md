# go-crud

This is an example of an API program for Create, Read, Update, and Delete (CRUD) using Go Language with the Echo framework.

## Config
First, you have to create a table named "barang" using a MySQL database. The following are the columns contained in the following table.

id = int, primaryKey, Auto Increment,
nama = varchar,
stok = int,
harga = float,

Last, you must change the variable value on config/database.go as you create at MySQL database.

## Run the app

    go run main.go

# REST API

The REST API to the example app is described below.

## Get All Data

### Request

`GET /api/v1/barang/get_all`

### Response

    {
        "status": 200,
        "messages": "Success to get items",
        "data": [
            {
            "id": 4,
            "nama": "gelas",
            "stok": 100,
            "harga": 10000
            },
            {
            "id": 5,
            "nama": "Tas",
            "stok": 200,
            "harga": 100000
            }
        ]
    }

## Create a new data

### Request

`POST /api/v1/barang/create`

    {
    "nama" : "Baju",
    "stok" : 50,
    "harga" : 75000
    }

### Response

    {
        "status": 200,
        "messages": "Success to create new item",
        "data": null
    }

## Get a specific Data

### Request

`GET /api/v1/barang/detail/:id`

### Response

    {
        "status": 200,
        "messages": "Success to get items",
        "data": {
            "id": 4,
            "nama": "gelas",
            "stok": 100,
            "harga": 10000
        }
    }

## Update a Data

### Request

`PUT /api/v1/barang/update/:id`

{
  "nama" : "Meja",
  "stok" : 50,
  "harga" : 75000
}

### Response

    {
        "status": 200,
        "messages": "Success to update item",
        "data": null
    }

## Delete a Data

### Request

`DELETE /api/v1/barang/delete/:id`

### Response

    {
        "status": 200,
        "messages": "Success to delete item",
        "data": null
    }