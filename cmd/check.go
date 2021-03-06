package cmd

import (
	"github.com/katta/jabfinder/pkg/cowin"
	"github.com/katta/jabfinder/pkg/models"
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
		vaccine, _ := cmd.Flags().GetString("vaccine")
		notify, _ := cmd.Flags().GetBool("notify")
		date, _ := cmd.Flags().GetString("date")

		cowin.CheckAvailability(&models.Filters{
			Age:          age,
			Dose:         dose,
			DistrictCode: districtCode,
			Vaccine:      vaccine,
			Date:         date,
		}, notify)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP("districtCode", "d", "", "Numeric district code")
	checkCmd.Flags().IntP("age", "a", 18, "18 or 45 - Age group to find slots for")
	checkCmd.Flags().IntP("dose", "e", 1, "1 or 2 - Dose to filter by")
	checkCmd.Flags().StringP("vaccine", "v", "COVISHIELD", "COVISHIELD or COVAXIN - vaccine types available")
	checkCmd.Flags().BoolP("notify", "n", false, "Notifies via email when the slots are available")
	checkCmd.Flags().StringP("date", "t", "", "Starting date in dd-mm-yyyy ( Checks for 7 days from starting date )")

	checkCmd.MarkFlagRequired("districtCode")
}
