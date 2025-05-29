---
name: Service Implementation
about: Implement a new gRPC service for game entities
title: 'feat: implement [SERVICE_NAME]Service'
labels: ['enhancement', 'service', 'backend']
assignees: ''
---

## Service Implementation: [SERVICE_NAME]Service

### Description
Implement the [SERVICE_NAME]Service following the same patterns as WorldService to provide CRUD operations for [entity_name] entities.

### Requirements
- [ ] Implement service methods in `internal/services/[service_name].go`:
  - [ ] `Create[EntityName]` - Create new [entity_name]
  - [ ] `Get[EntityName]` - Retrieve [entity_name] by ID
  - [ ] `Update[EntityName]` - Update [entity_name] details
  - [ ] `Delete[EntityName]` - Delete [entity_name] (return deleted entity)
  - [ ] `List[EntityName]s` - List with pagination and filtering
- [ ] Add proper error handling with Connect status codes
- [ ] Implement database queries for all CRUD operations
- [ ] Register service handler in `cmd/server/main.go`
- [ ] Add unit tests in `internal/services/[service_name]_test.go`
- [ ] Update README.md to move service from "Planned" to "Implemented"

### Database Table
The service should interact with the `[table_name]` table from the existing database schema.

### Acceptance Criteria
- All CRUD operations work correctly
- Proper error handling for not found, validation errors, etc.
- Pagination works for list operations
- Service is registered and accessible via gRPC/Connect
- Unit tests cover main functionality
- Code follows Go standards and existing patterns

### Dependencies
- Depends on existing database schema and connection
- Should follow patterns established in WorldService implementation
