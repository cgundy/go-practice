package main 

import "fmt"

var AvailableProperties = map[string]Property {
    "p1": Property{Value: 100000, DownPayment: 20000, MinPayment: 2000, Rent: 3000},
    "p2": Property{Value: 40000, DownPayment: 10000, MinPayment: 1000, Rent: 1500},
    "p3": Property{Value: 20000, DownPayment: 5000, MinPayment: 500, Rent: 800},
    "p4": Property{Value: 200000, DownPayment: 50000, MinPayment: 5000, Rent: 8000},
}

func PayMortgage(a *Assets) {
    for _, property := range a.Properties{
        property.AmountPayed += property.CheckMonthlyPayment()
        a.Cash -= property.CheckMonthlyPayment()
    }
}

func CollectRent(a *Assets){
    for _, property := range a.Properties{
        a.Cash += property.Rent
    }
    
}

func OneMonth(a *Assets){
    PayMortgage(a)
    CollectRent(a)
    a.ApplyInterest()
}


func BuyProperty(a *Assets, ap *map[string]Property, name string){
    p := AvailableProperties[name]
    p.init()
    a.Properties = append(a.Properties, &p)
    delete(AvailableProperties, name)
}

func ListAvailableProperties() {
    for key, property := range AvailableProperties{  
        fmt.Println(key, ":", property)
    }
}

 func main() {
    a := Assets{10000, 50000, 0.01, []*Property{}}
    ListAvailableProperties()
    BuyProperty(&a, &AvailableProperties, "p1")
    ListAssets(a)
    OneMonth(&a)
    ListAssets(a)
    a.BuyStocks(2000)
    ListAssets(a)
    ListAvailableProperties()
 }
