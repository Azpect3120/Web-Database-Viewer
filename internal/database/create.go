package database

import (
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Create a new connection to a database and store the details
// in the session.
func CreateConnection(c *gin.Context) {
	var (
		url  string = c.PostForm("db-url")
		name string = c.PostForm("db-conn-name")
	)

	fmt.Println(name, ": ", url)

	session := sessions.Default(c)

	var connections map[string]string

	session_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		connections = make(map[string]string)
	} else {
		if err := json.Unmarshal(session_bytes, &connections); err != nil {
			fmt.Println(err)
		}
	}

	connections[name] = url

	conn_bytes, err := json.Marshal(connections)
	if err != nil {
		fmt.Println(err)
	}

	session.Set("connections", []byte(conn_bytes))
	session.Set("current", name)
	session.Save()

	html := templates.ConnectionsList(connections, name)
	c.String(200, html)
}
