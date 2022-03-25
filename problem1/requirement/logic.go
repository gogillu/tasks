package requirement

import (
	"fmt"
	"problem1/concrete"
)

func TakeItemInputFromUser() []concrete.Item {
	// initialize vars
	takeFurtherInput := true
	itemList := []concrete.Item{}

	//loop for item input
	for takeFurtherInput == true {

		//input a new Item
		newItem := concrete.InputItem()
		// fmt.Println(newItem)

		//validate and add
		if !concrete.ValidateItem(newItem) {
			fmt.Println("Item entered is invalid, can't add, please try again...!!!")
		} else {
			itemList = append(itemList, newItem)
			fmt.Println("==> new Item Added", newItem)
			// fmt.Println(itemList)
		}

		takeFurtherInput := concrete.AskForUserChoice()
		// fmt.Println(takeFurtherInput)
		if concrete.ValidateYesNoChoice(takeFurtherInput) && takeFurtherInput == "n" {
			break
		}
	}

	// fmt.Println(itemList)
	return itemList
}

func CalculateTaxOnItems(items []concrete.Item) []concrete.Item {
	// itemList := []concrete.Item{}

	for index, item := range items {
		items[index] = *concrete.UpdateTaxOnItem(item)
		// itemList = append(itemList, item)
		// fmt.Println("--", item)
		// fmt.Println("--", items)
	}

	// items = append(*items, concrete.Item{})
	// concrete.PrintItemDetails(itemList)
	return items
}

func TheLogic() bool {
	fmt.Println("Logic")

	itemList := TakeItemInputFromUser()
	itemList = CalculateTaxOnItems(itemList)
	concrete.PrintItemDetails(itemList)
	return true
}
