package http

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	Router *gin.Engine
	config cors.Config
}

// Create a new instance of the http.Server struct
func New(port string) *Server {
	gin.SetMode(gin.DebugMode)

	var server *Server = &Server{
		port:   port,
		Router: gin.Default(),
		config: cors.DefaultConfig(),
	}
	server.config.AllowOrigins = []string{"*"}
	server.Router.Use(cors.New(server.config))

	// Session configuration
	gob.Register([]byte{})
	gob.Register(map[string]string{})
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   86400 * 7, // 7 days
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	server.Router.Use(sessions.Sessions("mysession", store))

	return server
}

// Setup the server with the necessary configurations
func (s *Server) Setup() *Server {
	v1 := s.Router.Group("/v1")
	web_g := v1.Group("/web")
	api_g := v1.Group("/api")

	web_g.Static("/static", "./web/static")
	web_g.Static("/assets", "./assets")

	s.Router.LoadHTMLGlob("./web/templates/*.html")

	populate(web_g, api_g)

	return s
}

// Run the server with the port defined on instantiation
func (s *Server) Start() {
	s.Router.Run(fmt.Sprintf(":%s", s.port))
}
