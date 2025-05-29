---
name: Implement ItemService
about: Add CRUD operations for item and inventory management
title: 'Implement ItemService'
labels: enhancement, backend, service-implementation
assignees: ''
---

## Description

Implement the `ItemService` to handle CRUD operations for items, equipment, and inventory management in the game world.

## Requirements

### Database Tables

- `items` - Base item definitions
- `equipment` - Equipment-specific properties
- `item_attributes` - Item stat modifiers
- `currencies` - Currency definitions

### Service Methods

Implement the following RPC methods:

- `CreateItem` - Create a new item
- `GetItem` - Retrieve an item by ID
- `UpdateItem` - Update item details
- `DeleteItem` - Delete an item
- `ListItems` - List items with pagination and filtering

### Features

- [ ] CRUD operations for items
- [ ] Handle equipment properties and relationships
- [ ] Manage item attributes and stat modifiers
- [ ] Proper error handling with Connect status codes
- [ ] Input validation for all fields
- [ ] Pagination support for listing
- [ ] Optional filtering by type, rarity, or category
- [ ] Handle currency relationships for item values

### Implementation Details

- Create service implementation in `internal/services/item.go`
- Register service with Connect handler in `cmd/server/main.go`
- Follow existing patterns from WorldService implementation
- Add comprehensive error handling
- Use proper SQL queries with parameterized statements

### Database Schema Reference

```sql
-- items table structure (from mud-schema-v2.sql)
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    item_type VARCHAR(50),
    -- Additional fields as defined in schema
);

-- Related tables
CREATE TABLE equipment (...);
CREATE TABLE item_attributes (...);
CREATE TABLE currencies (...);
```

### Testing

- [ ] Add unit tests for all service methods
- [ ] Test error cases and edge conditions
- [ ] Test pagination and filtering
- [ ] Test database constraints and relationships
- [ ] Test equipment and attribute relationships

### Acceptance Criteria

- All CRUD operations work correctly
- Proper error messages for invalid inputs
- Pagination works with configurable page sizes
- Filtering options function as expected
- Equipment and attribute relationships handled properly
- Service integrates properly with existing server setup
- Tests achieve good coverage of functionality

## Dependencies

- Requires database connection from `internal/database/db.go`
- Uses protobuf definitions from `protos/game/v1/item_service.proto`
- Follows patterns from existing WorldService implementation
