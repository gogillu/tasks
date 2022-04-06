package enum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTaxOnItem(t *testing.T) {
	// Positive Cases
	positiveTests := []struct {
		itemTypeString   string
		expectedItemType ItemType
		expectedError    error
	}{
		{"raw", Raw, nil},
		{"manufactured", Manufactured, nil},
		{"imported", Imported, nil},
	}

	for _, tc := range positiveTests {
		resp, err := GetItemTypeFromString(tc.itemTypeString)
		require.Equal(t, tc.expectedItemType, resp)
		require.Equal(t, tc.expectedError, err)
	}

	// Negative Cases
	negativeTests := []struct {
		itemTypeString   string
		expectedItemType ItemType
		expectedError    error
	}{
		{"new", Raw, nil},
		{"old", Manufactured, nil},
	}

	for _, tc := range negativeTests {
		resp, err := GetItemTypeFromString(tc.itemTypeString)
		require.NotEqual(t, tc.expectedItemType, resp)
		require.NotEqual(t, tc.expectedError, err)
	}

}
