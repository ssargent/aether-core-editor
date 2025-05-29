---
name: Implement RaceService
about: Add CRUD operations for character race management
title: 'Implement RaceService'
labels: enhancement, backend, service-implementation
assignees: ''
---

## Description

Implement the `RaceService` to handle CRUD operations for character races and their associated features in the game world.

## Requirements

### Database Tables

- `races` - Character race definitions
- `race_features` - Race-specific features and abilities
- `race_attribute_modifiers` - Attribute bonuses/penalties by race

### Service Methods

Implement the following RPC methods:

- `CreateRace` - Create a new character race
- `GetRace` - Retrieve a race by ID
- `UpdateRace` - Update race details
- `DeleteRace` - Delete a race
- `ListRaces` - List races with pagination and filtering

### Features

- [ ] CRUD operations for races
- [ ] Handle race features and abilities
- [ ] Manage attribute modifiers by race
- [ ] Proper error handling with Connect status codes
- [ ] Input validation for all fields
- [ ] Pagination support for listing
- [ ] Optional filtering by name or category
- [ ] Handle racial bonuses and penalties

### Implementation Details

- Create service implementation in `internal/services/race.go`
- Register service with Connect handler in `cmd/server/main.go`
- Follow existing patterns from WorldService implementation
- Add comprehensive error handling
- Use proper SQL queries with parameterized statements

### Database Schema Reference

```sql
-- races table structure (from mud-schema-v2.sql)
CREATE TABLE races (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    size_category VARCHAR(20),
    -- Additional fields as defined in schema
);

-- Related tables
CREATE TABLE race_features (...);
CREATE TABLE race_attribute_modifiers (...);
```

### Testing

- [ ] Add unit tests for all service methods
- [ ] Test error cases and edge conditions
- [ ] Test pagination and filtering
- [ ] Test database constraints and relationships
- [ ] Test feature and attribute modifier relationships

### Acceptance Criteria

- All CRUD operations work correctly
- Proper error messages for invalid inputs
- Pagination works with configurable page sizes
- Filtering options function as expected
- Features and attribute modifiers handled properly
- Service integrates properly with existing server setup
- Tests achieve good coverage of functionality

## Dependencies

- Requires database connection from `internal/database/db.go`
- Uses protobuf definitions from `protos/game/v1/race_service.proto`
- Follows patterns from existing WorldService implementation
