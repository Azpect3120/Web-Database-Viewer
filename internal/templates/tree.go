package templates

import (
	"fmt"
	"sort"
)

// Tree definition
const TREE_OPEN string = `<ul hx-swap-oob="outerHTML" id="database-table-tree" class="space-y-2">`
const TREE_CLOSE string = `</ul>`
const TREE_BODY_TEMPLATE string = `<li>%s</li>`

// Table definition
const TABLE_TEMPLATE string = `
	<button class="w-full text-left text-gray-700 font-medium hover:bg-gray-100 p-2 rounded flex items-center">
		<svg onclick="ToggleFields('%s');" id="icon-%s" class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" transform="rotate(-90)">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9l6 6 6-6"></path>
		</svg>
		<span class="hover:underline" title="Select this table" onclick="ToggleFields('%s');">%s</span>
		<svg class="w-8 h-8 ml-auto p-2 rounded-full hover:bg-gray-300 transition:all duration-150" xmlns="http://www.w3.org/2000/svg" class="ionicon" viewBox="0 0 512 512" onclick="LoadTableQuery('%s');">
			<path d="M464 428L339.92 303.9a160.48 160.48 0 0030.72-94.58C370.64 120.37 298.27 48 209.32 48S48 120.37 48 209.32s72.37 161.32 161.32 161.32a160.48 160.48 0 0094.58-30.72L428 464zM209.32 319.69a110.38 110.38 0 11110.37-110.37 110.5 110.5 0 01-110.37 110.37z"/>
		</svg>
	</button>
	`

// Fields definition
const FIELDS_LIST_OPEN string = `<ul id="fields-%s" class="hidden ml-6 mt-1 space-y-1 text-gray-600">`
const FIELDS_LIST_CLOSE string = `</ul>`
const FIELD_TEMPLATE string = `
	<li>
		<button class="flex items-center" title="Select this field">
			<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7">
				</path>
			</svg>
			<span>%s</span>
		</button>
	</li>
`

// This is not implemented yet
const PRIMARY_KEY string = `<span class="h-1.5 w-1.5 bg-yellow-500 rounded-full mx-2" title="Primary Key"></span>`

// Generate the tree based on the database tables and columns
func TableTree(tree map[string][]string) string {
	html := TREE_OPEN

	var body string
	for _, table := range getSortedKeys(tree) {
		body += fmt.Sprintf(TABLE_TEMPLATE, table, table, table, table, table)
		fields := fmt.Sprintf(FIELDS_LIST_OPEN, table)
		body += fields + generateFields(tree[table]) + FIELDS_LIST_CLOSE
	}

	html += fmt.Sprintf(TREE_BODY_TEMPLATE, body)
	return html + TREE_CLOSE
}

// Using a list of fields, generate the HTML for the fields
func generateFields(fields []string) string {
	var html string
	for _, field := range fields {
		html += fmt.Sprintf(FIELD_TEMPLATE, field)
	}
	return html
}

// Return a list of the keys in a map, sorted alphabetically
func getSortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
