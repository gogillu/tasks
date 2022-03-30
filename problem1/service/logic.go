package service

import (
	"fmt"
	"problem1/item"
	"problem1/util"
)

func TakeItemInputFromUser() []item.Item {
	// initialize vars
	takeFurtherInput := true
	itemList := []item.Item{}

	//loop for item input
	for takeFurtherInput == true {

		//input a new Item
		newItem := item.InputItem()
		// fmt.Println(newItem)

		//validate and add
		if !item.ValidateItem(newItem) {
			fmt.Println("Item entered is invalid, can't add, please try again...!!!")
		} else {
			itemList = append(itemList, newItem)
			fmt.Println("==> new Item Added", newItem)
			// fmt.Println(itemList)
		}

		takeFurtherInput := util.AskForUserChoice()
		// fmt.Println(takeFurtherInput)
		if util.ValidateYesNoChoice(takeFurtherInput) && takeFurtherInput == "n" {
			break
		}
	}

	// fmt.Println(itemList)
	return itemList
}

func CalculateTaxOnItems(items []item.Item) []item.Item {
	// itemList := []item.Item{}

	for index, itemi := range items {
		items[index] = *item.UpdateTaxOnItem(itemi)
		// itemList = append(itemList, item)
		// fmt.Println("--", index, item)
		// fmt.Println("--", items)
	}

	// items = append(*items, item.Item{})
	// item.PrintItemDetails(itemList)
	return items
}

func TheLogic() bool {
	fmt.Println("Logic")

	itemList := TakeItemInputFromUser()
	itemList = CalculateTaxOnItems(itemList)
	item.PrintItemDetails(itemList)
	return true
}
