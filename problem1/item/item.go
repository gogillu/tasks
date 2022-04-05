package item

type Item struct {
	Name     string
	Price    int
	Quantity int
	Ttype    string
	Tax      float64
}

func (it *Item) Init(name string, price int, quantity int, ttype string) {
	it.Name = name
	it.Price = price
	it.Quantity = quantity
	it.Ttype = ttype
}

func ValidateItem(givenItem Item) bool {

	if givenItem.Name == "" {
		return false
	}

	if givenItem.Price <= 0 {
		return false
	}

	if givenItem.Quantity <= 0 {
		return false
	}

	if !Validatettype(givenItem.Ttype) {
		return false
	}

	return true
}

func Validatettype(giventtype string) bool {
	validttypes := [3]string{"raw", "manufactured", "imported"}
	// fmt.Println(validttypes)
	for _, ttype := range validttypes {
		if ttype == giventtype {
			return true
		}
	}
	return false
}

func UpdateTaxOnItem(item Item) *Item {
	switch item.Ttype {
	case "raw":
		item.Tax = CalculateRawtax(item.Price, item.Quantity)

	case "manufactured":
		item.Tax = CalulateManufacturedtax(item.Price, item.Quantity)

	case "imported":
		item.Tax = CalulateImportedtax(item.Price, item.Quantity)
	}
	// fmt.Println(item)
	return &item
}

func CalculateRawtax(price int, quantity int) float64 {
	/*
		12.5% of the item cost
	*/
	return float64(price) * float64(quantity) * 12.5 / 100
}

func CalulateManufacturedtax(price int, quantity int) float64 {
	/*
		12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		=> item cost ( 12.5% + 2% (1 + 12.5%))
	*/
	return float64(price) * float64(quantity) * (12.5 + 2*(1+12.5/100)) / 100
}

func CalulateImportedtax(price int, quantity int) float64 {
	/*
		10% import duty on item cost
		+
		a surcharge(surcharge is:
			Rs. 5
				if the final cost after applying tax & import duty is up to Rs. 100,
			Rs. 10
				if the cost exceeds 100 and up to 200 and
			5% of the final cost
				if it exceeds 200)
		=> item cost ( 10% + 2% (1 + 12.5%))
	*/

	costWithoutTax := float64(price) * float64(quantity)

	if costWithoutTax*110/100 <= 100 {
		return costWithoutTax*10/100 + 5
	} else if costWithoutTax*110/100 <= 200 {
		return costWithoutTax*10/100 + 10
	} else {
		return costWithoutTax * (10.0 / 100.0) * (1.0 + 105.0/100.0)
	}

}
