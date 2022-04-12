package item

import (
	"fmt"

	"github.com/gogillu/tax-manager/item/enum"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Item struct {
	name     string
	price    int
	quantity int
	typ      enum.ItemType
}

type Invoice struct {
	itm            Item
	tax            float64
	effectivePrice float64
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

func New(name string, price int, quantity int, Type string) (Item, error) {

	itmType, err := enum.ItemTypeString(Type)
	if err != nil {
		return Item{}, fmt.Errorf("error : provided item type value is not valid")
	}

	itm := Item{}
	itm.name = name
	itm.price = price
	itm.quantity = quantity
	itm.typ = itmType

	if err := itm.validate(); err != nil {
		return itm, fmt.Errorf("error : item validation failed")
	}

	return itm, nil
}

func (itm Item) validate() error {
	return validation.ValidateStruct(&itm,
		validation.Field(&itm.price, validation.By(checkNegativeValue)),
		validation.Field(&itm.quantity, validation.By(checkNegativeValue)),
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

func (itm Item) GetInvoice() Invoice {
	return Invoice{
		itm:            itm,
		tax:            itm.CalculateTax(),
		effectivePrice: itm.CalculateTax() + float64(itm.GetPrice()*itm.GetQuantity()),
	}
}

func (itm Item) CalculateTax() float64 {

	var tax float64

	switch itm.typ {
	case enum.Raw:
		tax = float64(itm.price) * float64(itm.quantity) * RawItmTaxRate

	case enum.Manufactured:
		tax = float64(itm.price) * float64(itm.quantity) * (ManufacturedItmTaxRate + ManufacturedItmTaxRate*(1+ManufacturedItmExtraTaxRate))

	case enum.Imported:
		tax = float64(itm.price)*float64(itm.quantity)*ImportDutyRate + getSurcharge(float64(itm.price*itm.quantity)*(1.0+ImportDutyRate))
	}

	return tax
}

func getSurcharge(price float64) float64 {
	if price <= SurchargeCap1MaxAmt {
		return SurchargeAmt1
	} else if price <= SurchargeCap2MaxAmt {
		return SurchargeAmt2
	} else {
		return price * SurchargeRate3
	}
}
