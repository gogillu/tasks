package item

import (
	"fmt"
	"tax-manager/item/enum"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Item struct {
	name     string
	price    int
	quantity int
	typ      enum.ItemType
}

func (itm Item) GetName() string {
	return itm.name
}

func (itm Item) GetPrice() int {
	return itm.price
}

func (itm Item) GetQuantity() int {
	return itm.quantity
}

func (itm Item) GetType() enum.ItemType {
	return itm.typ
}

func New(name string, price int, quantity int, Type string) Item {

	itemType, err := enum.GetItemType(Type)
	if err != nil {
		panic(err)
	}

	item := Item{}
	item.name = name
	item.price = price
	item.quantity = quantity
	item.typ = itemType

	return item
}

func (item Item) Validate() error {

	return validation.ValidateStruct(&item,
		validation.Field(&item.price, validation.By(checkNegativeValue)),
		validation.Field(&item.quantity, validation.By(checkNegativeValue)),
	)
}

func checkNegativeValue(val interface{}) error {
	switch x := val.(type) {
	case int:
		if x < 0 {
			return fmt.Errorf("error : %d given value is negative", val)
		}
	}
	return nil
}

func (item Item) GetInvoice() {
	fmt.Printf("|%10s |%10d |%9d |%15s |%20f |\n", item.GetName(), item.GetPrice(), item.GetQuantity(), item.GetType(), item.CalculateTax())
}

func (item Item) CalculateTax() float64 {

	var tax float64

	switch item.typ {
	case enum.Raw:
		tax = float64(item.price) * float64(item.quantity) * RawItmTaxRate

	case enum.Manufactured:
		tax = float64(item.price) * float64(item.quantity) * (ManufacturedItmTaxRate + ManufacturedItmTaxRate*(1+ManufacturedItmExtraTaxRate))

	case enum.Imported:
		tax = float64(item.price)*float64(item.quantity)*ImportDutyRate + getSurcharge(float64(item.price*item.quantity)*(1.0+ImportDutyRate))
	}

	return tax
}

func getSurcharge(price float64) float64 {

	if price <= SurchargeCap1Amt {
		return SurchargeSlab1Amt
	} else if price <= SurchargeCap2Amt {
		return SurchargeSlab2Amt
	} else {
		return price * SurchargeSlab3Rate
	}

}
