package calc

func CalcDiscount(price int, voucher string) int {
	var discounts int = 0
	if voucher == "FAZZFOOD50" {
		if price >= 50000 {
			discounts = price * 50 / 100
			if discounts >= 50000 {
				discounts = 50000
			}
		}
	}
	if voucher == "DITRAKTIR60" {
		if price >= 25000 {
			discounts = price * 60 / 100
			if discounts >= 25000 {
				discounts = 30000
			}
		}
	}
	return discounts
}

func Calcfares(distance int) int {
	var fare int = 5000
	if distance > 2 {
		fare = fare + (distance-2)*3000
	}
	return fare
}

func CalcTaxs(price int, tax bool) int {
	var isTax int = 0
	if tax {
		isTax = price * 5 / 100
	}
	return isTax
}