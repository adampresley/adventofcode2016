package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Direction struct {
	DeltaX    int
	DeltaY    int
	TurnLeft  func() *Direction
	TurnRight func() *Direction
}

type Command struct {
	TurnDirection string
	MoveDistance  int
}

type Vector struct {
	X int
	Y int
}

func ParseInput(input string) *Command {
	input = strings.TrimSpace(input)
	result := &Command{}

	if input[0] == 'L' {
		result.TurnDirection = "left"
	} else {
		result.TurnDirection = "right"
	}

	result.MoveDistance, _ = strconv.Atoi(string(input[1:]))
	return result
}

func GetManhattanDistance(vector1, vector2 Vector) int {
	return int(math.Abs(float64(vector1.X)-float64(vector2.X)) + math.Abs(float64(vector1.Y)-float64(vector2.Y)))
}

func main() {
	var puzzleInput string

	east := &Direction{DeltaX: -1, DeltaY: 0}
	north := &Direction{DeltaX: 0, DeltaY: 1}
	south := &Direction{DeltaX: 0, DeltaY: -1}
	west := &Direction{DeltaX: 1, DeltaY: 0}

	east.TurnLeft = func() *Direction { return north }
	east.TurnRight = func() *Direction { return south }

	north.TurnLeft = func() *Direction { return west }
	north.TurnRight = func() *Direction { return east }

	south.TurnLeft = func() *Direction { return east }
	south.TurnRight = func() *Direction { return west }

	west.TurnLeft = func() *Direction { return south }
	west.TurnRight = func() *Direction { return north }

	//puzzleInput = "R2, L3"
	//puzzleInput = "R2, R2, R2"
	//puzzleInput = "R5, L5, R5, R3"
	puzzleInput = "L4, R2, R4, L5, L3, L1, R4, R5, R1, R3, L3, L2, L2, R5, R1, L1, L2, R2, R2, L5, R5, R5, L2, R1, R2, L2, L4, L1, R5, R2, R1, R1, L2, L3, R2, L5, L186, L5, L3, R3, L5, R4, R2, L5, R1, R4, L1, L3, R3, R1, L1, R4, R2, L1, L4, R5, L1, R50, L4, R3, R78, R4, R2, L4, R3, L4, R4, L1, R5, L4, R1, L2, R3, L2, R5, R5, L4, L1, L2, R185, L5, R2, R1, L3, R4, L5, R2, R4, L3, R4, L2, L5, R1, R2, L2, L1, L2, R2, L2, R1, L5, L3, L4, L3, L4, L2, L5, L5, R2, L3, L4, R4, R4, R5, L4, L2, R4, L5, R3, R1, L1, R3, L2, R2, R1, R5, L4, R5, L3, R2, R3, R1, R4, L4, R1, R3, L5, L1, L3, R2, R1, R4, L4, R3, L3, R3, R2, L3, L3, R4, L2, R4, L3, L4, R5, R1, L1, R5, R3, R1, R3, R4, L1, R4, R3, R1, L5, L5, L4, R4, R3, L2, R1, R5, L3, R4, R5, L4, L5, R2"

	commands := strings.Split(puzzleInput, ", ")
	parsedCommand := &Command{}

	currentDirection := north
	origin := Vector{X: 0, Y: 0}
	currentPosition := Vector{X: 0, Y: 0}
	nextCommand := ""

	for _, nextCommand = range commands {
		fmt.Printf("Next command: %s\n", nextCommand)
		parsedCommand = ParseInput(nextCommand)

		if parsedCommand.TurnDirection == "left" {
			currentDirection = currentDirection.TurnLeft()
		} else {
			currentDirection = currentDirection.TurnRight()
		}

		currentPosition.X += (currentDirection.DeltaX * parsedCommand.MoveDistance)
		currentPosition.Y += (currentDirection.DeltaY * parsedCommand.MoveDistance)
	}

	fmt.Printf("We are now %d blocks away from our starting point.\n", GetManhattanDistance(origin, currentPosition))
}
