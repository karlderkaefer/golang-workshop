package main

import (
	"testing"
)

func TestCharacterCreation(t *testing.T) {
	player := Character{Name: "Gopher", Health: 100, AttackPower: 10}

	if player.Name != "Gopher" {
		t.Errorf("Expected character name to be 'Gopher', got %v", player.Name)
	}

	if player.Health != 100 {
		t.Errorf("Expected character health to be 100, got %v", player.Health)
	}

	if player.AttackPower != 10 {
		t.Errorf("Expected character attack power to be 10, got %v", player.AttackPower)
	}
}

func TestItemCreation(t *testing.T) {
	item := Item{Name: "Health Potion", Description: "Restores 50 health points"}

	if item.Name != "Health Potion" {
		t.Errorf("Expected item name to be 'Health Potion', got %v", item.Name)
	}

	if item.Description != "Restores 50 health points" {
		t.Errorf("Expected item description to be 'Restores 50 health points', got %v", item.Description)
	}
}

func TestLocationCreation(t *testing.T) {
	location := Location{Name: "Forest", Description: "A dense and mysterious forest"}

	if location.Name != "Forest" {
		t.Errorf("Expected location name to be 'Forest', got %v", location.Name)
	}

	if location.Description != "A dense and mysterious forest" {
		t.Errorf("Expected location description to be 'A dense and mysterious forest', got %v", location.Description)
	}
}

func TestCreateForestLocation(t *testing.T) {
	forest := createForestLocation()
	if forest.Name != "Forest" {
		t.Errorf("Expected location name to be 'Forest', got %v", forest.Name)
	}
	if len(forest.Items) == 0 {
		t.Error("Expected Forest location to have items")
	}
	if len(forest.Enemies) == 0 {
		t.Error("Expected Forest location to have enemies")
	}
}

func TestCreateVillageLocation(t *testing.T) {
	village := createVillageLocation()
	if village.Name != "Village" {
		t.Errorf("Expected location name to be 'Village', got %v", village.Name)
	}
	if len(village.Items) == 0 {
		t.Error("Expected Village location to have items")
	}
	if len(village.Enemies) != 0 {
		t.Error("Expected Village location to have no enemies")
	}
}

