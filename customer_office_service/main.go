package main

import (
	"customer_office_service/controller"
)

func main() {
	controller.New().SetRouts().ServeHTTP()
}
