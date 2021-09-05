package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


// lbCmd represents the lb command
var lbCmd = &cobra.Command{
	Use:   "lb",
	Short: "View the binod leaderboard.",
	Long: "The `lb` command displays the binod leaderboard.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lb called")
	},
}

func init() {
	rootCmd.AddCommand(lbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
