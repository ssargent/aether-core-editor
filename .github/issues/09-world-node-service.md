---
name: Implement WorldNodeService
about: Add CRUD operations for world node and location management
title: 'Implement WorldNodeService'
labels: enhancement, backend, service-implementation
assignees: ''
---

## Description

Implement the `WorldNodeService` to handle CRUD operations for world nodes (locations/rooms) and their connections within game worlds.

## Requirements

### Database Tables

- `world_nodes` - Individual locations/rooms in worlds
- `node_connections` - Directional connections between nodes
- `node_exits` - Exit definitions and descriptions

### Service Methods

Implement the following RPC methods:

- `CreateWorldNode` - Create a new world node/location
- `GetWorldNode` - Retrieve a world node by ID
- `UpdateWorldNode` - Update world node details
- `DeleteWorldNode` - Delete a world node
- `ListWorldNodes` - List world nodes with pagination and filtering
- `ConnectNodes` - Create connections between nodes
- `DisconnectNodes` - Remove connections between nodes
- `GetNodeConnections` - Get all connections for a specific node

### Features

- [ ] CRUD operations for world nodes
- [ ] Handle node connections and exits
- [ ] Manage directional relationships between locations
- [ ] Support for different exit types (doors, passages, etc.)
- [ ] Proper error handling with Connect status codes
- [ ] Input validation for all fields
- [ ] Pagination support for listing
- [ ] Optional filtering by world, name, or type
- [ ] Prevent circular references and invalid connections

### Implementation Details

- Create service implementation in `internal/services/world_node.go`
- Register service with Connect handler in `cmd/server/main.go`
- Follow existing patterns from WorldService implementation
- Add comprehensive error handling
- Use proper SQL queries with parameterized statements
- Implement graph validation for node connections

### Database Schema Reference

```sql
-- world_nodes table structure (from mud-schema-v2.sql)
CREATE TABLE world_nodes (
    id SERIAL PRIMARY KEY,
    world_id INTEGER REFERENCES worlds(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    -- Additional fields as defined in schema
);

-- Related tables
CREATE TABLE node_connections (...);
CREATE TABLE node_exits (...);
```

### Testing

- [ ] Add unit tests for all service methods
- [ ] Test error cases and edge conditions
- [ ] Test pagination and filtering
- [ ] Test database constraints and relationships
- [ ] Test connection validation and graph integrity
- [ ] Test circular reference prevention

### Acceptance Criteria

- All CRUD operations work correctly
- Proper error messages for invalid inputs
- Pagination works with configurable page sizes
- Filtering options function as expected
- Node connections handled properly with validation
- Prevents invalid or circular connections
- Service integrates properly with existing server setup
- Tests achieve good coverage of functionality

## Dependencies

- Requires database connection from `internal/database/db.go`
- Uses protobuf definitions from `game_v1.proto`
- Follows patterns from existing WorldService implementation
- Related to WorldService for world associations
