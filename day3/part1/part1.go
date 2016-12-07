/*
--- Day 3: Squares With Three Sides ---

Now that you can think clearly, you move deeper into the labyrinth of hallways and office furniture that makes up this part of Easter Bunny HQ. This must be a graphic design department; the walls are covered in specifications for triangles.

Or are they?

The design document gives the side lengths of each triangle it describes, but... 5 10 25? Some of these aren't triangles. You can't help but mark the impossible ones.

In a valid triangle, the sum of any two sides must be larger than the remaining side. For example, the "triangle" given above is impossible, because 5 + 10 is not larger than 25.

In your puzzle input, how many of the listed triangles are possible?
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInputFile() ([][3]int, error) {
	var err error
	var file *os.File
	result := make([][3]int, 0, 20)
	var fields []string
	var values [3]int

	if file, err = os.Open("../input.txt"); err != nil {
		return result, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields = strings.Fields(line)

		for i, f := range fields {
			values[i], _ = strconv.Atoi(f)
		}

		sort.Ints(values[:])
		result = append(result, values)
	}

	return result, nil
}

func isTriangle(set [3]int) bool {
	if set[0]+set[1] > set[2] {
		return true
	}

	return false
}

func main() {
	input, _ := readInputFile()
	// input := [][3]int{
	// 	[3]int{5, 10, 25},
	// }

	possibleTriangles := 0

	fmt.Printf("Total inputs: %d\n", len(input))

	for _, set := range input {
		fmt.Printf("%v\n", set)
		if isTriangle(set) {
			possibleTriangles++
		}
	}

	fmt.Printf("Total possible triangles: %d\n", possibleTriangles)
}
