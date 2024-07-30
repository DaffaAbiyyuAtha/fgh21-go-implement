package main

import "fmt"

func fazzFood(price int, voucher string, distance int, tax bool) {
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
	var fare int = 5000
	if distance > 2 {
		fare = fare + (distance-2)*3000
	}

	var isTax int = 0
	if tax {
		isTax = isTax + price*5/100
	}

	var subTotal int = price - discounts + fare + isTax

	fmt.Printf("Harga :%d\n", price)
	fmt.Printf("Potongan :%d\n", discounts)
	fmt.Printf("Biaya Antar :%d\n", fare)
	fmt.Printf("Pajak :%d\n", isTax)
	fmt.Printf("SubTotal :%d\n", subTotal)
	
}

func main() {
	fazzFood(25000, "DITRAKTIR60", 1, false)
}