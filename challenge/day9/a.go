package day9 
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/Amr-Shams/aoc2024/util"
    "strings"
)

func aCommand() *cobra.Command{
    return &cobra.Command{
        Use: "a",
        Short: "Run day9 challenge a",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Answer: %v\n", partA(challenge.FromFile()))
        },
    }
}
func allZeros(list []int) bool{
    for _, v := range list{
        if v != 0{
            return false
        }
    }
    return true
}
func solve(list []int,part2 bool ) int{
    if allZeros(list){
        return 0
    }
    var D = make([]int,0)
    for i:=0; i<len(list)-1; i++{
        D = append(D,list[i+1]-list[i])
    }
     loc:=len(list)-1
     sign:=1
     if part2{
         loc=0
         sign=-1
        }        
     return list[loc] + sign*solve(D,part2)
}
    
func partA(challenge *challenge.Input) int {
    result:=0
    for history:= range challenge.Lines(){ 
        list:= util.AtoiList(strings.Fields(history)...)
        result+=solve(list,false)
    }
    return result
}
