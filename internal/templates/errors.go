package templates

import "fmt"

const TABLE_TREE_ERROR_BODY string = `
	<li class="py-2">
		<p class="text-xs text-red-500">Error Loading Tables: %s</p>
	</li>
`

const ENUM_TREE_ERROR_BODY string = `
	<li class="py-2">
		<p class="text-xs text-red-500">Error Loading Tables: %s</p>
	</li>
`

// When an error occurs while generating the table tree,
// this function will return the HTML for the error message.
func TableTreeError(err error) string {
	var html string = TABLE_TREE_OPEN
	html += fmt.Sprintf(TABLE_TREE_ERROR_BODY, err.Error())
	return html + TABLE_TREE_CLOSE
}

// When an error occurs while generating the enum tree,
// this function will return the HTML for the error message.
func EnumTreeError(err error) string {
	var html string = ENUM_TREE_OPEN
	html += fmt.Sprintf(ENUM_TREE_ERROR_BODY, err.Error())
	return html + ENUM_TREE_CLOSE
}
