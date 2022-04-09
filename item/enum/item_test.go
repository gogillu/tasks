package enum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetValidItemType(t *testing.T) {
	tests := []struct {
		itemTypeString   string
		expectedItemType ItemType
		expectedError    error
	}{
		{"raw", Raw, nil},
		{"manufactured", Manufactured, nil},
		{"imported", Imported, nil},
	}

	for _, tc := range tests {
		resp, err := GetItemType(tc.itemTypeString)
		require.Equal(t, tc.expectedItemType, resp)
		require.Equal(t, tc.expectedError, err)
	}

}
func TestGetInvalidItemType(t *testing.T) {

	tests := []struct {
		itemTypeString   string
		expectedItemType ItemType
		expectedError    error
	}{
		{"new", Raw, nil},
		{"old", Manufactured, nil},
	}

	for _, tc := range tests {
		resp, err := GetItemType(tc.itemTypeString)
		require.NotEqual(t, tc.expectedItemType, resp)
		require.NotEqual(t, tc.expectedError, err)
	}

}
