package main

import "testing"

func TestAddChallenge(t *testing.T) {
	location := Location{
		Name:        "Cave",
		Description: "A dark and spooky cave.",
		Items: []Item{
			{
				Name:        "Gold",
				Description: "Shiny gold coins.",
			},
		},
		Enemies: []Character{
			{
				Name:        "Skeleton",
				Health:      10,
				AttackPower: 5,
				Inventory: []Item{
					{
						Name:        "Rusty Sword",
						Description: "An old and worn sword.",
					},
				},
			},
		},
		Connections: map[string]*Location{},
		Challenges:  []Challenge{},
	}
	challenge := Challenge{
		Name:        "Clear the Cave",
		Description: "Defeat all the enemies in the cave.",
		Reward: []Item{
			{
				Name:        "Diamond",
				Description: "A valuable and rare gem.",
			},
		},
	}

	location.addChallenge(challenge)

	if len(location.Challenges) != 1 {
		t.Error("Expected location to have 1 challenge")
	}
	if location.Challenges[0].Name != challenge.Name {
		t.Errorf("Expected location's challenge name to be '%v', got '%v'", challenge.Name, location.Challenges[0].Name)
	}
}

func TestCompleteChallenge(t *testing.T) {
    location := Location{
        Name:        "Cave",
        Description: "A dark and spooky cave.",
        Items: []Item{
            {
                Name:        "Gold",
                Description: "Shiny gold coins.",
            },
        },
        Enemies: []Character{
            {
                Name:        "Skeleton",
                Health:      10,
                AttackPower: 5,
                Inventory: []Item{
                    {
                        Name:        "Rusty Sword",
                        Description: "An old and worn sword.",
                    },
                },
            },
        },
        Connections: map[string]*Location{},
        Challenges: []Challenge{
            {
                Name:        "Clear the Cave",
                Description: "Defeat all the enemies in the cave.",
                Reward: []Item{
                    {
                        Name:        "Diamond",
                        Description: "A valuable and rare gem.",
                    },
                },
            },
        },
    }

    player := &Character{
        Name:     "John",
        Health:   100,
        Inventory: []Item{
            {
                Name:        "Health Potion",
                Description: "Restores 50 health points.",
            },
        },
    }

    location.completeChallenge("Clear the Cave", player)

    if len(location.Enemies) != 0 {
        t.Error("Expected all enemies in location to be defeated")
    }

    if len(player.Inventory) != 2 {
        t.Error("Expected player to have an additional reward item")
    }

    found := false
    for _, item := range player.Inventory {
        if item.Name == "Diamond" {
            found = true
            break
        }
    }

    if !found {
        t.Error("Expected player to have the reward item")
    }
}
