package main

type Property struct{
	Value, DownPayment, AmountPayed, MonthlyPayment, Rent float64
}

func (p *Property) CheckMonthlyPayment() float64 {
	PaymentLeft := p.Value - p.AmountPayed
	if PaymentLeft < p.MonthlyPayment{
		p.MonthlyPayment = PaymentLeft
	}
	return p.MonthlyPayment
}