package templates

import (
	"fmt"
)

// Table wrapper definitions
const table_open string = `<table id="query-result" class="min-w-full divide-y divide-gray-200">`
const table_close string = `</table>`

// Header definitions
const table_head_open string = `<thead class="bg-gray-50"><tr>`
const table_head_close string = `</tr></thead>`
const table_head_row string = `<th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">%s</th>`

// Body definitions
const table_body_open string = `<tbody class="bg-white divide-y divide-gray-200">`
const table_body_close string = `</tbody>`
const table_body_row string = `<td class="px-6 py-4 whitespace-nowrap">%v</td>`

// Error message
const query_error_message string = `<p id="query-error" hx-swap-oob="outerHTML" class="text-red-500 py-2 text-sm">Query Error: %s</p>`
const query_error_message_blank string = `<p id="query-error" hx-swap-oob="outerHTML" class="text-red-500 py-2 text-sm hidden"></p>`

func ErrorQueryResults(e error) string {
	return table_open + table_close + fmt.Sprintf(query_error_message, e.Error())
}

func QueryResults(cols []string, rows []map[string]interface{}) string {
	head := generateHead(cols)

	body := table_body_open
	for _, row := range rows {
		body += generateRow(cols, row)
	}

	body += table_body_close

	return table_open + head + body + table_close + query_error_message_blank
}

// Generate the tables head row
func generateHead(cols []string) string {
	html := table_head_open

	for _, col := range cols {
		html += fmt.Sprintf(table_head_row, col)
	}

	return html + table_head_close
}

func generateRow(cols []string, data map[string]interface{}) string {
	row := "<tr>"

	for _, col := range cols {
		row += fmt.Sprintf(table_body_row, data[col])
	}

	return row + "</tr>"
}
