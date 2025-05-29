---
name: Implement NPCService
about: Add CRUD operations for NPC management
title: 'Implement NPCService'
labels: enhancement, backend, service-implementation
assignees: ''
---

## Description

Implement the `NPCService` to handle CRUD operations for non-player characters (NPCs) in the game world.

## Requirements

### Database Tables

- `npcs` - Main NPC definitions
- `npc_inventory` - NPC inventory items
- `npc_locations` - NPC locations and spawn points

### Service Methods

Implement the following RPC methods:

- `CreateNPC` - Create a new NPC
- `GetNPC` - Retrieve an NPC by ID
- `UpdateNPC` - Update NPC details
- `DeleteNPC` - Delete an NPC
- `ListNPCs` - List NPCs with pagination and filtering

### Features

- [ ] CRUD operations for NPCs
- [ ] Handle NPC inventory relationships
- [ ] Manage NPC location assignments
- [ ] Proper error handling with Connect status codes
- [ ] Input validation for all fields
- [ ] Pagination support for listing
- [ ] Optional filtering by name, type, or location
- [ ] Handle dialogue and interaction properties

### Implementation Details

- Create service implementation in `internal/services/npc.go`
- Register service with Connect handler in `cmd/server/main.go`
- Follow existing patterns from WorldService implementation
- Add comprehensive error handling
- Use proper SQL queries with parameterized statements

### Database Schema Reference

```sql
-- npcs table structure (from mud-schema-v2.sql)
CREATE TABLE npcs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    npc_type VARCHAR(50),
    -- Additional fields as defined in schema
);

-- Related tables
CREATE TABLE npc_inventory (...);
CREATE TABLE npc_locations (...);
```

### Testing

- [ ] Add unit tests for all service methods
- [ ] Test error cases and edge conditions
- [ ] Test pagination and filtering
- [ ] Test database constraints and relationships
- [ ] Test inventory and location relationships

### Acceptance Criteria

- All CRUD operations work correctly
- Proper error messages for invalid inputs
- Pagination works with configurable page sizes
- Filtering options function as expected
- Inventory and location relationships handled properly
- Service integrates properly with existing server setup
- Tests achieve good coverage of functionality

## Dependencies

- Requires database connection from `internal/database/db.go`
- Uses protobuf definitions from `game_v1.proto`
- Follows patterns from existing WorldService implementation
