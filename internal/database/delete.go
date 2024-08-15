package database

import (
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeleteConnections(c *gin.Context) {
	session := sessions.Default(c)
	current, ok := session.Get("current").(string)
	connections_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		fmt.Println("No connections found")
	}

	var connections map[string]string
	if err := json.Unmarshal(connections_bytes, &connections); err != nil {
		fmt.Println(err)
	}

	for _, conn := range c.PostFormArray("connections") {
		for name, url := range connections {
			if conn == url {
				delete(connections, name)
			}
		}
	}

	for name, url := range connections {
		newName := c.PostForm(url)

		if name == newName {
			continue
		}

		delete(connections, name)
		connections[newName] = url

		if name == current {
			session.Set("current", newName)
		}
	}

	connections_bytes, err := json.Marshal(connections)
	if err != nil {
		fmt.Println(err)
	}
	session.Set("connections", connections_bytes)
	session.Save()

	c.String(200, templates.MANAGER_CLOSED)
}
