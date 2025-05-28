# Aether Core Editor

A web-based editor for a text-based game engine with gRPC protobuf services to manage game world data.

## Architecture

- **Backend**: Go with Connect-RPC framework
- **Database**: PostgreSQL with schema from `mud-schema-v2.sql`
- **API**: gRPC/Connect services with protobuf definitions
- **Frontend**: (To be implemented) React TypeScript with connect-web

## Project Structure

```
├── cmd/server/          # Main server application
├── internal/
│   ├── database/        # Database connection and utilities
│   └── services/        # Business logic and RPC service implementations
├── gen/                 # Generated protobuf code
│   ├── game/v1/         # Protobuf message types
│   └── gamev1connect/   # Connect RPC service interfaces
├── game_v1.proto        # Protobuf service definitions
├── mud-schema-v2.sql    # PostgreSQL database schema
├── docker-compose.yml   # Development database
└── Makefile            # Build and development commands
```

## Quick Start

### Prerequisites

- Go 1.22+
- Docker and Docker Compose (for database)
- Protocol Buffers compiler (`protoc`)

### Setup

1. **Clone and setup dependencies:**
   ```bash
   go mod download
   make tools  # Install protobuf generation tools
   ```

2. **Start the development database:**
   ```bash
   docker-compose up -d
   ```

3. **Generate protobuf code:**
   ```bash
   make generate
   ```

4. **Build and run the server:**
   ```bash
   make build
   make run
   ```

   Or directly:
   ```bash
   go run ./cmd/server
   ```

## Configuration

Set environment variables or create a `.env` file (see `.env.example`):

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=aether
DB_SSLMODE=disable

# Server
PORT=8080
```

## API Services

The server implements the following gRPC/Connect services:

### WorldService
- `CreateWorld` - Create a new game world
- `GetWorld` - Retrieve a world by ID
- `UpdateWorld` - Update world details
- `DeleteWorld` - Delete a world
- `ListWorlds` - List worlds with pagination and filtering

### Planned Services
- AttributeService
- CapabilityService
- CharacterClassService
- CurrencyService
- EnemyService
- ItemService
- NPCService
- RaceService
- WorldNodeService

## Development

### Available Make Commands

```bash
make generate    # Generate protobuf code
make build      # Build the server
make run        # Run the server
make test       # Run tests
make clean      # Clean generated files
make deps       # Download dependencies
make tools      # Install protobuf tools
```

### Database Schema

The database schema is automatically loaded when starting the Docker Compose PostgreSQL instance. It includes tables for:

- Worlds and world nodes
- Character attributes, capabilities, and classes
- Items, currencies, and equipment
- NPCs and enemies
- Races and character features

### Testing the API

The server supports both gRPC and gRPC-Web protocols. You can test it using:

1. **grpcurl** (for gRPC):
   ```bash
   grpcurl -plaintext localhost:8080 list
   ```

2. **curl** (for Connect protocol):
   ```bash
   curl -X POST http://localhost:8080/game.v1.WorldService/CreateWorld \
     -H "Content-Type: application/json" \
     -d '{"code":"test","name":"Test World","description":"A test world"}'
   ```

3. **gRPC reflection** is enabled for development at:
   - `http://localhost:8080/grpc.reflection.v1.ServerReflection`
   - `http://localhost:8080/grpc.reflection.v1alpha.ServerReflection`

## Next Steps

1. **Complete remaining service implementations** - Add services for all entities in the database schema
2. **Add comprehensive tests** - Unit and integration tests for all services  
3. **Implement React frontend** - TypeScript application using connect-web
4. **Add authentication and authorization** - User management and permissions
5. **Add API documentation** - OpenAPI/Swagger docs generation
6. **Production deployment** - Docker images, Kubernetes manifests, CI/CD

## Contributing

1. Follow Go coding standards and conventions
2. Generate protobuf code after making changes to `.proto` files
3. Add tests for new functionality
4. Update documentation as needed
An editor for the AetherCore Engine
