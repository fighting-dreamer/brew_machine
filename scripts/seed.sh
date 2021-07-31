#!/bin/sh

## add ingredients
curl -X POST http://localhost:8080/v1/ingredient/add -d '{"Name":"water", "AvailableQuantity":1000}'
curl -X POST http://localhost:8080/v1/ingredient/add -d '{"Name":"milk", "AvailableQuantity":1000}'
curl -X POST http://localhost:8080/v1/ingredient/add -d '{"Name":"sugar", "AvailableQuantity":1000}'

## add beverage
curl -X POST http://localhost:8080/v1/beverage/add -d '{"Name":"chai", "IngredientsQuantityMap": {"water":100, "milk":200}}'

#
### make beverage
#curl -X GET 'http://localhost:8080/v1/dispenser/make?Name=chai&Outlet=2'