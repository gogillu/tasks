package item

import (
	"fmt"
	"problem1/item/enum"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Item struct {
	Name     string        `json:"name,omitempty"`
	Price    int           `json:"price,omitempty"`
	Quantity int           `json:"quantity,omitempty"`
	Type     enum.ItemType `json:"type,omitempty"`
}

func New(name string, price int, quantity int, Type string) Item {

	itemType, err := enum.GetItemTypeFromString(Type)
	if err != nil {
		panic(err)
	}

	item := Item{}
	item.Name = name
	item.Price = price
	item.Quantity = quantity
	item.Type = itemType

	return item
}

func (item Item) ValidateItem() error {

	return validation.ValidateStruct(&item,
		validation.Field(&item.Price, validation.By(validateNegativeInt)),
		validation.Field(&item.Quantity, validation.By(validateNegativeInt)),
	)
}

func validateNegativeInt(val interface{}) error {
	switch x := val.(type) {
	case int:
		if x < 0 {
			return fmt.Errorf("%d given value is negative", x)
		}
	}
	return nil
}

func (item Item) CalculateTaxOnItem() float64 {
	switch item.Type {
	case enum.Raw:
		return CalculateRawtax(item.Price, item.Quantity)

	case enum.Manufactured:
		return CalulateManufacturedtax(item.Price, item.Quantity)

	case enum.Imported:
		return CalulateImportedtax(item.Price, item.Quantity)
	}

	panic("error calculating tax on invalid item type")
}

func CalculateRawtax(price int, quantity int) float64 {
	/*
		12.5% of the item cost
	*/
	return float64(price) * float64(quantity) * RawItemTaxPercent / Hundread
}

func CalulateManufacturedtax(price int, quantity int) float64 {
	/*
		12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		=> item cost ( 12.5% + 2% (1 + 12.5%))
	*/
	return float64(price) * float64(quantity) * (ManufacturedItemPercentA + ManufacturedItemPercentB*(1+ManufacturedItemPercentA/Hundread)) / Hundread
}

func CalcualteImportDuty(price int, quantity int) float64 {
	/*
		10% import duty on item cost
	*/

	return float64(price) * float64(quantity) * ImportDutyPercent / Hundread
}

func CalulateImportedtax(price int, quantity int) float64 {
	/*
		import duty + surcharge
	*/

	itemPrice := float64(price) * float64(quantity)
	importDuty := CalcualteImportDuty(price, quantity)
	totalPriceWithImportDuty := itemPrice + importDuty

	if totalPriceWithImportDuty <= ImportDutyLevelOne {

		return importDuty + SurchargeLevelOneAmount

	} else if totalPriceWithImportDuty <= ImportDutyLevelTwo {

		return importDuty + SurchargeLevelOneAmount

	} else {

		return importDuty + totalPriceWithImportDuty*SurchargeLevelThreePercent/Hundread

	}

}
