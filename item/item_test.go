package item

import (
	"fmt"
	"testing"

	"github.com/gogillu/tax-manager/item/enum"

	"github.com/stretchr/testify/require"
)

func TestNewItem(t *testing.T) {
	tests := []struct {
		itmName     string
		itmPrice    int
		itmQuantity int
		itmType     string
		expErr      error
	}{
		{
			"item1",
			50,
			2,
			"raw",
			nil,
		},
		{
			"item2",
			10,
			125,
			"manufactured",
			nil,
		},
		{
			"item3",
			2,
			10,
			"imported",
			nil,
		},
		{
			"item4",
			110,
			4,
			"imported",
			nil,
		},
		{
			"item5",
			2,
			-10,
			"imported",
			fmt.Errorf("error : item validation failed"),
		},
		{
			"item6",
			1000,
			4,
			"new",
			fmt.Errorf("error : provided item type value is not valid"),
		},
	}

	for _, tc := range tests {
		_, err := New(tc.itmName, tc.itmPrice, tc.itmQuantity, tc.itmType)
		require.Equal(t, tc.expErr, err)
	}
}

func TestValidateItem(t *testing.T) {

	tests := []struct {
		item   Item
		expErr error
	}{
		{
			Item{"item1", 50, 2, enum.Raw},
			nil,
		},
		{
			Item{"item2", 10, 125, enum.Manufactured},
			nil,
		},
		{
			Item{"item3", 2, 10, enum.Imported},
			nil,
		},
		{
			Item{"item4", 110, 4, enum.Imported},
			nil,
		},
		{
			Item{"item5", 10, -20, enum.Manufactured},
			fmt.Errorf("invalid quantity"),
		},
		{
			Item{"item6", -2, 10, enum.Imported},
			fmt.Errorf("invalid price"),
		},
	}

	for _, tc := range tests {
		resp := tc.item.validate()
		if tc.expErr != nil {
			require.NotNil(t, resp)
		} else {
			require.Equal(t, tc.expErr, resp)
		}
	}
}

func TestCalculateTax(t *testing.T) {

	tests := []struct {
		item        Item
		expectedTax float64
	}{
		{Item{"item1", 100, 20, enum.Raw}, 250.0},
		{Item{"item2", 1000, 21, enum.Manufactured}, 5302.5},
		{Item{"item3", 40, 2, enum.Imported}, 13.0},
		{Item{"item4", 50, 3, enum.Imported}, 25.0},
		{Item{"item5", 100, 35, enum.Imported}, 542.5},
	}

	for _, tc := range tests {
		resp := tc.item.CalculateTax()
		require.Equal(t, tc.expectedTax, resp)
	}
}

func TestGetSurcharge(t *testing.T) {

	tests := []struct {
		price             float64
		expectedSurcharge float64
	}{
		{50, 5},
		{100, 5},
		{101, 10},
		{200, 10},
		{250, 12.5},
	}

	for _, tc := range tests {
		obtainedSurcharge := getSurcharge(tc.price)
		require.Equal(t, tc.expectedSurcharge, obtainedSurcharge)
	}
}
