# Authentication Backend Setup

This document describes the Go backend authentication system with database integration.

## Features Implemented

### Database Schema
- **Users table** with support for both email/password and OAuth users
- **Indexes** for performance optimization
- **UUID primary keys** for security
- **Soft deletes** with is_active flag

### Authentication Methods
1. **Email/Password Authentication**
   - Password hashing with bcrypt
   - User registration and login
   - JWT token generation

2. **Google OAuth Integration**
   - Goth library integration
   - Automatic user creation/update
   - Provider-specific user management

### API Endpoints

#### Authentication Routes
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `GET /api/auth/me` - Get current user (protected)
- `GET /api/auth/google` - Initiate Google OAuth
- `GET /api/auth/google/callback` - Google OAuth callback
- `POST /api/auth/logout` - User logout

#### Request/Response Examples

**Register User:**
```json
POST /api/auth/register
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Login:**
```json
POST /api/auth/login
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "user": {
    "id": "uuid",
    "name": "John Doe",
    "email": "john@example.com",
    "provider": "email",
    "is_active": true,
    "is_verified": false,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "token": "jwt-token"
}
```

## Environment Variables

Create a `.env` file in the server directory:

```env
# Server Configuration
PORT=8080
BASE_URL=http://localhost:8080

# Database Configuration
PG_DB_HOST=localhost
PG_DB_PORT=5432
PG_DB_USERNAME=postgres
PG_DB_PASSWORD=password
PG_DB_DATABASE=ngo_contribution
PG_DB_SCHEMA=public

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Google OAuth Configuration
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
```

## Database Setup

1. **Run migrations:**
   ```bash
   # The migration files are already created
   # Run them against your PostgreSQL database
   ```

2. **Database schema:**
   ```sql
   CREATE TABLE users (
       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
       name VARCHAR(255) NOT NULL,
       email VARCHAR(255) UNIQUE NOT NULL,
       password_hash VARCHAR(255),
       provider VARCHAR(50) DEFAULT 'email',
       provider_id VARCHAR(255),
       avatar_url TEXT,
       is_active BOOLEAN DEFAULT true,
       is_verified BOOLEAN DEFAULT false,
       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
   );
   ```

## Google OAuth Setup

1. **Create Google OAuth credentials:**
   - Go to [Google Cloud Console](https://console.cloud.google.com/)
   - Create a new project or select existing
   - Enable Google+ API
   - Create OAuth 2.0 credentials
   - Add authorized redirect URIs:
     - `http://localhost:8080/api/auth/google/callback` (development)
     - `https://yourdomain.com/api/auth/google/callback` (production)

2. **Set environment variables:**
   ```env
   GOOGLE_CLIENT_ID=your-client-id
   GOOGLE_CLIENT_SECRET=your-client-secret
   ```

## Running the Server

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Set up database:**
   - Ensure PostgreSQL is running
   - Create the database
   - Run the migration

3. **Start the server:**
   ```bash
   go run cmd/api/main.go
   ```

## Testing the API

### Using curl:

**Register:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

**Get current user:**
```bash
curl -X GET http://localhost:8080/api/auth/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Google OAuth:**
```bash
# Redirect to Google OAuth
curl -X GET http://localhost:8080/api/auth/google
```

## Security Features

- **Password hashing** with bcrypt
- **JWT tokens** with expiration
- **CORS configuration** for frontend integration
- **Input validation** on all endpoints
- **SQL injection protection** with parameterized queries
- **Soft deletes** for data integrity

## File Structure

```
server/
├── internal/
│   ├── config/
│   │   └── oauth.go           # OAuth configuration
│   ├── database/
│   │   └── database.go        # Database connection
│   ├── handlers/
│   │   └── auth_handler.go    # Authentication handlers
│   ├── middleware/
│   │   └── auth_middleware.go # JWT middleware
│   ├── models/
│   │   └── user.go           # User models
│   ├── repository/
│   │   └── user_repository.go # Database operations
│   ├── services/
│   │   ├── auth_service.go    # Authentication logic
│   │   └── jwt_service.go    # JWT token management
│   └── server/
│       ├── routes.go          # Route registration
│       └── server.go          # Server configuration
├── db/
│   └── migrations/
│       ├── 20250924181232_create_users_table.up.sql
│       └── 20250924181232_create_users_table.down.sql
└── cmd/
    └── api/
        └── main.go           # Application entry point
```

## Next Steps

1. **Frontend Integration**: Update frontend to use actual API endpoints
2. **Email Verification**: Add email verification for new users
3. **Password Reset**: Implement password reset functionality
4. **Rate Limiting**: Add rate limiting to prevent abuse
5. **Logging**: Add structured logging for better debugging
6. **Testing**: Add unit and integration tests
7. **Docker**: Containerize the application
8. **Production**: Configure for production deployment
