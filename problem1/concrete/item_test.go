package concrete

import (
	"testing"
)

func TestValidateItem(t *testing.T) {

	var expectedOutput, actualOutput bool

	// Positive cases
	//  case1
	actualOutput = ValidateItem(Item{"item1", 50, 2, "raw", 0})
	expectedOutput = true

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = ValidateItem(Item{"item2", 10, 5, "manufactured", 0})
	expectedOutput = true

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case3
	actualOutput = ValidateItem(Item{"item3", 150, 1, "imported", 0})
	expectedOutput = true

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	// Negative cases
	//  case4
	actualOutput = ValidateItem(Item{"", 6, 2, "raw", 0})
	expectedOutput = false

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case5
	actualOutput = ValidateItem(Item{"item5", -1, 3, "manufactured", 0})
	expectedOutput = false

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case6
	actualOutput = ValidateItem(Item{"item4", 13, 0, "raw", 0})
	expectedOutput = false

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case7
	actualOutput = ValidateItem(Item{"item4", 21, 3, "new", 0})
	expectedOutput = false

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

}

func TestValidateItemType(t *testing.T) {

	var expectedOutput, actualOutput bool

	// Positive case
	//  case1 :
	actualOutput = ValidateItemType("raw")
	expectedOutput = true

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case2 :
	actualOutput = ValidateItemType("manufactured")
	expectedOutput = true

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case3 :
	actualOutput = ValidateItemType("imported")
	expectedOutput = true

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	// Negative Cases
	//  case4 :
	actualOutput = ValidateItemType("rawwww")
	expectedOutput = false

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case5 :
	actualOutput = ValidateItemType("refurbished")
	expectedOutput = false

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

}

func TestUpdateTaxOnItem(t *testing.T) {

	var expectedOutput Item
	var actualOutput *Item

	// Positive case
	//  case1 :
	item1 := Item{"item1", 100, 25, "raw", 0}
	actualOutput = UpdateTaxOnItem(item1)
	expectedOutput = Item{"item1", 100, 25, "raw", 100.0 * 25.0 * 12.5 / 100}

	if actualOutput.itemTax != expectedOutput.itemTax {
		t.Errorf("expected %+v, actual %+v", expectedOutput, actualOutput)
	}

	//  case2 :
	item2 := Item{"item2", 40, 3, "manufactured", 0}

	actualOutput = UpdateTaxOnItem(item2)
	expectedOutput = Item{"item2", 40, 3, "raw", 40.0 * 3.0 * (12.5 + 2*(1+12.5/100)) / 100}

	if actualOutput.itemTax != expectedOutput.itemTax {
		t.Errorf("expected %+v, actual %+v", expectedOutput, actualOutput)
	}

	//  case3 :
	item3 := Item{"item3", 150, 300, "imported", 0}

	actualOutput = UpdateTaxOnItem(item3)
	expectedOutput = Item{"item3", 150, 300, "raw", 150.0 * 300.0 * (10.0 / 100.0) * (1.0 + 105.0/100.0)}

	if actualOutput.itemTax != expectedOutput.itemTax {
		t.Errorf("expected %+v, actual %+v", expectedOutput, actualOutput)
	}

}

func TestPrintItemDetails(t *testing.T) {

	// Positive case
	//  case1 :
	item1 := Item{"item1", 100, 25, "raw", 0}
	item2 := Item{"item2", 40, 3, "manufactured", 0}
	item3 := Item{"item3", 150, 300, "imported", 0}
	items := []Item{item1, item2, item3}

	for index, item := range items {
		items[index] = *UpdateTaxOnItem(item)
	}

	PrintItemDetails(items)
}
