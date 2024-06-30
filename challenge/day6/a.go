
package day6

import (
    "fmt"
    "strings"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/util"
)

func aCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "a",
        Short: "Run the solution for day6 a",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
        },
    }
    return cmd
}
func extractData(lines []string) (time, distanc []string){
    time=strings.Fields(strings.Split(lines[0], ":")[1])
    distanc=strings.Fields(strings.Split(lines[1], ":")[1])
    return time,distanc
}
func partA(challenge *challenge.Input) int {
    result:=1
    lines:=challenge.LineSlice()
    time,distance:=extractData(lines)
    for i:=0;i<len(time);i++{
        t:=util.Atoi(time[i])
        d:=util.Atoi(distance[i])
        result*=countWays(t,d)
    }
       return result
}

func countWays(t,d int) (cnt int){
    for i:=0;i<t;i++{
        if i*(t-i) > d{
            cnt++
        }
    }
    return cnt
}

