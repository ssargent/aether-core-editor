package services

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	gamev1 "github.com/ssargent/aether-core-editor/gen/game/v1"
)

func TestWorldService_CreateWorld(t *testing.T) {
	// This is a basic test structure - you'll need a test database for full testing

	req := connect.NewRequest(&gamev1.CreateWorldRequest{
		Code:        "test",
		Name:        "Test World",
		Description: "A test world",
	})

	// Verify request structure
	if req.Msg.Code != "test" {
		t.Error("Request code not set correctly")
	}
	if req.Msg.Name != "Test World" {
		t.Error("Request name not set correctly")
	}

	// Test with nil service should panic, so we test the panic occurs
	service := &WorldService{}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic with nil database connection")
		} else {
			t.Log("WorldService.CreateWorld properly panics with nil database as expected")
		}
	}()

	// This should panic due to nil database
	service.CreateWorld(context.Background(), req)
}
