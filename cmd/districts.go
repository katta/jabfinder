/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/katta/jabfinder/pkg/cowin"
	"github.com/spf13/cobra"
)

// districtsCmd represents the districts command
var districtsCmd = &cobra.Command{
	Use:   "districts",
	Short: "Lists the districts in the given state",
	Run: func(cmd *cobra.Command, args []string) {
		stateCode, _ := cmd.Flags().GetInt("stateCode")
		cowin.ListDistricts(stateCode)
	},
}

func init() {
	rootCmd.AddCommand(districtsCmd)

	districtsCmd.Flags().IntP("stateCode", "s", 0, "State code. Use states command to find one")
}
