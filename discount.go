package main

import (
	"fazztrack/demo/calc"
	"fmt"
)

func fazzFood(price int, voucher string, distance int, tax bool) {
	discount := calc.CalcDiscount(price, voucher)
	fares := calc.Calcfares(distance)
	taxs := calc.CalcTaxs(price,tax)

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