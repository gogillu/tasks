package concrete

import "testing"

func TestCalculateRawItemTax(t *testing.T) {

	var expectedOutput, actualOutput float64

	// Positive cases
	//  case1
	actualOutput = CalculateRawItemTax(100, 20)
	expectedOutput = 250.0

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = CalculateRawItemTax(5, 10)
	expectedOutput = 6.250

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

}

func TestCalulateManufacturedItemTax(t *testing.T) {

	var expectedOutput, actualOutput float64

	// Positive cases
	//  case1
	actualOutput = CalulateManufacturedItemTax(10, 2)
	expectedOutput = 10.0 * 2.0 * (12.5 + 2*(1+12.5/100)) / 100

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = CalulateManufacturedItemTax(5, 10)
	expectedOutput = 5.0 * 10 * (12.5 + 2*(1+12.5/100)) / 100

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

}

func TestCalulateImportedItemTax(t *testing.T) {

	var expectedOutput, actualOutput float64

	// Positive cases
	//  case1
	actualOutput = CalulateImportedItemTax(2, 3)
	expectedOutput = 2.0*3.0*10.0/100 + 5.0

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = CalulateImportedItemTax(50, 3)
	expectedOutput = 50.0*3.0*10.0/100 + 10.0

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}

	//  case3
	actualOutput = CalulateImportedItemTax(80, 40)
	expectedOutput = 80.0 * 40.0 * 10.0 / 100 * (1.0 + 105.0/100)

	if expectedOutput != actualOutput {
		t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
	}
	/*
		//  case4
		actualOutput = CalulateImportedItemTax(85, 45)
		expectedOutput = 85.0 * 45.0 * 10.0 / 100 * (1.0 + 105.0/100)

		if expectedOutput != actualOutput {
			t.Errorf("expected %f, actual %f", expectedOutput, actualOutput)
		}
	*/
}
