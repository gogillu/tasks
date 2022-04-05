package item

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateItemNew(t *testing.T) {
	tests := []struct {
		name     string
		price    int
		quantity int
		ttype    string
		tax      float64
		output   bool
	}{
		{
			"item1", 50, 2, "raw", 0.0, true,
		},
	}

	for _, tc := range tests {
		resp := ValidateItem(Item{tc.name, tc.price, tc.quantity, tc.ttype, tc.tax})
		require.Equal(t, tc.output, resp)
	}
}

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

func TestValidatettype(t *testing.T) {

	var expectedOutput, actualOutput bool

	// Positive case
	//  case1 :
	actualOutput = Validatettype("raw")
	expectedOutput = true

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case2 :
	actualOutput = Validatettype("manufactured")
	expectedOutput = true

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case3 :
	actualOutput = Validatettype("imported")
	expectedOutput = true

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	// Negative Cases
	//  case4 :
	actualOutput = Validatettype("rawwww")
	expectedOutput = false

	if actualOutput != expectedOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case5 :
	actualOutput = Validatettype("refurbished")
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

	if actualOutput.Tax != expectedOutput.Tax {
		t.Errorf("expected %+v, actual %+v", expectedOutput, actualOutput)
	}

	//  case2 :
	item2 := Item{"item2", 40, 3, "manufactured", 0}

	actualOutput = UpdateTaxOnItem(item2)
	expectedOutput = Item{"item2", 40, 3, "raw", 40.0 * 3.0 * (12.5 + 2*(1+12.5/100)) / 100}

	if actualOutput.Tax != expectedOutput.Tax {
		t.Errorf("expected %+v, actual %+v", expectedOutput, actualOutput)
	}

	//  case3 :
	item3 := Item{"item3", 150, 300, "imported", 0}

	actualOutput = UpdateTaxOnItem(item3)
	expectedOutput = Item{"item3", 150, 300, "raw", 150.0 * 300.0 * (10.0 / 100.0) * (1.0 + 105.0/100.0)}

	if actualOutput.Tax != expectedOutput.Tax {
		t.Errorf("expected %+v, actual %+v", expectedOutput, actualOutput)
	}

}

func TestCalculateRawtax(t *testing.T) {

	var expectedOutput, actualOutput float64

	// Positive cases
	//  case1
	actualOutput = CalculateRawtax(100, 20)
	expectedOutput = 250.0

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = CalculateRawtax(5, 10)
	expectedOutput = 6.250

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

}

func TestCalulateManufacturedtax(t *testing.T) {

	var expectedOutput, actualOutput float64

	// Positive cases
	//  case1
	actualOutput = CalulateManufacturedtax(10, 2)
	expectedOutput = 10.0 * 2.0 * (12.5 + 2*(1+12.5/100)) / 100

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = CalulateManufacturedtax(5, 10)
	expectedOutput = 5.0 * 10 * (12.5 + 2*(1+12.5/100)) / 100

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

}

func TestCalulateImportedtax(t *testing.T) {

	var expectedOutput, actualOutput float64

	// Positive cases
	//  case1
	actualOutput = CalulateImportedtax(2, 3)
	expectedOutput = 2.0*3.0*10.0/100 + 5.0

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = CalulateImportedtax(50, 3)
	expectedOutput = 50.0*3.0*10.0/100 + 10.0

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case3
	actualOutput = CalulateImportedtax(80, 40)
	expectedOutput = 80.0 * 40.0 * 10.0 / 100 * (1.0 + 105.0/100)

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}
	/*
		//  case4
		actualOutput = CalulateImportedtax(85, 45)
		expectedOutput = 85.0 * 45.0 * 10.0 / 100 * (1.0 + 105.0/100)

		if expectedOutput != actualOutput {
			t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
		}
	*/
}
