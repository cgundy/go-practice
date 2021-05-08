package main

import "fmt"

type Assets struct{
	Cash, Investments, Interest float64
	Properties []Property
}

func PayMortgage(a *Assets) {
    for i, property := range a.Properties{
        property.AmountPayed += property.CheckMonthlyPayment()
		if (property.CheckMonthlyPayment() > a.Cash){
			fmt.Print("Can't pay mortgage")
			return
		}
        a.Cash -= property.CheckMonthlyPayment()
		a.Properties[i].AmountPayed = property.AmountPayed
    }
}

func (a *Assets) ApplyInterest() {
	a.Investments += (a.Investments * a.Interest)
}

func (a *Assets) sellStocks(amount float64){
	if amount > a.Investments{
		fmt.Print("You don't have investments")
		return
	}
	a.Cash = a.Cash + amount
	a.Investments -= amount
}

func (a *Assets) BuyStocks(amount float64){
	if amount > a.Cash{
		fmt.Print("You need mo money")
		return
	}
	a.Cash = a.Cash - amount
	a.Investments += amount
}

func ListAssets(a Assets) {
    fmt.Println("Cash:", a.Cash)
    fmt.Println("Investments:", a.Investments)
    for _, property := range a.Properties{
		fmt.Println("Property:", property)
        fmt.Println("Amount Payed:", property.AmountPayed)
    }
}