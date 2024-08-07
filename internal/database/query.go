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

func QueryCurrent(c *gin.Context) {
	query := c.PostForm("sql")
	conn := getConnection(c)

	cols, data, err := queryConnection(query, conn)
	if err != nil {
		c.String(200, templates.ErrorQueryResults(err))
		return
	}

	c.String(200, templates.QueryResults(cols, data))
}

func queryConnection(query, url string) ([]string, []map[string]interface{}, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return []string{}, []map[string]interface{}{}, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return []string{}, []map[string]interface{}{}, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return []string{}, []map[string]interface{}{}, err
	}

	// Create values pointer and value pointer slices
	// We need to use pointers because the Scan function
	// requires a pointer to the value. However, there
	// is no simple way to create an array of pointers.
	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))

	for i := range cols {
		valuePtrs[i] = &values[i]
	}

	// Final data structure to store the results
	// An array of maps, where each map is a row
	// and the keys are the column names.
	var result []map[string]interface{}

	for rows.Next() {
		// Scan the result into the value pointers
		err := rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Println(err)
		}

		// Create a map to store the column data
		row := make(map[string]interface{})
		for i, col := range cols {
			var v interface{}
			val := values[i]

			// Convert the value to a string representation
			// I can't see when this would fail and return
			// something that isn't a string, but it's better
			// to be safe than sorry. However, this might be
			// hard to handle in the frontend.
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}

			row[col] = v
		}

		// Append the row to the result set
		result = append(result, row)
	}

	return cols, result, nil
}

func getConnection(c *gin.Context) (url string) {
	session := sessions.Default(c)
	conn_bytes, ok := session.Get("connections").([]byte)
	curr, ok := session.Get("current").(string)
	if !ok {
		fmt.Println("No current connection")
		return ""
	}

	var connections map[string]string
	if err := json.Unmarshal(conn_bytes, &connections); err != nil {
		fmt.Println(err)
		return ""
	}

	return connections[curr]
}
