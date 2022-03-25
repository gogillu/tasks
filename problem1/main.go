package main

import (
	"flag"
	"fmt"
	"problem1/requirement"
)

func getCmdArgs() {
	var itemName = flag.String("name", "", "Name of the item")
	var itemPrice = flag.Int("price", 0, "Name of the item")
	var itemQuantity = flag.Int("quantity", 0, "Name of the item")
	var itemType = flag.String("type", "", "Name of the item")
	flag.Parse()
	fmt.Println(*itemName, *itemPrice, *itemQuantity, *itemType)
}

func main() {
	fmt.Println("main starts")

	//call the task
	requirement.TheLogic()
}
