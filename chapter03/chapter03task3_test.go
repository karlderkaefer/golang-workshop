package main

import "testing"

func TestAddItemToInventory(t *testing.T) {
	player := Character{
		Name:        "Gopher",
		Health:      100,
		AttackPower: 10,
		Inventory:   []Item{},
	}

	item := Item{
		Name:        "Healing Potion",
		Description: "A potion that restores health.",
	}

	player.addItemToInventory(item)

	if len(player.Inventory) != 1 {
		t.Error("Expected player's inventory to have 1 item")
	}
	if player.Inventory[0].Name != item.Name {
		t.Errorf("Expected player's inventory to contain '%v', got '%v'", item.Name, player.Inventory[0].Name)
	}
}

func TestRemoveItemFromInventory(t *testing.T) {
	player := Character{
		Name:        "Gopher",
		Health:      100,
		AttackPower: 10,
		Inventory: []Item{
			{
				Name:        "Healing Potion",
				Description: "A potion that restores health.",
			},
		},
	}
	item := Item{
		Name:        "Healing Potion",
		Description: "A potion that restores health.",
	}

	player.removeItemFromInventory(item)

	if len(player.Inventory) != 0 {
		t.Error("Expected player's inventory to be empty")
	}
}