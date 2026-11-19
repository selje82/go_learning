package main

import (
	"fmt"
)

func main() {
	var storage string = "Mediamarkt Frankfurt"
	var nrPallets int = 45
	var storageOpen bool = false
	var temperature float64 = 1.2

	fmt.Println("Storage:", storage)
	fmt.Println("Number of pallets:", nrPallets)
	fmt.Println("Is storage open?", storageOpen)
	fmt.Println("Temperature:", temperature)

}
