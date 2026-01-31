# Snippetbox

A full-stack web application built with Go that allows users to create, share, and manage text snippets with time-based expiration. Think Pastebin or GitHub Gists, but built from scratch to learn Go web development.

## Overview

Snippetbox is a production-ready web application following best practices from Alex Edwards' "Let's Go" book. It demonstrates modern Go web development with minimal dependencies, focusing on the standard library and idiomatic Go patterns.

## Features

### Core Functionality
- **Create Snippets** - Paste and save text snippets
- **View Snippets** - Browse and read shared snippets
- **Time-Based Expiration** - Snippets automatically expire (1 day, 7 days, or 365 days)
- **User Authentication** - Secure registration and login system
- **Session Management** - Stateful user sessions
- **Protected Routes** - Only authenticated users can create snippets

### Technical Features
- Server-side rendering with HTML templates
- MySQL database integration
- Middleware architecture (logging, recovery, authentication)
- HTTPS/TLS support with HTTP/2
- CSRF protection
- Secure password hashing with bcrypt
- Custom error handling and logging
- RESTful routing patterns

## Technologies Used

- **Go** - Core programming language
- **MySQL** - Database for data persistence
- **html/template** - Server-side templating
- **crypto/bcrypt** - Password hashing
- **net/http** - HTTP server (standard library)
- **TLS/HTTPS** - Secure connections
- **Middleware** - Custom middleware chain

## Getting Started

### Prerequisites

- Go 1.16 or higher
- MySQL 5.7+ or MariaDB
- Optional: TLS certificates for HTTPS

### Installation

```bash
# Clone the repository
git clone https://github.com/BlackMunda/snippetbox.git
cd snippetbox

# Install dependencies
go mod download

# Set up the database
mysql -u root -p < setup.sql

# Generate TLS certificates (for development)
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

# Run the application
go run ./cmd/web
```

The application will start on `https://localhost:4000`

## Database Setup

```sql
-- Create database
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create snippets table
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

-- Create users table
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

-- Create sessions table
CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```

## Project Structure

```
snippetbox/
├── cmd/
│   └── web/
│       ├── main.go              # Application entry point
│       ├── handlers.go          # HTTP handlers
│       ├── helpers.go           # Helper functions
│       ├── middleware.go        # Middleware chain
│       └── routes.go            # Route definitions
├── pkg/
│   ├── models/
│   │   ├── snippets.go          # Snippet database model
│   │   └── users.go             # User database model
│   └── forms/
│       └── forms.go             # Form validation
├── ui/
│   ├── html/
│   │   ├── base.layout.tmpl     # Base template
│   │   ├── home.page.tmpl       # Home page
│   │   ├── show.page.tmpl       # Show snippet
│   │   ├── create.page.tmpl     # Create snippet
│   │   └── ...
│   └── static/
│       ├── css/
│       │   └── main.css
│       └── js/
│           └── main.js
├── tls/
│   ├── cert.pem                 # TLS certificate
│   └── key.pem                  # TLS private key
├── go.mod
└── go.sum
```

## Configuration

The application accepts the following command-line flags:

```bash
go run ./cmd/web -addr=":4000" -dsn="user:pass@/snippetbox?parseTime=true"
```

**Flags:**
- `-addr` - HTTP network address (default: ":4000")
- `-dsn` - MySQL data source name
- `-static-dir` - Path to static assets (default: "./ui/static")

## Routes

```
GET  /                      - Home page (list recent snippets)
GET  /snippet/:id           - View a specific snippet
GET  /snippet/create        - Show create snippet form (auth required)
POST /snippet/create        - Create a new snippet (auth required)

GET  /user/signup           - Show signup form
POST /user/signup           - Create new user account
GET  /user/login            - Show login form
POST /user/login            - Authenticate user
POST /user/logout           - Logout user (auth required)

GET  /static/*              - Serve static files
```

## Middleware Chain

1. **secureHeaders** - Set security headers (X-Frame-Options, X-XSS-Protection)
2. **logRequest** - Log all HTTP requests
3. **recoverPanic** - Recover from panics gracefully
4. **requireAuthentication** - Protect authenticated routes

## Security Features

- **HTTPS/TLS** - All traffic encrypted
- **HTTP/2** - Modern protocol support
- **CSRF Protection** - Token-based CSRF prevention
- **Secure Headers** - X-Frame-Options, X-XSS-Protection, X-Content-Type-Options
- **Password Hashing** - bcrypt with cost factor 12
- **Session Security** - HTTP-only, secure cookies
- **SQL Injection Prevention** - Prepared statements
- **Input Validation** - Form validation and sanitization

## Example Usage

### Creating a Snippet

1. Navigate to `https://localhost:4000/snippet/create`
2. Log in or sign up
3. Enter snippet title and content
4. Select expiration time (1 day, 7 days, or 365 days)
5. Submit

### Viewing Snippets

- Visit home page to see recent snippets
- Click on any snippet to view full content
- Expired snippets are automatically hidden

## Error Handling

The application includes comprehensive error handling:
- **404 Not Found** - Custom error page
- **500 Internal Server Error** - Custom error page with logging
- **400 Bad Request** - Invalid form submissions
- **403 Forbidden** - Unauthorized access attempts

## Logging

Structured logging with different levels:
- **INFO** - General application events
- **ERROR** - Error conditions with stack traces
- **DEBUG** - Detailed debugging information (development only)

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...
```

## Learning Outcomes

This project taught me:
- Idiomatic Go web development
- HTTP server implementation
- Middleware patterns and composition
- Template rendering and layouts
- Database integration with MySQL
- Session management
- Authentication and authorization
- HTTPS/TLS configuration
- Error handling strategies
- Structured logging
- Testing HTTP handlers
- Dependency injection patterns

## Key Go Concepts Demonstrated

- **Interfaces** - For database abstraction
- **Goroutines** - Implicit in HTTP server
- **Context** - Request-scoped values
- **Middleware** - Function composition
- **Error Handling** - Explicit error checking
- **Struct Embedding** - Template data composition
- **Method Sets** - HTTP handler methods

## Performance Considerations

- Connection pooling for database
- Template caching
- Static file serving optimization
- Graceful shutdown handling
- Request timeout configuration

## Future Enhancements

- [ ] Syntax highlighting for code snippets
- [ ] Snippet search functionality
- [ ] User profiles and snippet history
- [ ] Snippet editing (within time limit)
- [ ] API endpoints for programmatic access
- [ ] Rate limiting
- [ ] Email verification
- [ ] Password reset functionality
- [ ] Snippet categories/tags
- [ ] Private/public snippet visibility

## Production Deployment

For production deployment:

1. Use environment variables for sensitive config
2. Enable HTTPS with proper certificates (Let's Encrypt)
3. Set up proper logging and monitoring
4. Configure database connection pooling
5. Use a reverse proxy (nginx)
6. Enable rate limiting
7. Set up automated backups
8. Use systemd or similar for process management

## Contributing

This is a learning project, but feedback and suggestions are welcome!

## Acknowledgments

Built following Alex Edwards' excellent "Let's Go" book. Highly recommended for learning Go web development!

## Contact

**Devashish Singh**
- GitHub: [@BlackMunda](https://github.com/BlackMunda)
- Email: devashishsingh488@gmail.com

---

*Built with 🐹 Go and a passion for clean code*
