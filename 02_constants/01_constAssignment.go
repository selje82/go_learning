package main

import (
	"fmt"
)

func main() {
	const maxPallets int = 750
	storageName := "Gøteborg storage"
	storageTemp := 2.3
	storedPallets := 375
	spaceUtilization := float64(storedPallets) / float64(maxPallets) * 100
	availablePallets := maxPallets - storedPallets
	palletWeight := 450.5
	totalWeight := float64(storedPallets) * palletWeight

	fmt.Println("Number of pallets free:", availablePallets)
	fmt.Printf("Warehouse Space Utilization: %.2f%%\n", spaceUtilization)

	fmt.Printf("Storage has %d of %d pallets. %d free\n", storedPallets, maxPallets, availablePallets)

	message := fmt.Sprintf("%s has %d of %d pallets. The temperature is: %.1f", storageName, storedPallets, maxPallets, storageTemp)
	fmt.Println(message)

	fmt.Printf("Total weight of %d pallets: %.2fkg\n", storedPallets, totalWeight)

}
