# Frontend-Backend Integration Guide

This document describes how the React frontend integrates with the Go backend authentication system.

## 🔗 **Integration Overview**

The frontend and backend are now fully integrated with:
- **Real API calls** to Go backend endpoints
- **JWT token management** for authentication state
- **Google OAuth** flow with proper redirects
- **Error handling** for API responses
- **Loading states** for better UX

## 📁 **Updated Frontend Structure**

```
client/src/
├── config/
│   ├── api.js              # API configuration and endpoints
│   └── environment.js      # Environment configuration
├── contexts/
│   └── AuthContext.jsx     # Updated with real API calls
├── hooks/
│   └── useOAuthCallback.js # OAuth callback handling
├── pages/
│   └── OAuthCallback.jsx   # OAuth callback page
└── components/
    └── auth/
        ├── Login.jsx       # Updated with real API integration
        └── Signup.jsx       # Updated with real API integration
```

## 🔧 **API Integration**

### **Authentication Flow**
1. **User Registration/Login**: Frontend calls backend API
2. **JWT Token**: Backend returns JWT token
3. **Token Storage**: Frontend stores token in localStorage
4. **API Requests**: Token sent in Authorization header
5. **Token Validation**: Backend validates token on protected routes

### **API Endpoints Used**
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `GET /api/auth/me` - Get current user (protected)
- `GET /api/auth/google` - Google OAuth initiation
- `POST /api/auth/logout` - User logout

## 🚀 **Setup Instructions**

### **1. Backend Setup**
```bash
cd server
# Set up environment variables (see AUTH_BACKEND_SETUP.md)
# Start the Go server
go run cmd/api/main.go
```

### **2. Frontend Setup**
```bash
cd client
# Install dependencies (already done)
npm install

# Set environment variables (optional)
# Create .env file with:
# VITE_API_BASE_URL=http://localhost:8080

# Start the frontend
npm run dev
```

### **3. Database Setup**
```bash
# Ensure PostgreSQL is running
# Create the database
# Run the migration (see AUTH_BACKEND_SETUP.md)
```

## 🔐 **Authentication Features**

### **Email/Password Authentication**
- ✅ User registration with validation
- ✅ User login with credentials
- ✅ Password hashing (bcrypt)
- ✅ Input validation and error handling
- ✅ JWT token management

### **Google OAuth Authentication**
- ✅ Google OAuth initiation
- ✅ OAuth callback handling
- ✅ Automatic user creation/update
- ✅ Seamless redirect flow

### **Security Features**
- ✅ JWT token expiration (24 hours)
- ✅ Token validation on protected routes
- ✅ CORS configuration for frontend
- ✅ Input sanitization and validation
- ✅ Secure password handling

## 🧪 **Testing the Integration**

### **1. Test Email/Password Authentication**
```bash
# Start both servers
# Backend: http://localhost:8080
# Frontend: http://localhost:5173

# Navigate to http://localhost:5173
# Try registering a new user
# Try logging in with credentials
```

### **2. Test Google OAuth**
```bash
# Ensure Google OAuth is configured in backend
# Click "Continue with Google" button
# Complete OAuth flow
# Verify user is created and logged in
```

### **3. Test Protected Routes**
```bash
# Login successfully
# Navigate to /dashboard
# Verify user data is displayed
# Test logout functionality
```

## 🔧 **Configuration**

### **Environment Variables**

**Backend (.env):**
```env
PORT=8080
BASE_URL=http://localhost:8080
PG_DB_HOST=localhost
PG_DB_PORT=5432
PG_DB_USERNAME=postgres
PG_DB_PASSWORD=password
PG_DB_DATABASE=ngo_contribution
PG_DB_SCHEMA=public
JWT_SECRET=your-super-secret-jwt-key
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
```

**Frontend (.env):**
```env
VITE_API_BASE_URL=http://localhost:8080
```

### **CORS Configuration**
The backend is configured to allow requests from:
- `http://localhost:5173` (Vite dev server)
- `https://*` (Production HTTPS)
- `http://*` (Development HTTP)

## 🐛 **Troubleshooting**

### **Common Issues**

1. **CORS Errors**
   - Ensure backend CORS is configured correctly
   - Check that frontend URL is allowed in backend CORS settings

2. **API Connection Errors**
   - Verify backend is running on correct port
   - Check VITE_API_BASE_URL in frontend
   - Ensure database is connected

3. **Authentication Issues**
   - Check JWT_SECRET is set in backend
   - Verify token is being sent in requests
   - Check browser localStorage for token

4. **Google OAuth Issues**
   - Verify Google OAuth credentials are set
   - Check redirect URI configuration
   - Ensure BASE_URL matches OAuth settings

### **Debug Steps**

1. **Check Network Tab**
   - Open browser DevTools
   - Monitor API requests and responses
   - Check for 401/403 errors

2. **Check Console Logs**
   - Look for JavaScript errors
   - Check API response errors
   - Verify token presence

3. **Check Backend Logs**
   - Monitor server logs for errors
   - Check database connection
   - Verify OAuth configuration

## 📈 **Next Steps**

1. **Production Deployment**
   - Configure production environment variables
   - Set up HTTPS certificates
   - Configure production database

2. **Additional Features**
   - Email verification
   - Password reset functionality
   - User profile management
   - Role-based access control

3. **Monitoring & Logging**
   - Add structured logging
   - Set up error monitoring
   - Add performance metrics

## 🔒 **Security Considerations**

- **JWT Secret**: Use a strong, random secret in production
- **HTTPS**: Always use HTTPS in production
- **CORS**: Restrict CORS to specific domains in production
- **Database**: Use connection pooling and proper indexing
- **Rate Limiting**: Consider adding rate limiting for auth endpoints
- **Input Validation**: All inputs are validated on both frontend and backend

The authentication system is now fully integrated and ready for development and testing!
