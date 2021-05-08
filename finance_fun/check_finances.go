package main 

import (
    "fmt"
	"bufio"
	"os"
    "strings"
    "strconv"
)

var AvailableProperties = map[string]Property {
    "p1": Property{Value: 100000, DownPayment: 20000, MinPayment: 2000, Rent: 3000},
    "p2": Property{Value: 40000, DownPayment: 10000, MinPayment: 1000, Rent: 1500},
    "p3": Property{Value: 20000, DownPayment: 5000, MinPayment: 500, Rent: 800},
    "p4": Property{Value: 200000, DownPayment: 50000, MinPayment: 5000, Rent: 8000},
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
    if (p.DownPayment > a.Cash){
        fmt.Print("You need mo money")
		return
    }
    a.Cash -= p.DownPayment
    p.init()
    a.Properties = append(a.Properties, p)
    delete(AvailableProperties, name)
}

func ListAvailableProperties() {
    for key, property := range AvailableProperties{  
        fmt.Println(key, ":", property)
    }
}

 func main() {
    a := Assets{10000, 50000, 0.01, []Property{}}
    for{
        fmt.Print("p: list available properties, b: buy property, a: list assets, i: invest, s: sell stocks, r: run \n")
        reader := bufio.NewReader(os.Stdin)
	    input, err := reader.ReadString('\n')
        if (err != nil) {
            fmt.Println("An error occured while reading input. Please try again", err)
        }
        input = strings.TrimSuffix(input, "\n")
        if (input == "p"){
            ListAvailableProperties()
        } else if (input == "b"){
            ListAvailableProperties()
            fmt.Print("Enter property you wish to purchase \n")
            input, err := reader.ReadString('\n')
            input = strings.TrimSuffix(input, "\n")
            if (err == nil){
                BuyProperty(&a, &AvailableProperties, input)
            }
        } else if (input == "a"){
            ListAssets(a)
        } else if (input == "i"){
            fmt.Print("Enter amount of money you wish to invest: \n")
            input, err := reader.ReadString('\n')
            input = strings.TrimSuffix(input, "\n")
            if (err == nil){
                amount, err := strconv.ParseFloat(input, 32)
                if (err == nil){
                    a.BuyStocks(amount)
                }
            }
        } else if (input == "s"){
            fmt.Print("Enter amount of stocks you wish to sell: \n")
            input, err := reader.ReadString('\n')
            input = strings.TrimSuffix(input, "\n")
            if (err == nil){
                amount, err := strconv.ParseFloat(input, 32)
                if (err == nil){
                    a.sellStocks(amount)
                }
            }
        } else if (input == "r"){
            OneMonth(&a)
        }
    }
 }
