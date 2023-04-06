# Chapter 3: Structs, Composition, and Design Patterns - Building the Game World

In Chapter 3, we can expand the game world by adding new locations, items, and challenges. Let's introduce two new locations: a Cave and a Castle. We will create two new functions, `createCaveLocation()` and `createCastleLocation()`, to set up these locations. Additionally, we'll implement an inventory feature for the player, so they can pick up and use items.

**Objective**: In this chapter, you will learn about structs, composition, and design patterns in Golang by creating the game world for the text-based adventure game. You will use the Location and Character structs from Chapter 1 to create more complex game entities, and you will implement common design patterns to make the game more flexible and maintainable.

## Teaser
In Chapter 3, our fearless Gopher delves deeper into the world map, discovering hidden caves and ancient castles filled with treasure and peril. Using their wits and skills, the Gopher must navigate through these complex labyrinths, solving puzzles and defeating enemies to uncover the secrets of the realm. Along the way, they encounter new allies and foes, each with unique abilities and weaknesses. This exciting chapter of Gopher's Adventures will challenge our hero's strategic thinking and creativity, as they explore new frontiers and unravel the mysteries of the game world.

## Task 1: Create the Cave Location

- Implement the `createCaveLocation` function that returns a `*Location` pointer representing the cave. The cave should have the following properties:
  - Name: "Cave"
  - Description: "A dark and spooky cave, filled with dangerous creatures."
  - Items: One `Item` named "Key" with the description "A rusty old key."
  - Enemies: Two `Character`s, a "Goblin" with 20 health and 5 attack power, and a "Troll" with 50 health and 10 attack power.
  - Connections: An empty slice of `Locations`s `make(map[string]*Location)`

## Task 2: Create the Castle Location

- Implement the `createCastleLocation` function that returns a `*Location` pointer representing the castle. The castle should have the following properties:
  - Name: "Castle"
  - Description: "A magnificent castle."
  - Items: An empty slice of `Item`s.
  - Enemies: An empty slice of `Character`s.
  - Connections: An empty slice of `Locations`s `make(map[string]*Location)`

## Task 3: Implement the Inventory Feature

- Implement the `addItemToInventory` function that takes an `Item` and a `*Character` pointer as arguments. It should add the item to the character's inventory.
- Implement the `removeItemFromInventory` function that takes an `Item` and a `*Character` pointer as arguments. It should remove the item from the character's inventory.

## Task 4: Implement the Challenge Feature

- `addChallenge(challenge Challenge)`: This method should add the given Challenge instance to the Challenges slice of the location.
- `completeChallenge(challengeName string)`: This method should remove the Challenge instance with the given Name from the Challenges slice of the location and add its Reward slice to the Inventory slice of the player.
