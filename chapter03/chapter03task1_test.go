package main

import "testing"
import "reflect"

// ... Chapter 1 and 2 tests ...

func TestCreateWorldMap(t *testing.T) {
	worldMap := createWorldMap()
	forest := createForestLocation()
	village := createVillageLocation()
	castle := createCastleLocation()
	cave := createCaveLocation()

	// Check connections
	if worldMap["Forest"].Connections["east"].Name != village.Name {
		t.Errorf("Expected Forest to connect east to Village, got %v", worldMap["Forest"].Connections["east"])
	}
	if worldMap["Village"].Connections["west"].Name != forest.Name {
		t.Errorf("Expected Village to connect west to Forest, got %v", worldMap["Village"].Connections["west"])
	}
	if worldMap["Village"].Connections["south"].Name != cave.Name {
		t.Errorf("Expected Village to connect south to Cave, got %v", worldMap["Village"].Connections["south"])
	}
	if worldMap["Village"].Connections["north"].Name != castle.Name {
		t.Errorf("Expected Village to connect north to Castle, got %v", worldMap["Village"].Connections["north"])
	}
	if worldMap["Castle"].Connections["south"].Name != village.Name {
		t.Errorf("Expected Cave to connect south to Forest, got %v", worldMap["Cave"].Connections["south"])
	}
    if worldMap["Cave"].Connections["north"].Name != village.Name {
        t.Errorf("Expected Cave to connect north to Village, got %v", worldMap["Cave"].Connections["north"])
    }

	// Check locations
	if worldMap["Forest"].Name != "Forest" {
		t.Errorf("Expected Forest location to have name 'Forest', got %v", worldMap["Forest"].Name)
	}
	if worldMap["Village"].Name != "Village" {
		t.Errorf("Expected Village location to have name 'Village', got %v", worldMap["Village"].Name)
	}
	if worldMap["Castle"].Name != "Castle" {
		t.Errorf("Expected Castle location to have name 'Castle', got %v", worldMap["Castle"].Name)
	}
	if worldMap["Cave"].Name != "Cave" {
		t.Errorf("Expected Cave location to have name 'Cave', got %v", worldMap["Cave"].Name)
	}

	// Check items and enemies
	if len(worldMap["Forest"].Items) != 1 {
		t.Errorf("Expected Forest location to have 1 item, got %d", len(worldMap["Forest"].Items))
	}
	if len(worldMap["Forest"].Enemies) != 1 {
		t.Errorf("Expected Village location to have 1 enemy, got %d", len(worldMap["Forest"].Enemies))
	}
	if len(worldMap["Castle"].Items) != 0 {
		t.Errorf("Expected Castle location to have 0 items, got %d", len(worldMap["Castle"].Items))
	}
	if len(worldMap["Cave"].Enemies) != 2 {
		t.Errorf("Expected Cave location to have 2 enemies, got %d", len(worldMap["Cave"].Enemies))
	}
}

func TestCreateCaveLocation(t *testing.T) {
	cave := createCaveLocation()
	if cave == nil {
		t.Errorf("Expected cave location, but got nil")
	}

	// Check location name and description
	if cave.Name != "Cave" {
		t.Errorf("Expected cave name to be 'Cave', but got '%s'", cave.Name)
	}
	if cave.Description != "A dark and spooky cave, filled with dangerous creatures." {
		t.Errorf("Expected cave description to be 'A dark and spooky cave, filled with dangerous creatures.', but got '%s'", cave.Description)
	}

	// Check items in location
	expectedItems := []Item{
		{Name: "Key", Description: "A rusty old key."},
	}
	if !reflect.DeepEqual(cave.Items, expectedItems) {
		t.Errorf("Expected items in cave to be '%v', but got '%v'", expectedItems, cave.Items)
	}

	// Check enemies in location
	expectedEnemies := []Character{
		{Name: "Goblin", Health: 20, AttackPower: 5},
		{Name: "Troll", Health: 50, AttackPower: 10},
	}
	if !reflect.DeepEqual(cave.Enemies, expectedEnemies) {
		t.Errorf("Expected enemies in cave to be '%v', but got '%v'", expectedEnemies, cave.Enemies)
	}

}


