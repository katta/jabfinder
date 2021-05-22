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
		districtCode, _ := cmd.Flags().GetString("districtCode")
		age, _ := cmd.Flags().GetInt("age")
		dose, _ := cmd.Flags().GetInt("dose")

		cowin.CheckAvailability(districtCode, &cowin.Filters{
			Age:  age,
			Dose: dose,
		})
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP("districtCode", "d", "", "Numeric district code")
	checkCmd.Flags().IntP("age", "a", 18, "18 or 45 - Age group to find slots for")
	checkCmd.Flags().IntP("dose", "e", 1, "1 or 2 - Dose to filter by")

	checkCmd.MarkFlagRequired("districtCode")
}
