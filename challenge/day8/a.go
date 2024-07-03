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


func partA(challenge *challenge.Input) int {
    var steps string
    // Create a map of maps, initializing the nested maps
    Go := map[string]map[string]string{
        "L": make(map[string]string),
        "R": make(map[string]string),
    }

    // Iterate over challenge sections
    for section := range challenge.Sections() {
        if len(steps) == 0 {
            steps = section
        }else{
        // Split the section by " = "
        parts := strings.Split(section, " = ")
        start := parts[0]
        leftRight := strings.Split(parts[1], ",")
        left := strings.TrimSpace(leftRight[0][1:])
        right := strings.TrimSpace(leftRight[1][:len(leftRight[1])-1])
        Go["L"][start] = left
        Go["R"][start] = right
        fmt.Printf("Go[%s][%s] = %s\n", "L", start, left)
        }
    }
    
    
    t := 0
   // pos := "AAA" // Example starting position, replace with actual start if needed
    //for pos != "ZZZ" && t<10 { // Example end position, replace with actual end if needed
      //  direction := string(steps[t % len(steps)])
       // pos = Go[direction][pos]
        //t++
   // }

    return t
}
