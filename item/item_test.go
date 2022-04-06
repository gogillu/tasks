package item

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateItemNew(t *testing.T) {

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
		resp := tc.item.ValidateItem()
		require.Equal(t, tc.expectedOutput, resp)
	}

	// Negative Cases
	negativeTests := []struct {
		item           Item
		expectedOutput error
	}{
		{Item{"item2", 10, -20, "manufactured"}, nil},
		{Item{"item3", -2, 10, "imported"}, nil},
	}

	for _, tc := range negativeTests {
		resp := tc.item.ValidateItem()
		require.NotEqual(t, tc.expectedOutput, resp)
	}

}

func TestCalculateTaxOnItem(t *testing.T) {
	// Positive Cases
	positiveTests := []struct {
		item        Item
		expectedTax float64
	}{
		{Item{"item1", 100, 20, "raw"}, 250.0},
		{Item{"item2", 5, 1000, "imported"}, 775.0},
		{Item{"item3", 1000, 21, "manufactured"}, 3097.5},
		{Item{"item4", 50, 2, "imported"}, 15.0},
	}

	for _, tc := range positiveTests {
		resp := tc.item.CalculateTaxOnItem()
		require.Equal(t, tc.expectedTax, resp)
	}
}
