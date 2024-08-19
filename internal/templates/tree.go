package templates

import (
	"fmt"
	"sort"

	"github.com/Azpect3120/Web-Database-Viewer/internal/model"
)

// Table tree definition
const TABLE_TREE_OPEN string = `<ul hx-swap-oob="outerHTML" id="database-table-tree" class="space-y-2">`
const TABLE_TREE_CLOSE string = `</ul>`
const TABLE_TREE_BODY_TEMPLATE string = `<li>%s</li>`

// Table definition
const TABLE_TEMPLATE string = `
	<button class="w-full text-left text-gray-700 font-medium hover:bg-gray-100 p-2 rounded flex items-center">
		<svg onclick="ToggleFields('%s');" id="icon-%s" class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" transform="rotate(-90)">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9l6 6 6-6"></path>
		</svg>
		<span class="hover:underline w-full" title="Select this table" onclick="ToggleFields('%s');">%s</span>
		<svg class="w-8 h-8 ml-auto p-2 rounded-full hover:bg-gray-300 transition:all duration-150" xmlns="http://www.w3.org/2000/svg" class="ionicon" viewBox="0 0 512 512" onclick="LoadTableQuery('%s');">
			<path d="M464 428L339.92 303.9a160.48 160.48 0 0030.72-94.58C370.64 120.37 298.27 48 209.32 48S48 120.37 48 209.32s72.37 161.32 161.32 161.32a160.48 160.48 0 0094.58-30.72L428 464zM209.32 319.69a110.38 110.38 0 11110.37-110.37 110.5 110.5 0 01-110.37 110.37z"/>
		</svg>
	</button>
	`

// Fields definition
const TABLE_FIELDS_LIST_OPEN string = `<ul id="fields-%s" class="hidden ml-6 mt-1 space-y-1 text-gray-600">`
const TABLE_FIELDS_LIST_CLOSE string = `</ul>`
const TABLE_FIELD_TEMPLATE string = `
	<li>
		<button onclick="LoadTableQueryWithFields('%s', '%s')" class="flex items-center w-full" title="Select this field">
			<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7">
				</path>
			</svg>
			<span>%s</span>
			<span class="text-xs ml-auto">%s</span>
		</button>
	</li>
`

// Enum tree definition
const ENUM_TREE_OPEN string = `<ul hx-swap-oob="outerHTML" id="database-enum-tree" class="space-y-2">`
const ENUM_TREE_CLOSE string = `</ul>`
const ENUM_TREE_BODY_TEMPLATE string = `<li>%s</li>`

// Enum definition
const ENUM_TEMPLATE string = `
	<button class="w-full text-left text-gray-700 font-medium hover:bg-gray-100 p-2 rounded flex items-center">
		<svg onclick="ToggleEnumValues('%s');" id="icon-enum-squeeze" class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" transform="rotate(-90)">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9l6 6 6-6"></path>
		</svg>
		<span class="hover:underline py-1 w-full" onclick="ToggleEnumValues('%s');">%s</span>
	</button>
`

// Enum values definition
const ENUM_VALUES_LIST_OPEN string = `<ul id="enum-values-%s" class="hidden ml-6 mt-1 space-y-1 text-gray-600">`
const ENUM_VALUES_LIST_CLOSE string = `</ul>`
const ENUM_VALUE_TEMPLATE string = `
	<li>
		<div class="flex items-center w-full py-2">
			<svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24"
				xmlns="http://www.w3.org/2000/svg">
				<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7">
				</path>
			</svg>
			<span>%s</span>
		</div>
	</li>
`

// Generate the tree based on the database tables and columns
func TableTree(tree map[string][]model.Column) string {
	html := TABLE_TREE_OPEN

	var body string
	for _, table := range getSortedKeys(tree) {
		body += fmt.Sprintf(TABLE_TEMPLATE, table, table, table, table, table)
		fields := fmt.Sprintf(TABLE_FIELDS_LIST_OPEN, table)
		body += fields + generateFields(table, tree[table]) + TABLE_FIELDS_LIST_CLOSE
	}

	html += fmt.Sprintf(TABLE_TREE_BODY_TEMPLATE, body)
	return html + TABLE_TREE_CLOSE
}

// Using a list of fields, generate the HTML for the fields
func generateFields(table string, fields []model.Column) string {
	var html string
	for _, field := range fields {
		html += fmt.Sprintf(TABLE_FIELD_TEMPLATE, table, field.Name, field.Name, generateType(field))
	}
	return html
}

// Return a list of the keys in a map, sorted alphabetically
func getSortedKeys[T model.Column | string](m map[string][]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

// Generate the string for the type of column base on the column definition
func generateType(col model.Column) string {
	var str string

	if col.PrimaryKey {
		str = `<span class="text-yellow-500">PK</span>, `
	} else if col.ForeignKey.Column != "" {
		str = fmt.Sprintf(`<span class="text-blue-500">FK: %s(%s)</span>, `, col.ForeignKey.ForeignTable, col.ForeignKey.ForeignColumn)
	}

	if col.Nullable == "NO" {
		str += `<span class="text-red-500">R</span>, `
	}

	if col.Unique {
		str += `<span class="text-green-500">U</span>, `
	}

	if col.MaxLength.Valid {
		str += fmt.Sprintf("%s(%d)", col.Type, col.MaxLength.Int64)
	} else {
		str += col.Type
	}

	return str
}

// Generate the HTML string for the enum tree
func EnumTree(enums map[string][]string) string {
	html := ENUM_TREE_OPEN
	var body string

	for _, enum := range getSortedKeys(enums) {
		body += fmt.Sprintf(ENUM_TEMPLATE, enum, enum, enum)
		valuesList := fmt.Sprintf(ENUM_VALUES_LIST_OPEN, enum)
		body += valuesList + generateEnumValues(enums[enum]) + ENUM_VALUES_LIST_CLOSE
	}

	html += fmt.Sprintf(ENUM_TREE_BODY_TEMPLATE, body)

	return html + ENUM_TREE_CLOSE
}

// Convert a list of values into a list of HTML elements
func generateEnumValues(values []string) string {
	var html string
	for _, value := range values {
		html += fmt.Sprintf(ENUM_VALUE_TEMPLATE, value)
	}
	return html
}
