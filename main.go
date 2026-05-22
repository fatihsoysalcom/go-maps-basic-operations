package main

import "fmt"

func main() {
	// Declare and initialize a map using make()
	// Keys are strings (city names), values are integers (populations)
	cityPopulations := make(map[string]int)

	fmt.Println("Initial map:", cityPopulations)

	// Add key-value pairs to the map
	cityPopulations["İstanbul"] = 16000000
	cityPopulations["Ankara"] = 5700000
	cityPopulations["İzmir"] = 4400000
	cityPopulations["Bursa"] = 3100000

	fmt.Println("\nMap after adding elements:", cityPopulations)

	// Access a value by its key
	fmt.Println("\nPopulation of İstanbul:", cityPopulations["İstanbul"])

	// Accessing a non-existent key returns the zero value for the value type (0 for int)
	fmt.Println("Population of Antalya (non-existent key):", cityPopulations["Antalya"])

	// Check for key existence using the "comma ok" idiom
	if pop, ok := cityPopulations["Ankara"]; ok {
		fmt.Printf("Ankara population exists: %d\n", pop)
	} else {
		fmt.Println("Ankara population not found.")
	}

	if pop, ok := cityPopulations["Adana"]; ok {
		fmt.Printf("Adana population exists: %d\n", pop)
	} else {
		fmt.Println("Adana population not found.") // This will be printed
	}

	// Update an existing value
	cityPopulations["İstanbul"] = 16100000
	fmt.Println("\nUpdated İstanbul Population:", cityPopulations["İstanbul"])

	// Iterate over the map using a for...range loop
	fmt.Println("\nCity Populations (Iterating over map):")
	for city, population := range cityPopulations {
		fmt.Printf("%s: %d\n", city, population)
	}

	// Delete a key-value pair from the map
	delete(cityPopulations, "Bursa")
	fmt.Println("\nMap after deleting Bursa:", cityPopulations)

	// Deleting a non-existent key or an already deleted key does not cause an error
	delete(cityPopulations, "Bursa")
	fmt.Println("Map after attempting to delete Bursa again (no change):")

	// Get the number of elements in the map
	fmt.Println("\nCurrent map length:", len(cityPopulations))
}