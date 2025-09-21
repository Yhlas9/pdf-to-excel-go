package parser

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

// ReadPDF reads text from a PDF within the specified page range
func ReadPDF(file string, startPage, endPage int) (string, error) {
	f, r, err := pdf.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	totalPage := r.NumPage()
	if endPage == 0 || endPage > totalPage {
		endPage = totalPage
	}

	var text string
	for i := startPage; i <= endPage; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		p, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		text += p
	}

	// Joining line breaks
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")

	return text, nil
}