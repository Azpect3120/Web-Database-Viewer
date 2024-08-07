package http

import (
	"time"

	"github.com/Azpect3120/Web-Database-Viewer/internal/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Populate the server with routes
func populate(web, api *gin.RouterGroup) {
	web.GET("/view", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	api.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"time":   time.Now().String(),
		})
	})

	api.POST("/query", func(c *gin.Context) {
		sql := c.PostForm("sql")
		c.JSON(200, gin.H{"sql": sql})
	})

	api.POST("/connections/test", database.TestConnectionURL)
	api.POST("/connections", database.CreateConnection)
	api.GET("/connections", func(c *gin.Context) {
		session := sessions.Default(c)
		connections, ok := session.Get("connections").(map[string]string)

		c.JSON(200, gin.H{"okay": ok, "connections": connections})

	})
}
