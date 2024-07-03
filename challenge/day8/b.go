package day8 

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/Amr-Shams/aoc2024/util/gmath"

)

func bCommand()*cobra.Command{
    return &cobra.Command{
        Use: "b",
        Short: "Run the solution for part b",
        Run: func(cmd *cobra.Command, args []string){
            fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
        },
    }
}

func partB(challenge *challenge.Input)int{
    startPos,steps,Go := parseInput(challenge)
    t :=0 
    T :=make(map[int]int)
    for {
        NewPos := make([]string,0)
        for i,_:= range startPos{
            direction := string(steps[t%len(steps)])
            pos:= Go[direction][startPos[i]]
            if pos[len(pos)-1] == byte('Z'){
                T[i] = t+1
                if len(T) == len(startPos){
                    values := make([]int,len(T))
                    for key,v := range T{
                        values[key] = v
                    }
                    return gmath.LCMArray(values)
                }
            }
            NewPos = append(NewPos,pos)
        }
        startPos = NewPos
        t++   
    }   
    return 0
}
