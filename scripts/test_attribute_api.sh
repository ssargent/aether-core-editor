#!/bin/bash

# Test script for AttributeService API endpoints
# This script tests the basic API structure without requiring a database
# Run from project root: ./scripts/test_attribute_api.sh

# Change to project root directory
cd "$(dirname "$0")/.."

echo "Testing AttributeService API endpoints..."
echo

# Start the server in the background
echo "Starting server..."
go run ./cmd/server &
SERVER_PID=$!

# Wait for server to start
sleep 3

echo "Testing API endpoints..."

# Test gRPC reflection to see if AttributeService is registered
echo "1. Checking gRPC reflection for AttributeService:"
grpcurl -plaintext localhost:8080 list | grep -i attribute || echo "grpcurl not available or service not found"

echo

# Test creating an attribute (will fail without database but should show proper error)
echo "2. Testing CreateAttribute endpoint:"
curl -s -X POST http://localhost:8080/game.v1.AttributeService/CreateAttribute \
  -H "Content-Type: application/json" \
  -d '{"name":"Strength","description":"Physical power"}' | head -c 200

echo
echo

# Test listing attributes (will fail without database but should show proper error)
echo "3. Testing ListAttributes endpoint:"
curl -s -X POST http://localhost:8080/game.v1.AttributeService/ListAttributes \
  -H "Content-Type: application/json" \
  -d '{}' | head -c 200

echo
echo

# Clean up
echo "Stopping server..."
kill $SERVER_PID 2>/dev/null || true
wait $SERVER_PID 2>/dev/null || true

echo "Test completed!"
