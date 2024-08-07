package database

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const CONNECTION_SUCCESS string = `
	<span id="connection-status" class="w-3 h-3 bg-green-500 rounded-full"></span>
	<span hx-swap-oob="outerHTML" id="connection-message" class="text-green-500"> Connection successful! </span>
	<span hx-swap-oob="outerHTML" id="db-url-invalid" class="text-xs text-red-500 hidden"></span>
`

const CONNECTION_FAILURE string = `
	<span id="connection-status" class="w-3 h-3 bg-red-500 rounded-full"></span>
	<span hx-swap-oob="outerHTML" id="connection-message" class="text-red-500"> Connection failed! </span>
	<span hx-swap-oob="outerHTML" id="db-url-invalid" class="text-xs text-red-500">%s</span>
`

// Test a connection to a database
func TestConnectionURL(c *gin.Context) {
	var driver string
	switch c.PostForm("db-driver") {
	case "postgresql":
		driver = "postgres"
	default:
		c.String(200, fmt.Sprintf(CONNECTION_FAILURE, "Unsupported driver"))
		return
	}

	// Open connection
	conn, err := sql.Open(driver, c.PostForm("db-url"))
	if err != nil {
		fmt.Println(err)
		c.String(200, CONNECTION_FAILURE)
	}

	// Ping/test connection
	if err := conn.Ping(); err != nil {
		c.String(200, fmt.Sprintf(CONNECTION_FAILURE, err.Error()))
		return
	}

	c.String(200, CONNECTION_SUCCESS)
}
