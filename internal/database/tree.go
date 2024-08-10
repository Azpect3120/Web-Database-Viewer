package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/model"
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

	return templates.TableTree(tree)
}

// Generate the tree of the database tables
func generateTree(url string) (map[string][]model.Column, error) {
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return map[string][]model.Column{}, err
	}
	defer conn.Close()

	tree, err := tableList(conn)
	if err != nil {
		return map[string][]model.Column{}, err
	}

	if err := fillColumns(conn, tree); err != nil {
		return map[string][]model.Column{}, err
	}

	return tree, nil
}

// Return a map with the keys being the table names and the values
// being blank which can be later used to store the columns.
func tableList(conn *sql.DB) (map[string][]model.Column, error) {
	rows, err := conn.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE';")
	if err != nil {
		return map[string][]model.Column{}, err
	}
	defer rows.Close()

	tree := make(map[string][]model.Column)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return map[string][]model.Column{}, err
		}
		tree[table] = []model.Column{}
	}

	return tree, nil
}

// Fill the columns of the tables in the tree using the keys found
// in the tableList function.
//
// For now, the only data stored is the
// column name, but in the future this could be expanded to store
// datatype, constraints, primary keys, relationship, etc.
func fillColumns(conn *sql.DB, tree map[string][]model.Column) error {
	var pkey string
	var fkeys []model.ForeignKey
	for table := range tree {
		unique, err := getUniqueColumns(conn, table)
		if err != nil {
			return err
		}

		pk, err := conn.Query(fmt.Sprintf("SELECT kcu.column_name FROM information_schema.table_constraints tc JOIN  information_schema.key_column_usage kcu  ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema WHERE tc.constraint_type = 'PRIMARY KEY' AND tc.table_name = '%s';", table))
		if err != nil {
			return err
		}
		defer pk.Close()
		for pk.Next() {
			if err := pk.Scan(&pkey); err != nil {
				return err
			}
		}

		fk, err := conn.Query(fmt.Sprintf("SELECT tc.table_schema, tc.table_name, kcu.column_name, ccu.table_schema AS foreign_table_schema, ccu.table_name AS foreign_table_name, ccu.column_name AS foreign_column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema JOIN information_schema.constraint_column_usage AS ccu ON ccu.constraint_name = tc.constraint_name AND ccu.table_schema = tc.table_schema WHERE tc.constraint_type = 'FOREIGN KEY' AND tc.table_name = '%s';", table))
		if err != nil {
			return err
		}
		defer fk.Close()
		for fk.Next() {
			var fkey model.ForeignKey
			if err := fk.Scan(new(interface{}), new(interface{}), &fkey.Column, new(interface{}), &fkey.ForeignTable, &fkey.ForeignColumn); err != nil {
				return err
			}
			fkeys = append(fkeys, fkey)
		}

		rows, err := conn.Query(fmt.Sprintf("SELECT column_name, is_nullable, data_type, character_maximum_length FROM information_schema.columns WHERE table_name = '%s';", table))
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var column model.Column
			if err := rows.Scan(&column.Name, &column.Nullable, &column.Type, &column.MaxLength); err != nil {
				return err
			}
			if column.Name == pkey {
				column.PrimaryKey = true
			}
			for _, fkey := range fkeys {
				if column.Name == fkey.Column {
					column.ForeignKey = fkey
				} else {
					column.ForeignKey = model.ForeignKey{}
				}
			}

			for _, u := range unique {
				if column.Name == u {
					column.Unique = true
				}
			}

			tree[table] = append(tree[table], column)
		}
	}

	return nil
}

// Returns a list of the unique columns in a table
func getUniqueColumns(conn *sql.DB, table string) ([]string, error) {
	var cols []string
	rows, err := conn.Query(fmt.Sprintf("SELECT kcu.column_name FROM information_schema.table_constraints AS tc JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name AND tc.table_schema = kcu.table_schema WHERE tc.constraint_type = 'UNIQUE' AND kcu.table_name = '%s';", table))
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var col string
		if err := rows.Scan(&col); err != nil {
			return []string{}, err
		}
		cols = append(cols, col)
	}

	return cols, nil
}
