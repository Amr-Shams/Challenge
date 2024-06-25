package day4 

import (
    "fmt"
    "slices"
 "github.com/spf13/cobra"
 "github.com/Amr-Shams/aoc2024/challenge"
 "strings"
)

func aCommand() *cobra.Command{
    return &cobra.Command{
        Use:   "a",
        Short: "Problems for Day 4A",   
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
        },
    }
}

func partA(challenge *challenge.Input) int {
    sum:=0
    for line:=range challenge.Lines(){
      if cnt:=scratchsOff(line);cnt>0{
        sum+=1<<(cnt-1)
      }
    }
    return sum
}
func scratchsOff(card string) int{
    part:=strings.Split(card,":")
    part= strings.Split(part[1],"|")

    winning:=strings.Fields(part[0])
    my:=strings.Fields(part[1])

    var cnt int 
    for _,v:=range my{
        if slices.Contains(winning,v){
            cnt++
        }
    }
    return cnt
}
