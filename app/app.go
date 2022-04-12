package app

import (
	"fmt"

	"github.com/gogillu/tax-manager/item"
)

const (
	Yes = "y"
	No  = "n"
)

func inputItem() []item.Item {

	itemList := []item.Item{}

	for {

		var itmName string
		var itmPrice int
		var itmQuantity int
		var itmType string

		fmt.Println("\n--Enter Item Details--")
		fmt.Print("name : ")
		fmt.Scan(&itmName)
		fmt.Print("price : ")
		fmt.Scan(&itmPrice)
		fmt.Print("quantity : ")
		fmt.Scan(&itmQuantity)
		fmt.Print("type : ")
		fmt.Scan(&itmType)

		itm, err := item.New(itmName, itmPrice, itmQuantity, itmType)
		if err != nil {
			fmt.Println("error : invalid item entered, try again ! ", err)
			continue
		}

		itemList = append(itemList, itm)

		if choice := getChoice(); choice == No {
			break
		}
	}

	return itemList
}

func getChoice() string {
	var choice string

	for {
		fmt.Print("\nAdd more items (y/n) : ")
		fmt.Scan(&choice)
		if err := validateChoice(choice); err != nil {
			continue
		}

		return choice
	}
}

func validateChoice(choice string) error {
	if choice == Yes || choice == No {
		return nil
	}
	return fmt.Errorf("error : invalid choice")
}

func showInvoice(items []item.Item) {
	for _, itm := range items {
		fmt.Println(itm.GetInvoice())
	}
}

func Start() {

	items := inputItem()
	showInvoice(items)

}
