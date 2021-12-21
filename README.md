# Product CRUD Example

## Get started
- Install docker with docker-compose

## Features
-  Get all products
-  Get a product by sku
-  Create/Update a product
-  Delete a product by sku

## Run application
```bash
docker-compose up
```

## Example
- Create a product
```bash
curl --location --request PUT 'localhost:8080/products' \
--header 'Content-Type: application/json' \
--data-raw '{
    "sku": "FAL-8406270",
    "name": "500 Zapatilla Urbana Mujer",
    "brand": "New Balance",
    "size": "37",
    "price": 42990.00,
    "principalImage": "https://falabella.scene7.com/is/image/Falabella/8406270_1",
    "otherImages": [
        "https//falabella.scene7.com/is/image/Falabella/8406270_1"
    ]
}'
```
- Get products
```bash
curl --location --request GET 'localhost:8080/products'
```
- Get a product
```bash
curl --location --request GET 'localhost:8080/products/FAL-8406270'
```
- Delete a product
```bash
curl --location --request DELETE 'localhost:8080/products/FAL-8406270'
```