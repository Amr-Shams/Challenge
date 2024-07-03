package day8 

import ( 
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
    "strings"
)

func aCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "a",
        Short: "Run the solution for day 8 part 1",
        Run: func(_ *cobra.Command,_ []string) {
            fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
        },
    }
}
func parseInput(challenge *challenge.Input)([]string,string, map[string]map[string]string){
    Go := map[string]map[string]string{
        "L": make(map[string]string),
        "R": make(map[string]string),
    }
    startPos := make([]string, 0)
    seondSection := false
    steps := ""
    for line := range challenge.Lines() {
        if line == "" {
            seondSection = true
            continue
        }
        if !seondSection {
            steps+=line
        } else{
            //XKM = (FRH, RLM)
            parts := strings.Split(line, " = ")
            parts[1] = strings.ReplaceAll(parts[1], "(", "")
            parts[1] = strings.ReplaceAll(parts[1], ")", "")
            parts[1] = strings.ReplaceAll(parts[1], ",", "")
            left := strings.Split(parts[1], " ")[0]
            right := strings.Split(parts[1], " ")[1]
            if parts[0][len(parts[0])-1] == byte('A'){
                startPos = append(startPos, parts[0])
            }
            Go["L"][parts[0]] = left
            Go["R"][parts[0]] = right
        }
    }
    return startPos,steps,Go
}

func partA(challenge *challenge.Input) int {
    var steps string
    _,steps,Go := parseInput(challenge)    
    t := 0
    pos := "AAA" // Example starting position, replace with actual start if needed
    for pos != "ZZZ" { // Example end position, replace with actual end if needed
       direction := string(steps[t % len(steps)])
        pos = Go[direction][pos]
        t++
    }
    return t
}
