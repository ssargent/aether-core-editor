# Aether Core Editor

A modern web-based editor for creating immersive text-based gaming worlds. The Aether Core Editor provides a beautiful, intuitive interface for managing game world data through powerful gRPC protobuf services.

## Architecture

- **Backend**: Go with Connect-RPC framework
- **Database**: PostgreSQL with schema from `mud-schema-v2.sql`
- **API**: gRPC/Connect services with protobuf definitions
- **Frontend**: React 19 + TypeScript with Vite, TailwindCSS v4, and modern UI components

## Project Structure

```
├── client/              # React TypeScript frontend
│   ├── src/
│   │   ├── components/  # React components (LoginPage, etc.)
│   │   ├── App.tsx      # Main application component
│   │   ├── main.tsx     # React app entry point
│   │   └── index.css    # Global styles with Tailwind imports
│   ├── public/          # Static assets
│   ├── package.json     # Frontend dependencies and scripts
│   ├── vite.config.ts   # Vite build configuration
│   ├── tailwind.config.js  # TailwindCSS v4 configuration
│   ├── postcss.config.js   # PostCSS configuration
│   └── .nvmrc           # Node.js version specification (v22.16.0)
├── protos/              # Protobuf service definitions
│   └── aethercore/v1/   # Individual service proto files
├── server/              # Go backend server
│   ├── cmd/server/      # Main server application
│   ├── internal/
│   │   ├── database/    # Database connection and utilities
│   │   └── services/    # Business logic and RPC service implementations
│   └── go.mod           # Server dependencies
├── gen/                 # Generated protobuf code
│   ├── aethercore/v1/   # Protobuf message types
│   └── aethercorev1connect/  # Connect RPC service interfaces
├── scripts/             # Utility scripts
├── mud-schema-v2.sql    # PostgreSQL database schema
├── docker-compose.yml   # Development database
├── buf.gen.yaml         # Buf protobuf generation config
└── Makefile            # Build and development commands
```

## Quick Start

### Prerequisites

- **Node.js** 22.16.0+ (specified in `.nvmrc`)
- **Go** 1.22+
- **Docker and Docker Compose** (for database)
- **Protocol Buffers compiler** (`protoc`)

### Setup

1. **Clone and setup dependencies:**

   ```bash
   # Install protobuf generation tools at root level
   make tools
   
   # Setup server dependencies
   cd server && go mod download
   
   # Setup client dependencies  
   cd ../client && npm install
   ```

2. **Start the development database:**

   ```bash
   docker-compose up -d
   ```

3. **Generate protobuf code:**

   ```bash
   make generate
   ```

4. **Start the backend server:**

   ```bash
   cd server
   go run ./cmd/server
   ```

   The server will start on `http://localhost:8080`

5. **Start the frontend development server:**

   ```bash
   cd client
   npm run dev
   ```

   The client will start on `http://localhost:5173` (or next available port)

### Alternative Setup (Individual Components)

**Backend only:**

```bash
cd server
go build -o ../bin/server ./cmd/server
./bin/server
```

**Frontend only:**

```bash
cd client
npm run dev
```

## Client (Frontend)

The Aether Core Editor features a modern React-based frontend with a beautiful login interface and responsive design.

### Technology Stack

- **React 19** - Latest React with modern features
- **TypeScript** - Type-safe development
- **Vite** - Fast build tool and development server
- **TailwindCSS v4** - Utility-first CSS framework with custom design system
- **PostCSS** - CSS processing and optimization

### Features

- 🎨 **Beautiful Login Interface** - Modern glassmorphism design with gradient backgrounds
- 🌟 **Animated UI Elements** - Smooth transitions and floating background orbs
- 📱 **Responsive Design** - Works seamlessly on desktop and mobile devices
- 🎯 **Brand Identity** - Consistent Aether Core Editor branding and color scheme
- ♿ **Accessibility** - Proper ARIA labels, focus management, and semantic HTML
- 🔐 **Form Validation** - Client-side validation with TypeScript type safety

### Development Commands

```bash
cd client

# Start development server (with hot reload)
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview

# Run linting
npm run lint

# Type checking
npx tsc --noEmit
```

### Custom Design System

The client includes a custom TailwindCSS configuration with:

