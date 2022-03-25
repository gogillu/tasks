package concrete

import "fmt"

type Item struct {
	itemName     string
	itemPrice    int
	itemQuantity int
	itemType     string
	itemTax      float64
}

/*
	Doubt :
		here the object should be created first and then vales to be taken as input
		or
		values to be taken first and then to create object ?
*/
// input an item
func InputItem() Item {

	// create empty item object
	newItem := Item{}

	// take all four values as input
	fmt.Println("\n--------------------")
	fmt.Print("Enter item name : ")
	fmt.Scan(&newItem.itemName)
	fmt.Print("Enter item price : ")
	fmt.Scan(&newItem.itemPrice)
	fmt.Print("Enter item quantity : ")
	fmt.Scan(&newItem.itemQuantity)
	fmt.Print("Enter item type : ")
	fmt.Scan(&newItem.itemType)

	return newItem
}

func ValidateItem(givenItem Item) bool {

	if givenItem.itemName == "" {
		return false
	}

	if givenItem.itemPrice <= 0 {
		return false
	}

	if givenItem.itemQuantity <= 0 {
		return false
	}

	if !ValidateItemType(givenItem.itemType) {
		return false
	}

	return true
}

func ValidateItemType(givenItemType string) bool {
	validItemTypes := [3]string{"raw", "manufactured", "imported"}
	// fmt.Println(validItemTypes)
	for _, itemType := range validItemTypes {
		if itemType == givenItemType {
			return true
		}
	}
	return false
}

func UpdateTaxOnItem(item Item) *Item {
	switch item.itemType {
	case "raw":
		item.itemTax = CalculateRawItemTax(item.itemPrice, item.itemQuantity)

	case "manufactured":
		item.itemTax = CalulateManufacturedItemTax(item.itemPrice, item.itemQuantity)

	case "imported":
		item.itemTax = CalulateImportedItemTax(item.itemPrice, item.itemQuantity)
	}
	// fmt.Println(item)
	return &item
}

func PrintItemDetails(items []Item) {
	fmt.Println()
	fmt.Println("+-----------+-----------+----------+----------------+---------------------+")
	fmt.Printf("|%10s |%10s |%9s |%15s |%20s |\n", "Item Name", "Price", "Quantity", "Type", "Tax")

	for _, item := range items {
		fmt.Println("+-----------+-----------+----------+----------------+---------------------+")
		fmt.Printf("|%10s |%10d |%9d |%15s |%20f |\n", item.itemName, item.itemPrice, item.itemQuantity, item.itemType, item.itemTax)
	}
	fmt.Println("+-----------+-----------+----------+----------------+---------------------+")

	/*
		fmt.Printf("name     : %20s\n", item.itemName)
		fmt.Println("price    : ", item.itemPrice)
		fmt.Println("quantity : ", item.itemQuantity)
		fmt.Println("type     : ", item.itemType)
		fmt.Println("tax      : ", item.itemTax)
	*/
}
