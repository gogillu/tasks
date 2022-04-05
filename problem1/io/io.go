package io

import (
	"fmt"
	"problem1/item"
)

/*
	Doubt :
		here the object should be created first and then vales to be taken as input
		or
		values to be taken first and then to create object ?
*/
// input an item
func InputItem() item.Item {

	// create empty item object
	newItem := item.Item{}

	// take all four values as input
	fmt.Println("\n--------------------")
	fmt.Print("Enter item name : ")
	fmt.Scan(&newItem.Name)
	fmt.Print("Enter item price : ")
	fmt.Scan(&newItem.Price)
	fmt.Print("Enter item quantity : ")
	fmt.Scan(&newItem.Quantity)
	fmt.Print("Enter item type : ")
	fmt.Scan(&newItem.Ttype)

	return newItem
}

func PrintItemDetails(items []item.Item) {
	fmt.Println()
	fmt.Println("+-----------+-----------+----------+----------------+---------------------+")
	fmt.Printf("|%10s |%10s |%9s |%15s |%20s |\n", "Item Name", "Price", "Quantity", "Type", "Tax")

	for _, item := range items {
		fmt.Println("+-----------+-----------+----------+----------------+---------------------+")
		fmt.Printf("|%10s |%10d |%9d |%15s |%20f |\n", item.Name, item.Price, item.Quantity, item.Ttype, item.Tax)
	}
	fmt.Println("+-----------+-----------+----------+----------------+---------------------+")

	/*
		fmt.Printf("name     : %20s\n", item.name)
		fmt.Println("price    : ", item.price)
		fmt.Println("quantity : ", item.quantity)
		fmt.Println("type     : ", item.ttype)
		fmt.Println("tax      : ", item.tax)
	*/
}
