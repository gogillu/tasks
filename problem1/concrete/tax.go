package concrete

func CalculateRawItemTax(itemPrice int, itemQuantity int) float64 {
	/*
		12.5% of the item cost
	*/
	return float64(itemPrice) * float64(itemQuantity) * 12.5 / 100
}

func CalulateManufacturedItemTax(itemPrice int, itemQuantity int) float64 {
	/*
		12.5% of the item cost + 2% of (item cost + 12.5% of the item cost)
		=> item cost ( 12.5% + 2% (1 + 12.5%))
	*/
	return float64(itemPrice) * float64(itemQuantity) * (12.5 + 2*(1+12.5/100)) / 100
}

func CalulateImportedItemTax(itemPrice int, itemQuantity int) float64 {
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

	costWithoutTax := float64(itemPrice) * float64(itemQuantity)

	if costWithoutTax*110/100 <= 100 {
		return costWithoutTax*10/100 + 5
	} else if costWithoutTax*110/100 <= 200 {
		return costWithoutTax*10/100 + 10
	} else {
		return costWithoutTax * (10.0 / 100.0) * (1.0 + 105.0/100.0)
	}

}
