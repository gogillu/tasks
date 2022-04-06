package enum

import "fmt"

type ItemType string

const (
	Raw          ItemType = "raw"
	Manufactured ItemType = "manufactured"
	Imported     ItemType = "imported"
)

var allItemType = []ItemType{
	Raw,
	Manufactured,
	Imported,
}

func GetItemTypeFromString(s string) (ItemType, error) {
	for itemType := range allItemType {
		if string(allItemType[itemType]) == s {
			return allItemType[itemType], nil
		}
	}
	return "", fmt.Errorf("%s provided string is not a valid ItemType", s)
}
