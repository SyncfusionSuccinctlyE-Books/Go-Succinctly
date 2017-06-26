// Code listing 33: https://play.golang.org/p/bXq0z1GghK

package main

import (
	"fmt"
)

type Discount struct {
	percent     float32
	promotionId string
}

type ManagersSpecial struct {
	Discount
	extraoff float32
}

func main() {

	normalPrice := float32(99.99)

	januarySale := Discount{15.00, "January"}
	managerSpecial := ManagersSpecial{januarySale, 10.00}

	discountedPrice := januarySale.Calculate(normalPrice)
	managerDiscount := managerSpecial.Calculate(normalPrice)

	fmt.Printf("Original price: $%4.2f\n", normalPrice)
	fmt.Printf("Discount percentage: %2.2f\n", 
                                               januarySale.percent)
	fmt.Printf("Discounted price: $%4.2f\n", discountedPrice)
	fmt.Printf("Manager's special: $%4.2f\n", managerDiscount)
}

func (d Discount) Calculate(originalPrice float32) float32 {
	return originalPrice - (originalPrice / 100 * d.percent)
}

func (ms ManagersSpecial) Calculate(originalPrice float32) float32 {
	return ms.Discount.Calculate(originalPrice) - ms.extraoff
}