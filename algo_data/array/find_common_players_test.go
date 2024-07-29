package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWhoPlaysInBothSports(t *testing.T) {
	footBallPlayers := []Players{
		{FirstName: "Jill", LastName: "Huang", Team: "Gators"},
		{FirstName: "Janko", LastName: "Barton", Team: "Sharks"},
		{FirstName: "Wand", LastName: "Vakulskas", Team: "Sharks"},
		{FirstName: "Jill", LastName: "Moloney", Team: "Gators"},
		{FirstName: "Luuk", LastName: "Watkins", Team: "Gators"},
	}

	basquetBallPlayers := []Players{
		{FirstName: "Hanzla", LastName: "Radosti", Team: "32ers"},
		{FirstName: "Tina", LastName: "Watkins", Team: "Barleycorns"},
		{FirstName: "Alex", LastName: "Patel", Team: "32ers"},
		{FirstName: "Jill", LastName: "Huang", Team: "Barleycorns"},
		{FirstName: "Wand", LastName: "Vakulskas", Team: "Barleycorns"},
	}

	expected := []string{
		"Jill Huang",
		"Wand Vakulskas",
	}

	got := FindWhoPlaysInBothSports(footBallPlayers, basquetBallPlayers)

	assert.ElementsMatch(t, got, expected)
}
