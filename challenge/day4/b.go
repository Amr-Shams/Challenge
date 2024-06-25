package day4 
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
)

func bCommand()*cobra.Command{
    return &cobra.Command{
        Use:   "b",
        Short: "Problems for Day 4B",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
        },
    }
}

func partB(challenge *challenge.Input) int {
    cards:=challenge.LineSlice()
    counts:=make([]int,len(cards))
    for i := range counts {
        counts[i] = 1
    }
    for i,card:=range cards{
        if cnt:=scratchsOff(card);cnt>0{
            for j:=i+1;j<=cnt+i;j++{
                counts[j]+=counts[i]
            }
        }
    }
    var sum int 
    for _,v:=range counts{
        sum+=v
    }

    return sum
}

