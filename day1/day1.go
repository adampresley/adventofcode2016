package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	North int = 1
	East  int = 2
	South int = 3
	West  int = 4
)

type Command struct {
	TurnDirection string
	MoveDistance  int
}

type Mover struct {
	input       []string
	inputIndex  int
	blocksAway  int
	orientation int
	xPos        int
	xDirection  int
	yPos        int
	yDirection  int
}

func NewMover(input string) *Mover {
	return &Mover{
		input:       strings.Split(input, ","),
		inputIndex:  -1,
		orientation: North,
		xDirection:  -1,
		yDirection:  -1,
	}
}

func (mover *Mover) TurnAndMove(command string) {
	parsedCommand := mover.parseCommand(command)

	if parsedCommand.TurnDirection == "left" {
		mover.TurnLeft()
	} else {
		mover.TurnRight()
	}

	if mover.orientation == North {
		if mover.yPos < 0 {
			mover.yPos -= parsedCommand.MoveDistance
		} else {
			mover.yPos += parsedCommand.MoveDistance
		}
	}

	if mover.orientation == South {
		if mover.yPos < 0 {
			mover.yPos += parsedCommand.MoveDistance
		} else {
			mover.yPos -= parsedCommand.MoveDistance
		}
	}

	if mover.orientation == East {
		if mover.xPos < 0 {
			mover.xPos -= parsedCommand.MoveDistance
		} else {
			mover.xPos += parsedCommand.MoveDistance
		}
	}

	if mover.orientation == West {
		if mover.xPos > 0 {
			mover.xPos -= parsedCommand.MoveDistance
		} else {
			mover.xPos += parsedCommand.MoveDistance
		}
	}

	mover.blocksAway = int(math.Abs(float64(mover.xPos))) + int(math.Abs(float64(mover.yPos)))
}

func (mover *Mover) GetDistanceFromOrigin() int {
	return int(math.Abs(float64(mover.blocksAway)))
}

func (mover *Mover) TurnRight() {
	newOrientation := int(mover.orientation) + 1

	if newOrientation > 4 {
		newOrientation = 1
	}

	mover.orientation = newOrientation
}

func (mover *Mover) TurnLeft() {
	newOrientation := int(mover.orientation) - 1

	if newOrientation < 1 {
		newOrientation = 4
	}

	mover.orientation = newOrientation
}

func (mover *Mover) NextCommand() (string, bool) {
	hasMore := false
	result := ""

	if mover.inputIndex < len(mover.input)-1 {
		mover.inputIndex++

		if mover.inputIndex < len(mover.input)-1 {
			hasMore = true
		}

		result = mover.input[mover.inputIndex]
	}

	return result, hasMore
}

func (mover *Mover) parseCommand(command string) Command {
	command = strings.TrimSpace(command)
	result := Command{}

	if command[0] == 'L' {
		result.TurnDirection = "left"
	} else {
		result.TurnDirection = "right"
	}

	result.MoveDistance, _ = strconv.Atoi(string(command[1]))
	return result
}

func main() {
	var puzzleInput string

	//puzzleInput = "R2, L3"
	//puzzleInput = "R2, R2, R2"
	//puzzleInput = "R5, L5, R5, R3"
	puzzleInput = "L4, R2, R4, L5, L3, L1, R4, R5, R1, R3, L3, L2, L2, R5, R1, L1, L2, R2, R2, L5, R5, R5, L2, R1, R2, L2, L4, L1, R5, R2, R1, R1, L2, L3, R2, L5, L186, L5, L3, R3, L5, R4, R2, L5, R1, R4, L1, L3, R3, R1, L1, R4, R2, L1, L4, R5, L1, R50, L4, R3, R78, R4, R2, L4, R3, L4, R4, L1, R5, L4, R1, L2, R3, L2, R5, R5, L4, L1, L2, R185, L5, R2, R1, L3, R4, L5, R2, R4, L3, R4, L2, L5, R1, R2, L2, L1, L2, R2, L2, R1, L5, L3, L4, L3, L4, L2, L5, L5, R2, L3, L4, R4, R4, R5, L4, L2, R4, L5, R3, R1, L1, R3, L2, R2, R1, R5, L4, R5, L3, R2, R3, R1, R4, L4, R1, R3, L5, L1, L3, R2, R1, R4, L4, R3, L3, R3, R2, L3, L3, R4, L2, R4, L3, L4, R5, R1, L1, R5, R3, R1, R3, R4, L1, R4, R3, R1, L5, L5, L4, R4, R3, L2, R1, R5, L3, R4, R5, L4, L5, R2"

	nextCommand := ""
	hasMore := true

	mover := NewMover(puzzleInput)

	for hasMore {
		nextCommand, hasMore = mover.NextCommand()
		fmt.Printf("Next command: %s\n", nextCommand)
		mover.TurnAndMove(nextCommand)
	}

	fmt.Printf("We are now %d blocks away from our starting point.\n", mover.GetDistanceFromOrigin())
}
