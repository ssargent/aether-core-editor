package services

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	gamev1 "github.com/ssargent/aether-core-editor/gen/game/v1"
)

func TestWorldService_CreateWorld(t *testing.T) {
	// This is a basic test structure - you'll need a test database for full testing
	
	// Example test that just verifies the method signature
	service := &WorldService{}
	
	req := connect.NewRequest(&gamev1.CreateWorldRequest{
		Code:        "test",
		Name:        "Test World",
		Description: "A test world",
	})
	
	// This would fail without a database connection, but demonstrates the API
	_, err := service.CreateWorld(context.Background(), req)
	
	// In a real test, you'd setup a test database and verify the result
	if err == nil {
		t.Error("Expected error without database connection")
	}
}
