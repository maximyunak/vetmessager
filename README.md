# Vetmessager

Backend for [Vetmessager](https://github.com/vetalana28/svetmessager).

## Installation

### Clone repository

```bash
git clone https://github.com/maximyunak/vetmessager.git
cd vetmessager
```

### Start PostgreSQL

```bash
make env-up
```

### Run database migrations

```bash
make migrate-up
```

### Start application in development mode

```bash
make app-run
```

The application will be available at:

```text
http://localhost:5050
```

### Build and deploy with Docker

```bash
make app-deploy
```

## API

### Base URL

```text
http://localhost:5050/api/{version}/{route}
```

For example:

```text
http://localhost:5050/api/v1/users
```

### Swagger

Swagger API documentation is available at:

```text
http://localhost:5050/swagger
```


## Project Structure

```text
.
├── cmd/                         # Application entry points
│   └── vetmessager/
│       ├── main.go              # Application entry point
│       └── Dockerfile           # Docker image configuration
│
├── docs/                        # Swagger API documentation
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── internal/                    # Application internal code
│   ├── core/                    # Shared application components
│   │   ├── auth/                # JWT authentication and token management
│   │   ├── domain/              # Domain models
│   │   ├── errors/              # Common application errors
│   │   ├── logger/              # Logging configuration and implementation
│   │   ├── password/            # Password hashing and verification
│   │   ├── repository/           # Shared repository infrastructure
│   │   │   └── postgres/
│   │   │       └── pull/         # PostgreSQL connection and configuration
│   │   └── transport/            # Shared HTTP infrastructure
│   │       └── http/
│   │           ├── middleware/   # HTTP middleware
│   │           ├── request/      # HTTP request decoding
│   │           ├── response/     # HTTP response handling
│   │           ├── server/       # HTTP server and routing
│   │           └── utils/        # HTTP utility functions
│   │
│   └── features/                 # Application features
│       └── users/                # User-related functionality
│           ├── repository/       # User data access
│           │   └── postgres/     # PostgreSQL implementation
│           ├── service/          # User business logic
│           └── transport/        # User HTTP API
│               └── http/
│
├── migrations/                   # Database migrations
│   ├── 000001_init.up.sql
│   └── 000001_init.down.sql
│
├── Makefile                      # Development and deployment commands
├── docker-compose.yml            # Docker infrastructure configuration
├── go.mod                        # Go module and dependencies
└── go.sum                        # Dependency checksums