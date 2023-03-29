# Chapter 1: Data Types in Golang - Task Description

**Objective**: In this chapter, you will learn about different data types in Golang by creating basic components for a text-based adventure game. You will use basic and composite data types to define game entities like characters, items, and locations. Pre-defined tests will help you to complete this chapter successfully.

## Teaser

In the mystical land of Go, a brave Gopher sets off on an epic adventure to uncover ancient secrets and protect its homeland. As the journey begins, our Gopher hero encounters a dense and mysterious forest, filled with peculiar creatures and hidden treasures. With each step, the Gopher learns more about the magical world and the powerful forces that govern it. The first chapter of Gopher's Adventures will challenge our hero to navigate through the forest, overcome obstacles, and unlock the mysteries that lie ahead.

## Task 1: Create Forest Location

Take a look at the available component in `main.go` and try to understand their relation and the meaning of the fields. To verfiy your progress you can run the tests with 
```bash
go test -v
```

In this task, you will implement the `createForestLocation` function, which should create a `*Location` representing the Forest. The Forest location should have the following attributes:

- Name: "Forest"
- Description: A suitable description for the Forest area
- Items: A slice of `Item` containing at least one item
- Enemies: A slice of `Character` containing at least one enemy
- Connections: An empty map of string to `*Location` (This will be filled in Chapter 2)

## Task 2: Create Village Location

In this task, you will implement the `createVillageLocation` function, which should create a `*Location` representing the Village. The Village location should have the following attributes:

- Name: "Village"
- Description: A suitable description for the Village area
- Items: A slice of `Item` containing at least one item
- Enemies: An empty slice of `Character` (The Village should not have any enemies)
- Connections: An empty map of string to `*Location` (This will be filled in Chapter 2)

## Task 3: Optional - Explore data types further

1. Experiment with other basic data types, such as integers, floating-point numbers, booleans, and strings. Create variables and manipulate their values to understand how they work in Golang.
2. Practice creating and using composite data types like arrays, slices, and maps. Understand how they can be used to store collections of data in your game.

## Outro

Upon completing this chapter, you will have created two essential locations for Gopher's adventure. The Forest and Village locations will be used to build the world map in Chapter 2.