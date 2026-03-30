package main

import (
	"fmt"
	"strings"
)

var productPrice = map[string]float64{
	"TSHIRT": 50.00,
	"MUG":    34.00,
	"CUP":    60.00,
	"PHONES": 1000,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := productPrice[itemCode]
	if !found {
		if strings.HasSuffix(itemCode, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemCode, "_SALE")
			basePrice, found = productPrice[originalItemCode]

			if found {
				salePrice := basePrice * 0.90
				fmt.Printf("-- Item %s  (!Sale Original : $%.2f,  Sales Price: %.2f)\n ", originalItemCode, basePrice, salePrice)

				return salePrice, true
			}
		}
		fmt.Printf("- Items: %s (Product not found)\n", itemCode)
		return 0.0, false
	}

	return basePrice, true

}

func main() {

	fmt.Println("===================PRODUCT PRICES=============")
	orderItem := []string{
		"TSHIRT", "MUG_SALE", "CUP", "PHONES",
	}

	var subTotal float64

	fmt.Println("===============PROCESSING ORDER===========")

	for _, item := range orderItem {
		price, found := calculateItemPrice(item)

		if found {
			subTotal += price
		}
	}
	fmt.Printf("Total price: %.2f\n", subTotal)
}
