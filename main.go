package main

import "github.com/karismapa/ama-billing/handler"

func main() {
	server := handler.NewHTTPServer()
	server.Serve()
}
