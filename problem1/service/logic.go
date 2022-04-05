package service

import (
	"fmt"
	"problem1/io"
	"problem1/item"
	"problem1/util"
)

func TakeItemInputFromUser() []item.Item {

	itemList := []item.Item{}

	//loop for item input
	for {

		newItem := io.InputItem()

		//validate item and add
		if !item.ValidateItem(newItem) {
			fmt.Println("Item entered is invalid, can't add, please try again...!!!")
		} else {
			itemList = append(itemList, newItem)
			fmt.Println("==> new Item Added", newItem)
		}

		takeFurtherInput := util.AskForUserChoice()
		if util.ValidateYesNoChoice(takeFurtherInput) && takeFurtherInput == "n" {
			break
		}
	}

	// fmt.Println(itemList)
	return itemList
}

func CalculateTaxOnItems(items []item.Item) []item.Item {

	for index, itemi := range items {
		items[index] = *item.UpdateTaxOnItem(itemi)
	}

	return items
}

func CliItemInputAndCalculation() bool {

	itemList := TakeItemInputFromUser()
	itemList = CalculateTaxOnItems(itemList)
	io.PrintItemDetails(itemList)
	return true
}
