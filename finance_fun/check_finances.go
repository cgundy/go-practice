package main 

 import "fmt"


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

func ListAssets(a Assets) {
    fmt.Println("Cash:", a.Cash)
    fmt.Println("Investments:", a.Investments)
    for _, property := range a.Properties{
        fmt.Println("Properties:", property.AmountPayed)
    }
    
}

 func main() {
    p1 := Property{100000, 20000, 30000, 2000, 3000}
    a := Assets{10000, 50000, 0.01, []*Property{&p1}}
    
    ListAssets(a)
    OneMonth(&a)
    ListAssets(a)
    a.BuyStocks(2000)
    ListAssets(a)
 }
