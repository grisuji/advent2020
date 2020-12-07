package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("i.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	group := make(map[byte]int)

	scanner := bufio.NewScanner(file)

	union := 0
	intersection := 0
	all := 0

	for scanner.Scan() { // read file line by line
		line := scanner.Text() // get the content of the next line
		if line == "" {        // new group
			for _, count := range group {
				union++           // count in every case
				if count == all { // count only if all have the same answer
					intersection++
				}
			}
			group = make(map[byte]int) // reset map
			all = 0
			// fmt.Printf("%d ... %d\n", union, intersection)
		} else { // just count the answers of every user
			for i := 0; i < len(line); i++ { // go thriugh the string
				group[line[i]]++
				// fmt.Printf("%c - %d\n", line[i], group[line[i]])
			}
			all++
		}
	}

	fmt.Printf("union: %d    intersection: %d\n", union, intersection)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
