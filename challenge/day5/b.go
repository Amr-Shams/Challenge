package day5

import (
	"fmt"
	"sort"
	"github.com/Amr-Shams/aoc2024/challenge"
	"github.com/spf13/cobra"
)

func bCommand()*cobra.Command{
    return &cobra.Command{
        Use:   "b",
        Short: "Problems for Day 5B",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
        },
    }
}


func createNegativeRanges(ranges []mappingRange) []mappingRange {
	// Sort the ranges by the src field
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	var start int
	for i := 0; i < len(ranges); i++ {
		currentRange := ranges[i]
		if currentRange.start > start {
			// Insert a new negative range before the current range
			negativeRange := mappingRange{
				start:   start,
				end:  start,
				offset: currentRange.start - start,
			}
			// Inserting the negative range into the slice at position i
			ranges = append(ranges[:i], append([]mappingRange{negativeRange}, ranges[i:]...)...)
			i++ // Increment i since we've added an element to the slice
		}
		start = currentRange.start+ currentRange.offset
	}
	return ranges
}

func partB(challenge *challenge.Input) int {
	var seeds []int
	var ranges [][]mappingRange

	for section := range challenge.Sections() {
        if len(seeds) == 0 { // First section is seeds
			seeds = parseInput(section)
		} else { // each section is a list of mapping ranges
          ranges = append(ranges, createNegativeRanges(getMappingRanges(section)))
		}
	}
	for _, r:= range ranges{
		for _, range_ := range r{
			fmt.Println(range_)
		}
	}
	return 0 
}

