## Service Implementation: CurrencyService

### Description
Implement the CurrencyService following the same patterns as WorldService to provide CRUD operations for currency entities.

### Requirements
- [ ] Implement service methods in `internal/services/currency.go`:
  - [ ] `CreateCurrency` - Create new currency
  - [ ] `GetCurrency` - Retrieve currency by ID
  - [ ] `UpdateCurrency` - Update currency details
  - [ ] `DeleteCurrency` - Delete currency (return deleted entity)
  - [ ] `ListCurrencies` - List with pagination and filtering
- [ ] Add proper error handling with Connect status codes
- [ ] Implement database queries for all CRUD operations
- [ ] Register service handler in `cmd/server/main.go`
- [ ] Add unit tests in `internal/services/currency_test.go`
- [ ] Update README.md to move service from "Planned" to "Implemented"

### Database Table
The service should interact with the `game.currency` table from the existing database schema.

### Acceptance Criteria
- All CRUD operations work correctly
- Proper error handling for not found, validation errors, etc.
- Pagination works for list operations
- Service is registered and accessible via gRPC/Connect
- Unit tests cover main functionality
- Code follows Go standards and existing patterns

**Labels:** enhancement, service, backend
**Title:** feat: implement CurrencyService
