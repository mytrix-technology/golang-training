package main

import (
	"github.com/gin-gonic/gin"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/http"
	_ "log"
)

type HelloTool struct {
}

func main() {
	// Create an HTTP transport
	//transport := http.NewHTTPTransport("/mcp")
	//transport.WithAddr(":8080")
	//
	//// Create server with the HTTP transport
	//server := mcp_golang.NewServer(transport)
	//
	//// Register your tools
	//server.RegisterTool("hello", &HelloTool{})
	//
	//// Start the server
	//if err := server.Serve(); err != nil {
	//	log.Fatal(err)
	//}

	// Create a Gin transport
	transport := http.NewGinTransport()

	// Create server with the Gin transport
	server := mcp_golang.NewServer(transport)

	// Register your tools
	server.RegisterTool("hello", &HelloTool{})

	// Set up Gin router
	router := gin.Default()
	router.POST("/mcp", transport.Handler())

	// Start the server
	router.Run(":8080")
}
