package main

import (
	"example.com/go_rest_postgre/controller"
	"example.com/go_rest_postgre/model"
)

func main() {
	model.Init()
	controller.Init()
}
