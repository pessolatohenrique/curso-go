package main

import (
	"curso-go/internal/entity"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	order := entity.NewOrder("1", 50, 10)
	order.CalculateFinalPrice()
	fmt.Println(order)
	fmt.Println(order.ID)
	fmt.Println(order.FinalPrice)
}
