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

func (v Vector) String() string {
	return fmt.Sprintf("%d:%d", v.X, v.Y)
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

func Visited(visitedLocations map[string]Vector, vector Vector) bool {
	var ok bool
	_, ok = visitedLocations[vector.String()]
	return ok
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

	//puzzleInput = "R8, R4, R4, R8"
	//puzzleInput = "R2, L2, R2, R2, R2"
	puzzleInput = "R5, L2, L1, R1, R3, R3, L3, R3, R4, L2, R4, L4, R4, R3, L2, L1, L1, R2, R4, R4, L4, R3, L2, R1, L4, R1, R3, L5, L4, L5, R3, L3, L1, L1, R4, R2, R2, L1, L4, R191, R5, L2, R46, R3, L1, R74, L2, R2, R187, R3, R4, R1, L4, L4, L2, R4, L5, R4, R3, L2, L1, R3, R3, R3, R1, R1, L4, R4, R1, R5, R2, R1, R3, L4, L2, L2, R1, L3, R1, R3, L5, L3, R5, R3, R4, L1, R3, R2, R1, R2, L4, L1, L1, R3, L3, R4, L2, L4, L5, L5, L4, R2, R5, L4, R4, L2, R3, L4, L3, L5, R5, L4, L2, R3, R5, R5, L1, L4, R3, L1, R2, L5, L1, R4, L1, R5, R1, L4, L4, L4, R4, R3, L5, R1, L3, R4, R3, L2, L1, R1, R2, R2, R2, L1, L1, L2, L5, L3, L1"

	commands := strings.Split(puzzleInput, ", ")
	parsedCommand := &Command{}

	currentDirection := north
	origin := Vector{X: 0, Y: 0}
	currentPosition := Vector{X: 0, Y: 0}
	visited := make(map[string]Vector)

	nextCommand := ""

	visited[origin.String()] = origin

	for _, nextCommand = range commands {
		parsedCommand = ParseInput(nextCommand)

		if parsedCommand.TurnDirection == "left" {
			currentDirection = currentDirection.TurnLeft()
		} else {
			currentDirection = currentDirection.TurnRight()
		}

		for dx := 0; dx < parsedCommand.MoveDistance; dx++ {
			currentPosition.X += currentDirection.DeltaX
			currentPosition.Y += currentDirection.DeltaY

			if Visited(visited, currentPosition) {
				break
			}

			visited[currentPosition.String()] = currentPosition
		}
	}

	fmt.Printf("We are now %d blocks away from our starting point.\n", GetManhattanDistance(origin, currentPosition))
}
