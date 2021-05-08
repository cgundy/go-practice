package main

type Property struct{
	Value, DownPayment, MinPayment, Rent, MonthlyPayment, AmountPayed float64
}

func (p *Property) init(){
	p.MonthlyPayment = p.MinPayment
	p.AmountPayed = p.DownPayment
}

func (p *Property) CheckMonthlyPayment() float64 {
	PaymentLeft := p.Value - p.AmountPayed
	if PaymentLeft < p.MonthlyPayment{
		p.MonthlyPayment = PaymentLeft
	}
	return p.MonthlyPayment
}

func (p *Property) IncreasePayment(amount float64){
	p.MonthlyPayment = amount
}