- **Primary Color Palette** - Blue gradients (primary-50 to primary-900)
- **Dark Color Palette** - Sophisticated grays (dark-50 to dark-900)
- **Custom Animations** - fade-in, slide-up, pulse-slow
- **Inter Font Family** - Modern, readable typography
- **Glassmorphism Effects** - Semi-transparent containers with backdrop blur

### Environment Setup

The client requires Node.js 22.16.0+ (specified in `.nvmrc`). If using nvm:

```bash
cd client
nvm use  # Uses version from .nvmrc
npm install
```

### Available Routes

Currently implemented:
- `/` - Login page with authentication form

Planned:
- `/dashboard` - Main editor interface
- `/worlds` - World management
- `/settings` - User preferences

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

### AttributeService
- `CreateAttribute` - Create a new character attribute
- `GetAttribute` - Retrieve an attribute by ID
- `UpdateAttribute` - Update attribute details
- `DeleteAttribute` - Delete an attribute
- `ListAttributes` - List attributes with pagination and filtering

### Planned Services
- CapabilityService
- CharacterClassService
- CurrencyService
- EnemyService
- ItemService
- NPCService
- RaceService
- WorldNodeService

## Development

### Available Commands

**Root level (protobuf generation):**

```bash
make generate    # Generate protobuf code
make tools       # Install protobuf tools
make clean       # Clean generated files
```

**Server level (from /server directory):**

```bash
cd server
go build -o ../bin/server ./cmd/server  # Build the server
go run ./cmd/server                     # Run the server
go test ./...                           # Run tests
go mod download                         # Download dependencies
```

**Client level (from /client directory):**

```bash
cd client
npm run dev      # Start development server
npm run build    # Build for production
npm run preview  # Preview production build
npm run lint     # Run ESLint
npm install      # Install dependencies
npm update       # Update dependencies
```

### Development Workflow

1. **Start the database:**
   ```bash
   docker-compose up -d
   ```

2. **Start the backend (Terminal 1):**
   ```bash
   cd server && go run ./cmd/server
   ```

3. **Start the frontend (Terminal 2):**
   ```bash
   cd client && npm run dev
   ```

4. **Access the application:**
   - Frontend: `http://localhost:5173` (or next available port)
   - Backend API: `http://localhost:8080`
   - Database: `localhost:5432`

### Hot Reload Development

Both the client and server support hot reload:

- **Client**: Vite provides instant hot module replacement (HMR)
- **Server**: Use `air` or similar tools for Go hot reload:
  ```bash
  # Install air (Go hot reload tool)
  go install github.com/cosmtrek/air@latest
  
  # Run server with hot reload
  cd server && air
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

### Immediate Priorities

1. **Authentication Integration** - Connect login form to backend authentication service
2. **Dashboard Implementation** - Create main editor interface after successful login
3. **API Integration** - Connect frontend to backend gRPC services using connect-web
4. **World Management UI** - Build interfaces for creating and managing game worlds

### Medium Term Goals

1. **Complete Service Implementation** - Add remaining backend services for all game entities
2. **Advanced UI Components** - Rich text editors, drag-and-drop interfaces, data tables
3. **Real-time Features** - WebSocket connections for collaborative editing
4. **Data Visualization** - Charts and graphs for game analytics

### Long Term Vision

1. **Multi-user Support** - User management, permissions, and team collaboration
2. **Plugin System** - Extensible architecture for custom game mechanics
3. **Export/Import Tools** - Integration with popular game engines and formats
4. **Cloud Deployment** - Production-ready infrastructure and CI/CD pipelines

## Contributing

1. Follow Go coding standards and conventions for backend development
2. Use TypeScript and React best practices for frontend development
3. Generate protobuf code after making changes to `.proto` files
4. Maintain the established design system and component patterns
5. Add tests for new functionality (both frontend and backend)
6. Update documentation as needed

### Code Style

- **Backend**: Standard Go formatting with `gofmt` and `golint`
- **Frontend**: ESLint configuration with React and TypeScript rules
- **Commits**: Follow conventional commit message format

### Pull Request Process

1. Create feature branch from `main`
2. Implement changes with appropriate tests
3. Update documentation if needed
4. Ensure all tests pass and linting is clean
5. Submit PR with clear description of changes

---

*Built with ❤️ for game developers and world builders*
