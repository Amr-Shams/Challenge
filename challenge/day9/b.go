package day9 
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/Amr-Shams/aoc2024/util"
    "strings"
)

func bCommand() *cobra.Command{
    return &cobra.Command{
        Use: "b",
        Short: "Run day9 challenge b",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Answer: %v\n", partB(challenge.FromFile()))
        },
    }
}
func partB(challenge *challenge.Input) int {
    result:=0
    for history:= range challenge.Lines(){ 
        list:= util.AtoiList(strings.Fields(history)...)
        result+=solve(list,true)
    }
    return result
}
