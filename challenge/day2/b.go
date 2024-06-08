package day2 
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
)
func bCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "b",
        Short: "Problems for Day 2A",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
        },
    }
}


func partB(challenge *challenge.Input) int {
    var sum int 
    for game:= range challenge.Lines(){
        _,power,_ := gameScore(game)
        sum+=power
    }
    return sum
}

