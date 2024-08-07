package database

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateConnection(c *gin.Context) {
	session := sessions.Default(c)
	var (
		url      string = c.PostForm("db-url")
		database string = c.PostForm("db-database")
	)

	connections, ok := session.Get("connections").(map[string]string)
	if !ok {
		fmt.Println("Creating new connections map /internal/database/create.go:19")
		connections = make(map[string]string)
	}

	connections[database] = url

	session.Set("connections", connections)
	err := session.Save()
	if err != nil {
		fmt.Println("Failed to save session /internal/database/create.go:29")
		fmt.Println(err)
	}
}
