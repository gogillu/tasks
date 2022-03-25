package concrete

import (
	"testing"
)

func TestValidateYesNoChoice(t *testing.T) {

	var expectedOutput, actualOutput bool

	// Positive cases
	//  case1
	actualOutput = ValidateYesNoChoice("y")
	expectedOutput = true

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = ValidateYesNoChoice("n")
	expectedOutput = true

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}

	//  case2
	actualOutput = ValidateYesNoChoice("yes")
	expectedOutput = false

	if expectedOutput != actualOutput {
		t.Errorf("expected %t, actual %t", expectedOutput, actualOutput)
	}
}
