
package day6

import (
    "fmt"
    "strings"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/util"
)

func bCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "b",
        Short: "Run the solution for day6 a",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
        },
    }
    return cmd
}

func partB(challenge *challenge.Input) int {
    result:=1
    lines:=challenge.LineSlice()
   timeStr := strings.TrimSpace(strings.Split(lines[0], ":")[1])
    distanceStr := strings.TrimSpace(strings.Split(lines[1], ":")[1])
    
    // Join the fields without spaces
    timeJoined := strings.Join(strings.Fields(timeStr), "")
    distanceJoined := strings.Join(strings.Fields(distanceStr), "")
    
    // Convert joined strings to integer slices
    time := util.Atoi(timeJoined)
    distance := util.Atoi(distanceJoined)
    result*=countWays(time,distance)
       return result
}

