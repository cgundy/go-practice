package main

import "fmt"

type Assets struct{
	Cash, Investments, Interest float64
	Properties []*Property
}

func (a *Assets) ApplyInterest() {
	a.Investments += (a.Investments * a.Interest)
}

func (a *Assets) sellStocks(amount float64){
	a.Cash = a.Cash + amount
	a.Investments -= amount
}

func (a *Assets) BuyStocks(amount float64){
	a.Cash = a.Cash - amount
	a.Investments += amount
}

func ListAssets(a Assets) {
    fmt.Println("Cash:", a.Cash)
    fmt.Println("Investments:", a.Investments)
    for _, property := range a.Properties{
        fmt.Println("Properties:", property.AmountPayed)
    }
}