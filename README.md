# PDF to Excel Parser

Extract data from PDF files and save it to Excel using a YAML config.

---

## 1️⃣ Setup `config.yaml`

_Edit this file to match your own PDF!_

```yaml
input_file: "your_file.pdf"         # Path to your PDF
output_file: "excel_file/output.xlsx"  # Excel output path
start_page: 1
end_page: 0                          # 0 = until the end
headers:
  - Kod
  - BS
  - Hünär
  - Tariff
start_column: "Kod"                  # Marks start of a new record (CRUCIAL!)
text_column: "Hünär"                 # For text not matching any regex
patterns:
  Kod: '\d{6}'                       # Adjust regex for your PDF
  BS: '^\d$'
  Tariff: '^\d{1,2}(-\d{1,2})?$'
```
⚠️ Important:

start_column must match your PDF. Parser treats everything from one start_column value to the next as one Excel row.

Example:
In your PDF, you might have numbers 123456 and 654321 that mark the start of records.
With start_column: "Kod", everything between 123456 and 654321 will become one row in Excel.
If start_column is set incorrectly, the parser won't know where a new row starts.

Update patterns to fit your PDF structure. Each regex should match the format of the data in your PDF.

2️⃣ Notes

Errors like "Error loading" or "Excel write error" indicate issues with config or file paths:

"Error loading" → config file (config.yaml) not found or has wrong format.

"Excel write error" → cannot write Excel file (check output_file path and write permissions).

Any text that doesn't match regex goes into text_column (e.g., "Hünär").
This is, in my opinion, the best way to separate a text column, because it ensures all unmatched text is still captured and won't be lost.
