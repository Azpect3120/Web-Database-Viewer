package templates

import "fmt"

// List item templates
const LIST_OPEN string = `<select id="connected-database" name="connected-database" class="mt-1 block p-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500 text-sm md:text-base" title="Change the connection to query.">`
const LIST_ITEM string = `<option value="%s"%s>%s</option>`
const LIST_CLOSE string = `</select>`

const TREE_VIEW_NAME string = `<span hx-swap-oob="outerHTML" id="database-name-tree">%s</span>`

// Generate a list of connections to display in the drop-down.
// Current connection will be toggled as selected
func ConnectionsList(connections map[string]string, current string) string {
	var html string = LIST_OPEN

	if len(connections) == 0 || connections == nil {
		html += fmt.Sprintf(LIST_ITEM, "", " selected", "No connections")
		return html + LIST_CLOSE
	}

	html += fmt.Sprintf(LIST_ITEM, connections[current], " selected", current)
	for _, name := range getSortedKeys(connections) {
		if name == current {
			continue
		} else {
			html += fmt.Sprintf(LIST_ITEM, connections[name], "", name)
		}
	}

	treeName := fmt.Sprintf(TREE_VIEW_NAME, current)

	html += LIST_CLOSE
	return html + treeName
}
