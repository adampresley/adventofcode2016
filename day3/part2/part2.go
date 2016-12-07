/*
--- Part Two ---

Now that you've helpfully marked up their design documents, it occurs to you that triangles are specified in groups of three vertically. Each set of three numbers in a column specifies a triangle. Rows are unrelated.

For example, given the following specification, numbers with the same hundreds digit would be part of the same triangle:

101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603
In your puzzle input, and instead reading by columns, how many of the listed triangles are possible?

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
