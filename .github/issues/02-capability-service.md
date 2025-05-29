## Service Implementation: CapabilityService

### Description
Implement the CapabilityService following the same patterns as WorldService to provide CRUD operations for character capability entities.

### Requirements
- [ ] Implement service methods in `internal/services/capability.go`:
  - [ ] `CreateCapability` - Create new capability
  - [ ] `GetCapability` - Retrieve capability by ID
  - [ ] `UpdateCapability` - Update capability details
  - [ ] `DeleteCapability` - Delete capability (return deleted entity)
  - [ ] `ListCapabilities` - List with pagination and filtering
- [ ] Add proper error handling with Connect status codes
- [ ] Implement database queries for all CRUD operations
- [ ] Register service handler in `cmd/server/main.go`
- [ ] Add unit tests in `internal/services/capability_test.go`
- [ ] Update README.md to move service from "Planned" to "Implemented"

### Database Table
The service should interact with the `game.capabilities` table from the existing database schema.

### Acceptance Criteria
- All CRUD operations work correctly
- Proper error handling for not found, validation errors, etc.
- Pagination works for list operations
- Service is registered and accessible via gRPC/Connect
- Unit tests cover main functionality
- Code follows Go standards and existing patterns

**Labels:** enhancement, service, backend
**Title:** feat: implement CapabilityService
