package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func setupTable(table *tablewriter.Table, header []string, rows [][]string, footer []string) *tablewriter.Table {
	table.SetHeader(header)
	var headerColors []tablewriter.Colors
	var columnColors []tablewriter.Colors

	for index := 0; index < len(header); index++ {
		headerColors = append(headerColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor})
		columnColors = append(columnColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor})
	}

	table.SetHeaderColor(headerColors...)
	table.SetColumnColor(columnColors...)
	for _, row := range rows {
		table.Append(row)
	}

	if footer != nil {
		var footerColors []tablewriter.Colors
		for index := 0; index < len(footer); index++ {
			footerColors = append(footerColors, tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor})

		}
		table.SetFooter(footer)
		table.SetFooterColor(footerColors...)

	}

	return table
}

func Render(header []string, rows [][]string, footer []string, showRowLine bool) {
	table := tablewriter.NewWriter(os.Stdout)
	setupTable(table, header, rows, footer)
	table.SetRowLine(showRowLine)
	totalColumns := len(header)
	table.SetAutoMergeCellsByColumnIndex([]int{totalColumns - 2, totalColumns - 1})
	table.Render()
}
