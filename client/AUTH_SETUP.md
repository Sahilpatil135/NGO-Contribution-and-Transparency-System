# Authentication Setup

This document describes the authentication system implemented for the NGO Contribution and Transparency System.

## Features

### Frontend Components
- **Login Page** (`/login`): Email/password login with Google OAuth option
- **Signup Page** (`/signup`): User registration with email/password and Google OAuth
- **Dashboard** (`/dashboard`): Protected area for authenticated users
- **AuthContext**: Centralized authentication state management

### Authentication Methods
1. **Email/Password Authentication**
   - Form validation (password confirmation, minimum length)
   - Mock API integration ready for backend connection
   - Error handling and loading states

2. **Google OAuth Integration**
   - Google OAuth button with proper styling
   - Redirects to `/api/auth/google` for backend OAuth flow
   - Ready for Goth integration

## Backend Integration Points

### API Endpoints (to be implemented)
- `POST /api/auth/login` - Email/password login
- `POST /api/auth/signup` - User registration
- `GET /api/auth/google` - Google OAuth initiation
- `GET /api/auth/google/callback` - Google OAuth callback
- `POST /api/auth/logout` - User logout
- `GET /api/auth/me` - Get current user info

### Goth Integration
The frontend is designed to work with Goth for Google OAuth:

1. **OAuth Flow**: Users click "Continue with Google" → redirects to `/api/auth/google`
2. **Backend Setup**: Configure Goth with Google provider
3. **Session Management**: Use JWT tokens or sessions for authentication state
4. **User Data**: Store user information from OAuth providers

## File Structure

```
src/
├── components/
│   ├── auth/
│   │   ├── Login.jsx          # Login component
│   │   ├── Signup.jsx         # Signup component
│   │   └── Auth.css           # Authentication styles
│   └── Dashboard.jsx          # Protected dashboard
├── contexts/
│   └── AuthContext.jsx        # Authentication context
└── App.jsx                    # Main app with routing
```

## Usage

### Development
```bash
npm run dev
```

### Authentication Flow
1. User visits the app → redirected to `/login` if not authenticated
2. User can sign up with email/password or use Google OAuth
3. After successful authentication → redirected to `/dashboard`
4. Logout clears authentication state and redirects to login

### Environment Variables (for backend)
```env
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
JWT_SECRET=your_jwt_secret
```

## Next Steps

1. **Backend Implementation**: Implement the API endpoints in Go
2. **Goth Setup**: Configure Google OAuth with Goth
3. **Database Integration**: Store user data and sessions
4. **Security**: Add CSRF protection, rate limiting, etc.
5. **Testing**: Add unit and integration tests

## Styling

The authentication pages feature:
- Modern gradient backgrounds
- Responsive design for mobile/desktop
- Dark mode support
- Loading states and error handling
- Google OAuth button with official styling
