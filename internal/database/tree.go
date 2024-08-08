package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func TableTree(c *gin.Context) string {
	session := sessions.Default(c)
	connections_bytes, ok := session.Get("connections").([]byte)
	current, ok := session.Get("current").(string)
	if !ok {
		fmt.Println("No connections found")
		return ""
	}

	var connections map[string]string
	if err := json.Unmarshal(connections_bytes, &connections); err != nil {
		fmt.Println(err)
		return ""
	}

	url := connections[current]

	tree, err := generateTree(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Printf("%+v\n", tree)

	return templates.TableTree(tree)
}

// Generate the tree of the database tables
func generateTree(url string) (map[string][]string, error) {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return map[string][]string{}, err
	}
	defer conn.Close()

	tree, err := tableList(conn)
	if err != nil {
		return map[string][]string{}, err
	}

	if err := fillColumns(conn, tree); err != nil {
		return map[string][]string{}, err
	}

	return tree, nil
}

// Return a map with the keys being the table names and the values
// being blank which can be later used to store the columns.
func tableList(conn *sql.DB) (map[string][]string, error) {
	rows, err := conn.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
	if err != nil {
		return map[string][]string{}, err
	}
	defer rows.Close()

	tree := make(map[string][]string)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return map[string][]string{}, err
		}
		tree[table] = []string{}
	}

	return tree, nil
}

// Fill the columns of the tables in the tree using the keys found
// in the tableList function.
//
// For now, the only data stored is the
// column name, but in the future this could be expanded to store
// datatype, constraints, primary keys, relationship, etc.
func fillColumns(conn *sql.DB, tree map[string][]string) error {
	for table := range tree {
		rows, err := conn.Query(fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name = '%s';", table))
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var column string
			if err := rows.Scan(&column); err != nil {
				return err
			}
			tree[table] = append(tree[table], column)
		}
	}

	return nil
}
