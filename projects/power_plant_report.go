package main

import "fmt"

// Compute the utilisation of the generated power and print the report
func main() {
	var gridLoad = 75.
	var plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	capacity := plantCapacities[0] + plantCapacities[1] +
		plantCapacities[2] + plantCapacities[3] + plantCapacities[4] +
		plantCapacities[5]
	utilisation := gridLoad / capacity

	fmt.Printf("%-15s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-15s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-15s%.1f%%\n", "Utilisation:", utilisation)
}
