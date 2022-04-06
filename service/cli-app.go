package service

import (
	"fmt"
	"problem1/io"
	"problem1/item"
)

func TakeItemInputFromUser() []item.Item {

	itemList := []item.Item{}

	//loop for item input
	for {

		newItem := io.InputItem()

		//validate item and add
		if newItem.ValidateItem() != nil {
			fmt.Println("Item entered is invalid, can't add, please try again...!!!")
		} else {
			itemList = append(itemList, newItem)
			fmt.Println("==> new Item Added", newItem)
		}

		takeFurtherInput := io.GetUserChoice()
		if io.ValidateYesNoChoice(takeFurtherInput) && takeFurtherInput == "n" {
			break
		}
	}

	// fmt.Println(itemList)
	return itemList
}

func CalculateTaxOnItems(items []item.Item) map[item.Item]float64 {
	var itemTaxMap map[item.Item]float64 = make(map[item.Item]float64)

	for _, itemi := range items {
		itemTaxMap[itemi] = itemi.CalculateTaxOnItem()
	}

	return itemTaxMap
}

func CliItemInputAndCalculation() bool {

	itemList := TakeItemInputFromUser()
	fmt.Println(itemList)
	itemTaxMap := CalculateTaxOnItems(itemList)
	io.PrintItemDetails(itemTaxMap)
	return true
}
