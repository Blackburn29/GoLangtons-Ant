package main

import (
	"fmt"
	"time"
)

var grid [40][40]string
var antx = 10
var anty = 10
var iterations = 500000

var NORTH = 0
var EAST = 1
var SOUTH = 2
var WEST = 3
var WHITE = "█"
var BLACK = " "
var ANT = "@"

var direction = EAST
var nextBlock = WHITE

func main() {
	initGrid()
	clearScreen()
	for i := 0; i < iterations; i++ {
		advanceStep()
		printGrid()
		fmt.Println(i)
		time.Sleep(50 * time.Millisecond)
	}
}

func initGrid() {
	for x, h := range grid {
		for y, _ := range h {
			grid[x][y] = BLACK
		}
	}
}

func printGrid() {
	clearScreen()
	for _, h := range grid {
		for _, cell := range h {
			fmt.Printf(" %s", cell)
		}
		fmt.Println()
	}
}

func advanceStep() {
	switch nextBlock {
	case WHITE:
		advanceWhite()
	case BLACK:
		advanceBlack()
	}

}

func advanceWhite() {
	switch direction {
	case NORTH:
		direction = EAST
	case EAST:
		direction = SOUTH
	case SOUTH:
		direction = WEST
	case WEST:
		direction = NORTH
	}
	grid[antx][anty] = BLACK
	moveForward()
}

func advanceBlack() {
	switch direction {
	case NORTH:
		direction = WEST
	case EAST:
		direction = NORTH
	case SOUTH:
		direction = EAST
	case WEST:
		direction = SOUTH
	}
	grid[antx][anty] = WHITE
	moveForward()
}

//TODO Fix logic for when ant reaches edges of grid.
func moveForward() {
	switch direction {
	case NORTH:
		if anty < len(grid[0]) {
			anty++
		}
	case EAST:
		if anty < len(grid[0]) {
			antx++
		}
	case SOUTH:
		if anty > 0 {
			anty--
		}
	case WEST:
		if anty > 0 {
			antx--
		}
	}
	nextBlock = getBlockType()
	grid[antx][anty] = ANT
}

func getBlockType() string {
	return grid[antx][anty]
}

func clearScreen() {
	fmt.Print("\033[2J")
}
