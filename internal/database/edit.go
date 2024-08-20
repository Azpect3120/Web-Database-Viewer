package database

import (
	"encoding/json"
	"fmt"

	"github.com/Azpect3120/Web-Database-Viewer/internal/templates"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func EditConnections(c *gin.Context) {
	session := sessions.Default(c)
	current, ok := session.Get("current").(string)
	connections_bytes, ok := session.Get("connections").([]byte)
	if !ok {
		fmt.Println("No connections found")
	}

	var connections map[string][2]string
	if err := json.Unmarshal(connections_bytes, &connections); err != nil {
		fmt.Println(err)
	}

	for _, conn := range c.PostFormArray("connections") {
		for name, data := range connections {
			if conn == data[0] {
				delete(connections, name)
			}
		}
	}

	for name, data := range connections {
		newName := c.PostForm(data[0])

		if name == newName {
			continue
		}

		delete(connections, name)
		connections[newName] = [2]string{data[0], data[1]}

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
