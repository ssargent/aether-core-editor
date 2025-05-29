## Service Implementation: AttributeService

### Description
Implement the AttributeService following the same patterns as WorldService to provide CRUD operations for character attribute entities.

### Requirements
- [ ] Implement service methods in `internal/services/attribute.go`:
  - [ ] `CreateAttribute` - Create new attribute
  - [ ] `GetAttribute` - Retrieve attribute by ID  
  - [ ] `UpdateAttribute` - Update attribute details
  - [ ] `DeleteAttribute` - Delete attribute (return deleted entity)
  - [ ] `ListAttributes` - List with pagination and filtering
- [ ] Add proper error handling with Connect status codes
- [ ] Implement database queries for all CRUD operations
- [ ] Register service handler in `cmd/server/main.go`
- [ ] Add unit tests in `internal/services/attribute_test.go`
- [ ] Update README.md to move service from "Planned" to "Implemented"

### Database Table
The service should interact with the `game.attributes` table from the existing database schema.

### Acceptance Criteria
- All CRUD operations work correctly
- Proper error handling for not found, validation errors, etc.
- Pagination works for list operations  
- Service is registered and accessible via gRPC/Connect
- Unit tests cover main functionality
- Code follows Go standards and existing patterns

**Labels:** enhancement, service, backend
**Title:** feat: implement AttributeService
