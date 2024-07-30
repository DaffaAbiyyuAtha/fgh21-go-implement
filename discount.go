package main

import "fmt"

func calcDiscount (price int, voucher string) int {
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

func calcfares(distance int)int{
	var fare int = 5000
	if distance > 2 {
		fare = fare + (distance-2)*3000
	}
	return fare
}

func calcTaxs(price int, tax bool)int{
	var isTax int = 0
	if tax {
		isTax = price*5/100
	}
	return isTax
}

func fazzFood(price int, voucher string, distance int, tax bool) {
	discount := calcDiscount(price, voucher)
	fares := calcfares(distance)
	taxs := calcTaxs(price,tax)

	var subTotal int = price - discount + fares + taxs

	fmt.Printf("Harga :%d\n", price)
	fmt.Printf("Potongan :%d\n", discount)
	fmt.Printf("Biaya Antar :%d\n", fares)
	fmt.Printf("Pajak :%d\n", taxs)
	fmt.Printf("SubTotal :%d\n", subTotal)
	
}

func main() {
	fazzFood(75000, "FAZZFOOD50", 5, true)
}