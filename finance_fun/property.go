package main

type Property struct{
	Value, DownPayment, MinPayment, Rent, MonthlyPayment, AmountPayed float64
}

func (p *Property) init(){
	if p.MonthlyPayment == 0{
		p.MonthlyPayment = p.MinPayment
	}
}

func (p *Property) CheckMonthlyPayment() float64 {
	PaymentLeft := p.Value - p.AmountPayed
	if PaymentLeft < p.MonthlyPayment{
		p.MonthlyPayment = PaymentLeft
	}
	return p.MonthlyPayment
}