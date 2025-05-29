package main

import (
	"log"
	"net/http"
	"os"

	"connectrpc.com/grpcreflect"
	"github.com/ssargent/aether-core-editor/gen/aethercore/v1/aethercorev1connect"
	"github.com/ssargent/aether-core-editor/internal/database"
	"github.com/ssargent/aether-core-editor/internal/services"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// Initialize database connection
	db, err := database.NewFromEnv()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create services
	worldService := services.NewWorldService(db)
	attributeService := services.NewAttributeService(db)

	// Create HTTP multiplexer
	mux := http.NewServeMux()

	// Register services
	path, handler := aethercorev1connect.NewWorldServiceHandler(worldService)
	mux.Handle(path, handler)

	attributePath, attributeHandler := aethercorev1connect.NewAttributeServiceHandler(attributeService)
	mux.Handle(attributePath, attributeHandler)

	// Add gRPC reflection for development
	reflector := grpcreflect.NewStaticReflector(
		"aethercore.v1.WorldService",
		"aethercore.v1.AttributeService",
		// Add other services here as they're implemented
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Add CORS middleware for web clients
	corsHandler := addCORS(mux)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server with HTTP/2 support
	log.Printf("Starting server on port %s", port)
	log.Printf("gRPC-Web and Connect clients can access services at http://localhost:%s", port)

	err = http.ListenAndServe(
		":"+port,
		// Use h2c for HTTP/2 without TLS (development only)
		h2c.NewHandler(corsHandler, &http2.Server{}),
	)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// addCORS adds CORS headers for web clients
func addCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Connect-Protocol-Version, Connect-Timeout-Ms")
		w.Header().Set("Access-Control-Expose-Headers", "Connect-Protocol-Version, Connect-Timeout-Ms")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
