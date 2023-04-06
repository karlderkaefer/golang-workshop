package main

type Item struct {
	Name        string
	Description string
}

type Character struct {
	Name        string
	Health      int
	AttackPower int
	Inventory   []Item
}

type Challenge struct {
	Name        string
	Description string
	Reward      []Item
}

type Location struct {
	Name        string
	Description string
	Items       []Item
	Enemies     []Character
	Connections map[string]*Location
	Challenges  []Challenge
}

func createWorldMap() map[string]*Location {
	forest := createForestLocation()
	village := createVillageLocation()
	cave := createCaveLocation()
	castle := createCastleLocation()

	forest.Connections = map[string]*Location{"east": village}
	village.Connections = map[string]*Location{"west": forest, "south": cave, "north": castle}
	cave.Connections = map[string]*Location{"north": village}
	castle.Connections = map[string]*Location{"south": village}
	worldMap := map[string]*Location{
		"Forest":  forest,
		"Village": village,
		"Cave":    cave,
		"Castle":  castle,
	}
	return worldMap
}

func createForestLocation() *Location {
	// Use the createForestLocation function from Chapter 1
	return nil
}

func createVillageLocation() *Location {
	// Use the createVillageLocation function from Chapter 1
	return nil
}

func createCaveLocation() *Location {
	// Implement the createCaveLocation function here
	return nil
}

func createCastleLocation() *Location {
	// Implement the createCastleLocation function here
	return nil
}

// addItemToInventory adds an item to the player's inventory.
func (c *Character) addItemToInventory(item Item) {
	// Implement the addItemToInventory function here
}

// removeItemFromInventory removes an item from the player's inventory.
func (c *Character) removeItemFromInventory(item Item) {
	// Implement the removeItemFromInventory function here
}

// addChallenge adds a challenge to the location's list of challenges.
func (l *Location) addChallenge(challenge Challenge) {
	// Implement the addChallenge function here
}

// completeChallenge removes the challenge from the location's list of challenges and grants the reward items to the player.
func (l *Location) completeChallenge(challengeName string, player *Character) {
	// Implement the completeChallenge function here
}
