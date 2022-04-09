package item

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateValidItem(t *testing.T) {

	// Positive Cases
	positiveTests := []struct {
		item           Item
		expectedOutput error
	}{
		{Item{"item1", 50, 2, "raw"}, nil},
		{Item{"item2", 10, 125, "manufactured"}, nil},
		{Item{"item3", 2, 10, "imported"}, nil},
		{Item{"item4", 110, 4, "imported"}, nil},
	}

	for _, tc := range positiveTests {
		resp := tc.item.Validate()
		require.Equal(t, tc.expectedOutput, resp)
	}
}

func TestValidateInvalidItem(t *testing.T) {

	// Negative Cases
	negativeTests := []struct {
		item           Item
		expectedOutput error
	}{
		{Item{"item2", 10, -20, "manufactured"}, nil},
		{Item{"item3", -2, 10, "imported"}, nil},
	}

	for _, tc := range negativeTests {
		resp := tc.item.Validate()
		require.NotEqual(t, tc.expectedOutput, resp)
	}

}

func TestCalculateTaxOnItem(t *testing.T) {

	tests := []struct {
		item        Item
		expectedTax float64
	}{
		{Item{"item1", 100, 20, "raw"}, 250.0},
		{Item{"item2", 1000, 21, "manufactured"}, 5302.5},
		{Item{"item3", 40, 2, "imported"}, 13.0},
		{Item{"item4", 50, 3, "imported"}, 25.0},
		{Item{"item5", 100, 35, "imported"}, 542.5},
	}

	for _, tc := range tests {
		resp := tc.item.CalculateTax()
		require.Equal(t, tc.expectedTax, resp)
	}
}
