package main

import "github.com/Azpect3120/Web-Database-Viewer/internal/http"

func main() {
	http.New("3001").Setup().Start()
}
