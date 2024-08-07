package database

import (
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ChangeConnection(c *gin.Context) {
	conn := c.PostForm("connected-database")

	// Do something to change the connection

	session := sessions.Default(c)
	conn_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		c.String(200, templates.ConnectionsList(nil, ""))
		return
	}

	var connections map[string]string
	if err := json.Unmarshal(conn_bytes, &connections); err != nil {
		c.String(200, templates.ConnectionsList(nil, ""))
		fmt.Println(err)
		return
	}

	var name string
	for n, c := range connections {
		if c == conn {
			name = n
			break
		}
	}

	session.Set("current", name)
	session.Save()

	c.String(200, templates.ConnectionsList(connections, name))
}
