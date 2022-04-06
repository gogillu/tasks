package io

import (
	"fmt"
	"problem1/item"
)

// input an item
func InputItem() item.Item {

	var itemName, itemType string
	var itemPrice, itemQuantity int

	// take all four values as input
	fmt.Println("\n--------------------")
	fmt.Print("Enter item name : ")
	fmt.Scan(&itemName)
	fmt.Print("Enter item price : ")
	fmt.Scan(&itemPrice)
	fmt.Print("Enter item quantity : ")
	fmt.Scan(&itemQuantity)
	fmt.Print("Enter item type : ")
	fmt.Scan(&itemType)

	newItem := item.New(itemName, itemPrice, itemQuantity, itemType)
	return newItem
}

func GetUserChoice() string {
	var choice string
	fmt.Print("\nDo you want to enter details of any other item (y/n):")
	fmt.Scan(&choice)
	return choice
}

func ValidateYesNoChoice(choice string) bool {
	if choice == "y" || choice == "n" {
		return true
	}
	return false
}

func PrintItemDetails(itemTaxMap map[item.Item]float64) {
	fmt.Println()
	fmt.Println("+-----------+-----------+----------+----------------+---------------------+---------------------+")
	fmt.Printf("|%10s |%10s |%9s |%15s |%20s |%20s |\n", "Item Name", "Price", "Quantity", "Type", "Tax", "Total Cost")

	for item, tax := range itemTaxMap {
		fmt.Println("+-----------+-----------+----------+----------------+---------------------+---------------------+")
		fmt.Printf("|%10s |%10d |%9d |%15s |%20f |%20f |\n", item.Name, item.Price, item.Quantity, item.Type, tax, tax+float64(item.Price)*float64(item.Quantity))
	}
	fmt.Println("+-----------+-----------+----------+----------------+---------------------+---------------------+")

	/*
		fmt.Printf("name     : %20s\n", item.name)
		fmt.Println("price    : ", item.price)
		fmt.Println("quantity : ", item.quantity)
		fmt.Println("type     : ", item.ttype)
		fmt.Println("tax      : ", item.tax)
	*/
}
