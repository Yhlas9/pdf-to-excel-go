package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"pdf_to_excel/config"
	"pdf_to_excel/parser"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("load error config.yaml:", err)
	}

	text, err := parser.ReadPDF(cfg.InputFile, cfg.StartPage, cfg.EndPage)
	if err != nil {
		log.Fatal("read error PDF:", err)
	}

	records, err := parser.ParseTextToRecords(text, cfg.Headers, cfg.Patterns, cfg.StartColumn, cfg.TextColumn)
	if err != nil {
		log.Fatal("parsing error PDF:", err)
	}

	fmt.Printf("// Found %d records\n", len(records))

    outputDir := filepath.Dir(cfg.OutputFile)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			log.Fatal("error creating folder:", err)
		}
	}

	if err := parser.WriteExcel(cfg.OutputFile, records, cfg.Headers); err != nil {
		log.Fatal("error writing Excel:", err)
	}

	fmt.Println("✅ Done! Saved to:", cfg.OutputFile)
}
