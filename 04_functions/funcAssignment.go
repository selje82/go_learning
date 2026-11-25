package main

import (
	"fmt"
)

func freeSpace(maxPallets, storedPallets int) int {
	return maxPallets - storedPallets
}

func storeCalc(maxPallets, storedPallets int) (int, float64) {
	availableStorage := freeSpace(maxPallets, storedPallets)
	var spaceUtilization float64 = float64(storedPallets) / float64(maxPallets) * 100

	return availableStorage, spaceUtilization
}

func main() {
	availableStorage := freeSpace(100, 50)
	fmt.Printf("Available storage: %d\n", availableStorage)

	freeStorage, spaceUtilization := storeCalc(100, 50)
	fmt.Printf("Available storage: %d\nStorage utilization: %.2f%%\n", freeStorage, spaceUtilization)
	freeStorage2, _ := storeCalc(200, 100)
	fmt.Printf("Available storage: %d\n", freeStorage2)
}
