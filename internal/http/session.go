package http

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

const SESSION_SECRET string = "lqMVy04h8KWll6lCU9XCDbfUsBIjD1eG"

// Initialize the session
// Session is a map of "db name": "connection url"
func (s *Server) initSession() {
	gob.Register(&map[string]string{})
	store := cookie.NewStore([]byte(SESSION_SECRET))
	s.Router.Use(sessions.Sessions("mysession", store))
}
