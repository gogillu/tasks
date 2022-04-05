package main

import (
	"flag"
	"fmt"
	"problem1/service"
)

func getCmdArgs() {
	var name = flag.String("name", "", "Name of the item")
	var price = flag.Int("price", 0, "Name of the item")
	var quantity = flag.Int("quantity", 0, "Name of the item")
	var ttype = flag.String("type", "", "Name of the item")
	flag.Parse()
	fmt.Println(*name, *price, *quantity, *ttype)
}

func main() {
	fmt.Println("main starts")

	//call the task
	service.CliItemInputAndCalculation()
}
