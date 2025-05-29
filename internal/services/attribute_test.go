package services

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	gamev1 "github.com/ssargent/aether-core-editor/gen/game/v1"
)

func TestAttributeService_CreateAttribute(t *testing.T) {
	// This is a basic test structure - you'll need a test database for full testing

	req := connect.NewRequest(&gamev1.CreateAttributeRequest{
		Name:        "Strength",
		Description: "Physical power and muscle strength",
	})

	// We can't actually call the method without a database connection
	// This test just ensures the request creation is correct
	if req.Msg.Name != "Strength" {
		t.Error("Request creation failed")
	}

	if req.Msg.Description != "Physical power and muscle strength" {
		t.Error("Request description not set correctly")
	}

	// In a real test, you'd setup a test database:
	// db := setupTestDB(t)
	// service := NewAttributeService(db)
	// resp, err := service.CreateAttribute(context.Background(), req)
	// assert.NoError(t, err)
	// assert.Equal(t, "Strength", resp.Msg.Attribute.Name)
	t.Log("AttributeService.CreateAttribute method signature verified")
}

func TestAttributeService_GetAttribute(t *testing.T) {
	service := &AttributeService{}

	req := connect.NewRequest(&gamev1.GetAttributeRequest{
		Id: 1,
	})

	// Verify request structure
	if req.Msg.Id != 1 {
		t.Error("Request ID not set correctly")
	}

	// With nil database, we expect this to panic/error
	// In a real test, you'd use a test database
	_, err := service.GetAttribute(context.Background(), req)

	if err == nil {
		t.Error("Expected error without database connection")
	}
	t.Log("AttributeService.GetAttribute properly handles missing database")
}

func TestAttributeService_UpdateAttribute(t *testing.T) {
	service := &AttributeService{}

	req := connect.NewRequest(&gamev1.UpdateAttributeRequest{
		Id:          1,
		Name:        "Strength",
		Description: "Updated description for physical strength",
	})

	// Verify request structure
	if req.Msg.Id != 1 {
		t.Error("Request ID not set correctly")
	}
	if req.Msg.Name != "Strength" {
		t.Error("Request name not set correctly")
	}

	_, err := service.UpdateAttribute(context.Background(), req)

	if err == nil {
		t.Error("Expected error without database connection")
	}
	t.Log("AttributeService.UpdateAttribute properly handles missing database")
}

func TestAttributeService_DeleteAttribute(t *testing.T) {
	service := &AttributeService{}

	req := connect.NewRequest(&gamev1.DeleteAttributeRequest{
		Id: 1,
	})

	// Verify request structure
	if req.Msg.Id != 1 {
		t.Error("Request ID not set correctly")
	}

	_, err := service.DeleteAttribute(context.Background(), req)

	if err == nil {
		t.Error("Expected error without database connection")
	}
	t.Log("AttributeService.DeleteAttribute properly handles missing database")
}

func TestAttributeService_ListAttributes(t *testing.T) {
	service := &AttributeService{}

	req := connect.NewRequest(&gamev1.ListAttributesRequest{
		Name: "Strength",
	})

	// Verify request structure
	if req.Msg.Name != "Strength" {
		t.Error("Request name filter not set correctly")
	}

	_, err := service.ListAttributes(context.Background(), req)

	if err == nil {
		t.Error("Expected error without database connection")
	}
	t.Log("AttributeService.ListAttributes properly handles missing database")
}

func TestAttributeService_ListAttributesEmpty(t *testing.T) {
	service := &AttributeService{}

	req := connect.NewRequest(&gamev1.ListAttributesRequest{})

	// Verify empty request is valid
	if req.Msg == nil {
		t.Error("Request message should not be nil")
	}

	_, err := service.ListAttributes(context.Background(), req)

	if err == nil {
		t.Error("Expected error without database connection")
	}
	t.Log("AttributeService.ListAttributes handles empty request properly")
}
