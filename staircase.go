package main

import "fmt"

// O(2n^2)
func staircase(n int) (levelString string) {
	for level := 0; level < n; level++ {
		// can also use strings.Repeat() for spaces and hashes
		numberOfSpaces := n - 1 - level
		for space := 0; space < numberOfSpaces; space++ {
			levelString += " "
		}
		numberOfHashes := level + 1
		for hash := 0; hash < numberOfHashes; hash++ {
			levelString += "#"
		}
		levelString += "\n"
	}
	fmt.Println(levelString)
	return
}

func main() {
	staircase(6)
}
