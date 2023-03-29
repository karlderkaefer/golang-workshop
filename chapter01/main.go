package main

type Item struct {
	Name        string
	Description string
}

type Character struct {
	Name        string
	Health      int
	AttackPower int
}

type Location struct {
	Name        string
	Description string
	Items       []Item
	Enemies     []Character
	Connections map[string]*Location
}

func createForestLocation() *Location {
	// Implement the createForestLocation function here
	return nil
}

func createVillageLocation() *Location {
	// Implement the createVillageLocation function here
	return nil
}
