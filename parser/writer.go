package parser

import "github.com/xuri/excelize/v2"

// WriteExcel writes records to Excel
func WriteExcel(file string, records []Record, headers []string) error {
	excel := excelize.NewFile()
	sheet := "Sheet1"
	excel.NewSheet(sheet)
	rowCounter := 1

	// Headers
	for j, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(j+1, rowCounter)
		excel.SetCellValue(sheet, cell, h)
	}

	style, _ := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	excel.SetCellStyle(sheet, "A1", "Z1", style)
	rowCounter++

    // Data
	for _, rec := range records {
		row := make([]string, len(headers))
		for i, h := range headers {
			row[i] = rec[h]
		}
		for j, col := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, rowCounter)
			excel.SetCellValue(sheet, cell, col)
		}
		rowCounter++
	}

	return excel.SaveAs(file)
}
