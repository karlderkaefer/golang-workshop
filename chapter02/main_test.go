package main

import (
	"testing"
)

// existing tests from chapter 1

func TestMove(t *testing.T) {
	worldMap := createWorldMap()
	currentLocation := worldMap["Forest"]

	currentLocation = move("north", currentLocation)
	if currentLocation.Name != "Forest" {
		t.Errorf("Expected location name to still be 'Forest', got %v", currentLocation.Name)
	}

	currentLocation = move("east", currentLocation)
	if currentLocation.Name != "Village" {
		t.Errorf("Expected location name to be 'Village', got %v", currentLocation.Name)
	}
}

func TestInteract(t *testing.T) {
	worldMap := createWorldMap()
	currentLocation := worldMap["Forest"]

	interact("pickup", currentLocation)
	if len(currentLocation.Items) != 0 {
		t.Error("Expected no items left in the location after picking up the health potion")
	}
}

func TestCombat(t *testing.T) {
	worldMap := createWorldMap()
	player := &Character{Name: "Gopher", Health: 100, AttackPower: 10}
	enemy := &worldMap["Forest"].Enemies[0]

	combat(player, enemy)

	if player.Health <= 0 {
		t.Error("Expected player to be alive after combat")
	}

	if enemy.Health > 0 {
		t.Error("Expected enemy to be defeated after combat")
	}
}
