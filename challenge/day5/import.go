package day5 

import "github.com/spf13/cobra"

func AddCommandsTo(root* cobra.Command){
    day:= &cobra.Command{
        Use: "5",
        Short: "Run day 5 solutions",
    }
    day.AddCommand(aCommand())
    day.AddCommand(bCommand())
    root.AddCommand(day)
}
