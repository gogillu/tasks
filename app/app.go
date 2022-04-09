package app

import (
	"fmt"
	"tax-manager/item"
)

const (
	YES = "y"
	NO  = "n"
)

func inputItem() []item.Item {

	itemList := []item.Item{}

	for {

		var itmName, itmType string
		var itmPrice, itmQuantity int

		fmt.Println("\n--Enter Item Details--")
		fmt.Print("name : ")
		fmt.Scan(&itmName)
		fmt.Print("price : ")
		fmt.Scan(&itmPrice)
		fmt.Print("quantity : ")
		fmt.Scan(&itmQuantity)
		fmt.Print("type : ")
		fmt.Scan(&itmType)

		itm := item.New(itmName, itmPrice, itmQuantity, itmType)
		if err := itm.Validate(); err != nil {
			fmt.Println("error : invalid item entered, try again ! ", err)
			continue
		}

		itemList = append(itemList, itm)

		var choice string
	loop:
		fmt.Print("\nAdd more items (y/n) : ")
		fmt.Scan(&choice)
		if err := validateChoice(choice); err != nil {
			goto loop
		}

		if choice == NO {
			break
		}
	}

	return itemList
}

func validateChoice(choice string) error {
	if choice == YES || choice == NO {
		return nil
	}
	return fmt.Errorf("error : invalid user choice")
}

func showInvoice(items []item.Item) {
	for _, itm := range items {
		itm.GetInvoice()
	}
}

func Start() {

	items := inputItem()
	showInvoice(items)

}
