package main

type Character struct {
	Name        string
	Health      int
	AttackPower int
}

type Item struct {
	Name        string
	Description string
}

type Location struct {
	Name        string
	Description string
	Items       []Item
	Enemies     []Character
	// is a map of directions to locations
	// example of a valid connection: "east": &Location{...}
	Connections map[string]*Location
}

// existing function from chapter one

func createForestLocation() *Location {
	// Implement the createForestLocation function here
	return nil
}

func createVillageLocation() *Location {
	// Implement the createVillageLocation function here
	return nil
}

func createWorldMap() map[string]*Location {
	forest := createForestLocation()
	village := createVillageLocation()

	// TODO add connections between locations

	return map[string]*Location{
		"Forest":  forest,
		"Village": village,
	}
}


// move takes a direction string and the current location pointer as arguments.
// It checks if there is a valid connection in the given direction using the
// Connections map inside the Location struct, and returns the updated location.
// If there is no valid connection in the given direction, it returns the
// current location unchanged and prints an error message.
func move(direction string, currentLocation *Location) *Location {
	// Implement the move function here
	return nil
}

// interact takes an action string and the current location pointer as arguments.
// Based on the action (e.g., "pickup"), it modifies the game state by picking
// up items, talking to NPCs, or interacting with other elements in the current location.
func interact(action string, currentLocation *Location) {
	// Implement the interact function here
}

// combat takes two Character pointers, representing the player and the enemy, as arguments.
// It simulates a combat encounter between the player and the enemy, modifying their health
// and other attributes as needed to reflect the outcome of the encounter.
func combat(player, enemy *Character) {
	// Implement the combat function here
}


