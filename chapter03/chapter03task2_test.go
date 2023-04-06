package main

import "testing"

func TestCreateCastleLocation(t *testing.T) {
    castleLocation := createCastleLocation()

    if castleLocation.Name != "Castle" {
        t.Errorf("Expected castleLocation name to be 'Castle', but got %s", castleLocation.Name)
    }

    if castleLocation.Description != "A magnificent castle" {
        t.Errorf("Expected castleLocation description to be 'A magnificent castle', but got %s", castleLocation.Description)
    }

    if castleLocation.Items == nil {
        t.Errorf("Expected castleLocation items to be a slice, but got nil")
    }

    if castleLocation.Enemies == nil {
        t.Errorf("Expected castleLocation enemies to be a slice, but got nil")
    }

    if castleLocation.Connections == nil {
        t.Errorf("Expected castleLocation connections to be a map, but got nil")
    }
}