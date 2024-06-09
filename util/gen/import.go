package gen

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "gen",
		Short: "Generate input files for all days",
	}

	day.AddCommand(genCommand())
	root.AddCommand(day)
}
