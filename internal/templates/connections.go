package templates

import "fmt"

const LIST_OPEN string = `<select id="connected-database" name="connected-database" class="mt-1 block p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 text-sm md:text-base">`
const LIST_ITEM string = `<option value="%s"%s>%s</option>`
const LIST_CLOSE string = `</select>`

// Generate a list of connections to display in the drop-down.
// Current connection will be toggled as selected
func ConnectionsList(connections map[string]string, current string) string {
	var html string = LIST_OPEN

	if len(connections) == 0 || connections == nil {
		html += fmt.Sprintf(LIST_ITEM, "", " selected", "No connections")
		return html + LIST_CLOSE
	}

	for name, url := range connections {
		if name == current {
			html += fmt.Sprintf(LIST_ITEM, url, " selected", name)
		} else {
			html += fmt.Sprintf(LIST_ITEM, url, "", name)
		}
	}

	html += LIST_CLOSE
	return html
}
