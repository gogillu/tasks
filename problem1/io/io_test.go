package io

import (
	"problem1/item"
	"testing"
)

func TestPrintItemDetails(t *testing.T) {

	// Positive case
	//  case1 :
	item1 := item.Item{}
	item1.Init("item1", 100, 25, "raw")
	item2 := item.Item{}
	item2.Init("item2", 40, 3, "manufactured")
	item3 := item.Item{}
	item3.Init("item3", 150, 300, "imported")
	items := []item.Item{item1, item2, item3}

	for index, it := range items {
		items[index] = *item.UpdateTaxOnItem(it)
	}

	PrintItemDetails(items)
}
