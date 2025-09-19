package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/xuri/excelize/v2"
)

type Record struct {
	Kod    string
	BS     string
	Hunar  string
	Tariff string
	IDK    string
}

func main() {
	f, r, err := pdf.Open("book.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var pdfText string
	totalPage := r.NumPage()

	for i := 11; i <= totalPage; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		text, err := page.GetPlainText(nil)
		if err != nil {
			log.Fatal(err)
		}
		// Merge the text into one line
		pdfText += " " + strings.TrimSpace(text)
	}

	// Regular expressions
	kodRegex := regexp.MustCompile(`\b\d{6}\b`)       //Kod    column
	bsRegex := regexp.MustCompile(`^\d$`)             //Bs     column
	tariffRegex := regexp.MustCompile(`^\d$|^\d-\d$`) //Tarif  column
	idkRegex := regexp.MustCompile(`^\d{4}$`)         // IDK   column

	// Find positions of all Kod
	kodIndices := kodRegex.FindAllStringIndex(pdfText, -1)
	var records []Record

	for i, pos := range kodIndices {
		start := pos[0]
		var end int
		if i+1 < len(kodIndices) {
			end = kodIndices[i+1][0]
		} else {
			end = len(pdfText)
		}

		block := strings.TrimSpace(pdfText[start:end])
		words := strings.Fields(block)

		if len(words) < 2 {
			continue
		}

		rec := Record{}
		rec.Kod = words[0]

		// BS — the first digit after Kod
		if bsRegex.MatchString(words[1]) {
			rec.BS = words[1]
		}

		// Search for Tariff and IDK in the words after BS
		hunarWords := []string{}
		for _, w := range words[2:] {
			if rec.Tariff == "" && tariffRegex.MatchString(w) {
				rec.Tariff = w
			} else if rec.Tariff != "" && rec.IDK == "" && idkRegex.MatchString(w) {
				rec.IDK = w
			} else {
				hunarWords = append(hunarWords, w)
			}
		}

		rec.Hunar = strings.Join(hunarWords, " ")
		records = append(records, rec)
	}

	fmt.Printf("Found %d records\n", len(records))

	// Create Excel file
	excel := excelize.NewFile()
	sheet := "Sheet1"
	excel.NewSheet(sheet)
	rowCounter := 1

	headers := []string{"Kod", "BS", "Hünärleriň ady", "Tariff zaryadynyň gerimi", "IDK kod"}
	for j, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(j+1, rowCounter)
		excel.SetCellValue(sheet, cell, h)
	}

	style, _ := excel.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
	})
	excel.SetCellStyle(sheet, "A1", "E1", style)
	rowCounter++

	for _, rec := range records {
		row := []string{rec.Kod, rec.BS, rec.Hunar, rec.Tariff, rec.IDK}
		for j, col := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, rowCounter)
			excel.SetCellValue(sheet, cell, col)
		}
		rowCounter++
	}

	if err := excel.SaveAs("book.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Готово! Файл сохранён как output.xlsx")
}
