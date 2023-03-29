# Chapter 2: Control Structures and Functions - Building the World Map and Game Mechanic

In this chapter, we'll implement the game's basic logic using control structures like if statements, for loops, and switch statements. We'll also create functions to handle player actions, such as moving, interacting with objects, and combat.

**Objective**: In this chapter, you will learn about control structures and functions in Golang by implementing the core logic for the text-based adventure game. You will create functions to handle player actions, such as moving, interacting with objects, and combat. Additionally, you will write tests to ensure that the functions work as expected.

## Teaser

In Chapter 2, our fearless Gopher discovers an ancient map hidden deep within the forest, revealing a network of interconnected realms. Guided by newfound knowledge, the Gopher navigates through treacherous terrains, solving intricate puzzles and overcoming challenges along the way. As the journey unfolds, the Gopher encounters strange creatures and uncovers lost artifacts, each interaction offering valuable lessons and shaping the hero's destiny. This action-packed chapter of Gopher's Adventures will test our hero's cunning and courage, as they explore uncharted territories and face formidable foes.

## Task 1: Create the World Map

In this task, you will implement the `createWorldMap` function, which should create a map of string to `*Location` representing the game world. You will use the `createForestLocation()` and `createVillageLocation()` functions you implemented in Chapter 1 to create the locations for the world map. After creating the Forest and Village locations, connect them as follows:

- The Forest location should have an "east" connection to the Village
- The Village location should have a "west" connection to the Forest

## Task 2: Implement the Move Function

In this task, you will implement the `move` function, which takes a direction string and the current location pointer as arguments. Based on the given direction, the function should update the current location and return the new location. If the given direction is not valid, the function should return the current location unchanged and print an error message.

## Task 3: Implement the Interact Function

In this task, you will implement the `interact` function, which takes an action string and the current location pointer as arguments. Based on the action (e.g., "pickup"), it should modify the game state by picking up items, talking to NPCs, or interacting with other elements in the current location.

## Task 4: Implement the Combat Function

In this task, you will implement the `combat` function, which takes two `Character` pointers, representing the player and the enemy, as arguments. It should simulate a combat encounter between the player and the enemy, modifying their health and other attributes as needed to reflect the outcome of the encounter.

## Outro 

Once you have completed these tasks, you will have a solid understanding of Golang's control structures and functions and how they can be used to create the core logic of a text-based adventure game. You will be ready to move on to the next chapter, where you will learn about more advanced topics such as interfaces, error handling, and concurrency.

Upon completing this chapter, you will have created the world map and implemented the core game mechanics, setting the stage for Gopher's adventure to unfold.