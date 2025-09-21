package parser

import (
	"regexp"
	"strings"

)

type Record map[string]string


// ParseTextToRecords parses PDF text into records based on the config
func ParseTextToRecords(text string, headers []string, patterns map[string]string, startColumn, textColumn string) ([]Record, error) {
	regexMap := make(map[string]*regexp.Regexp)
	for col, pat := range patterns {
		regexMap[col] = regexp.MustCompile(pat)
	}

	var records []Record

	// // Start of a new record
	startRe, ok := regexMap[startColumn]
	if !ok {
		return nil, nil
	}
	indices := startRe.FindAllStringIndex(text, -1)

	for i, pos := range indices {
		start := pos[0]
		end := len(text)
		if i+1 < len(indices) {
			end = indices[i+1][0]
		}

		recordText := strings.TrimSpace(text[start:end])
		fields := strings.Fields(recordText)
		if len(fields) == 0 {
			continue
		}

		record := make(Record)

		for _, f := range fields {
			matched := false
			for col, re := range regexMap {
				if record[col] == "" && re.MatchString(f) {
					record[col] = f
					matched = true
					break
				}
			}
			if !matched {
				// // Any word that doesn't match any regex → goes into textColumn
				if record[textColumn] != "" {
					record[textColumn] += " "
				}
				record[textColumn] += f
			}
		}

		records = append(records, record)
	}

	return records, nil
}
