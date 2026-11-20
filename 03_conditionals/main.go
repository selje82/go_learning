package main

import (
	"fmt"
)

func main() {
	const maxPallets int = 750
	storageName := "Gøteborg storage"
	storageTemp := 2.3
	iscoldStorage := true
	storedPallets := 200
	spaceUtilization := float64(storedPallets) / float64(maxPallets) * 100
	availablePallets := maxPallets - storedPallets
	palletWeight := 450.5
	totalWeight := float64(storedPallets) * palletWeight
	deliveryPallets := 200

	fmt.Println("Number of pallets free:", availablePallets)
	fmt.Printf("Warehouse Space Utilization: %.2f%%\n", spaceUtilization)

	fmt.Printf("Storage has %d of %d pallets. %d free\n", storedPallets, maxPallets, availablePallets)

	message := fmt.Sprintf("%s has %d of %d pallets. The temperature is: %.1f", storageName, storedPallets, maxPallets, storageTemp)
	fmt.Println(message)

	fmt.Printf("Total weight of %d pallets: %.2fkg\n", storedPallets, totalWeight)

	fmt.Printf("Storage utilization: %.2f\n", spaceUtilization)

	if spaceUtilization > 90 {
		fmt.Println("Warning: storage almost full")
	} else if spaceUtilization > 50 {
		fmt.Println("Storage OK")
	} else {
		fmt.Println("Plenty of space")
	}

	if free := maxPallets - storedPallets; deliveryPallets <= free {
		storedPallets += deliveryPallets
		fmt.Printf("Stored pallets: %d\n", storedPallets)
	} else {
		fmt.Printf("You have %d pallets too much\n", deliveryPallets-free)
	}

	/*if storedPallets+deliveryPallets <= maxPallets {
		storedPallets += deliveryPallets
		fmt.Printf("Stored pallets: %d\n", storedPallets)
	} else {
		overCapacity := deliveryPallets - (maxPallets - storedPallets)
		fmt.Printf("You have %d pallets too much\n", overCapacity)

	}*/

	if iscoldStorage && storageTemp > 4.0 {
		fmt.Printf("Warning! Storage temperature is: %.2f\n", storageTemp)
	} else if !iscoldStorage {
		fmt.Println("Dry storage, no temperature check")
	} else {
		fmt.Printf("Temperature is %.1f.\n", storageTemp)
	}
}
