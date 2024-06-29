package day5

import (
	"fmt"
	"strings"

	"github.com/Amr-Shams/aoc2024/challenge"
	"github.com/Amr-Shams/aoc2024/util"
	"github.com/spf13/cobra"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Problems for Day 5A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

type mappingRange struct {
	start  int
	end    int
	offset int
}

func parseInput(section string) []int {
	parts := strings.Fields(strings.Split(section, ":")[1])
	seeds := make([]int, 0, len(parts))
	for _, part := range parts {
		seeds = append(seeds, util.Atoi(part))
	}
	return seeds
}

func getMappingRanges(section string) []mappingRange {
	parts := strings.Fields(strings.Split(section, ":")[1])
	ranges := make([]mappingRange, 0, len(parts)/3)
	for i := 0; i < len(parts); i += 3 {
		start := util.Atoi(parts[i])
		end := util.Atoi(parts[i+1])
		offset := util.Atoi(parts[i+2])
		ranges = append(ranges, mappingRange{start, end, offset})
	}
	return ranges
}

func getFinalMapping(seed int, ranges [][]mappingRange) int {
	currentSeed := seed
	for _, r := range ranges {
		for _, range_ := range r {
            if currentSeed >= range_.end && currentSeed < range_.end + range_.offset {
                currentSeed = range_.start + (currentSeed - range_.end)
                break
            }
        }
	}
	return currentSeed
}

func partA(challenge *challenge.Input) int {
	var seeds []int
	var ranges [][]mappingRange

	for section := range challenge.Sections() {
        if len(seeds) == 0 { // First section is seeds
			seeds = parseInput(section)
		} else { // each section is a list of mapping ranges
          ranges = append(ranges, getMappingRanges(section))
		}
	}
	newSeeds := make([]int, len(seeds))
	for i, seed := range seeds {
		newSeeds[i] = getFinalMapping(seed, ranges)
	}
	var minNewSeed int
	for _, seed := range newSeeds {
		if minNewSeed == 0 || seed < minNewSeed {
			minNewSeed = seed
		}
	}
	return minNewSeed
}
