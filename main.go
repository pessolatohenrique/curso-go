package main

import (
	"curso-go/internal/entity"
	"fmt"
)

func main() {
	order, err := entity.NewOrder("123", 50, 10)
	if err == nil {
		order.CalculateFinalPrice()
		fmt.Println(order)
		fmt.Println(order.ID)
		fmt.Println(order.FinalPrice)
	} else {
		fmt.Println(err)
	}

}
