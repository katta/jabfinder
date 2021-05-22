package cmd

import (
	"github.com/katta/jabfinder/pkg/cowin"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the availability of vaccine",
	Run: func(cmd *cobra.Command, args []string) {
		cowin.CheckAvailability("", "")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP("districtCode", "d", "", "Numeric district code")
	checkCmd.Flags().StringP("age", "a", "18", "Age Limit")
}
