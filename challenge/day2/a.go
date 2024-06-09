package day2 
import (
    "fmt"
    "strings"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/Amr-Shams/aoc2024/util"
    "github.com/Amr-Shams/aoc2024/util/gmath"
)
func aCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "a",
        Short: "Problems for Day 2A",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
        },
    }
}
const (
    TOTAL_RED = 12
    TOTAL_GREEN = 13
    TOTAL_BLUE = 14
)

func partA(challenge *challenge.Input) int {
    var sum int 
    for game:= range challenge.Lines(){
        id,_,valid := gameScore(game)
        if  valid{
            sum+=id
        }
    }
    return sum
}

func gameScore(game string)(int,int,bool){
    var maxRed,maxGreen,maxBlue int
    valid:=true
    sections := strings.Split(game,":")
    id:= util.Atoi(strings.Split(sections[0]," ")[1])
    draws:= strings.Split(sections[1],";")
    for _,draw:= range draws{
       colors:= strings.Split(draw,",")   
       for _,color:= range colors{
         str:= strings.Fields(color)    
         count:= util.Atoi(str[0])
         switch str[1]{
            case "red":
                maxRed = gmath.Max(maxRed,count)
                valid = valid && count <= TOTAL_RED
            case "green":
                maxGreen = gmath.Max(maxGreen,count)
                valid = valid && count <= TOTAL_GREEN
            case "blue":
                maxBlue = gmath.Max(maxBlue,count)
                valid = valid && count <= TOTAL_BLUE
            default:
                valid = false
            }
        }
    }
    return id,maxRed*maxGreen*maxBlue,valid
}
