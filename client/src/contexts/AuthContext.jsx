import { createContext, useContext, useState, useEffect } from 'react';
import { apiRequest, API_ENDPOINTS } from '../config/api';

const AuthContext = createContext();

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  const fetchCurrentUser = async () => {
    const result = await apiRequest(API_ENDPOINTS.ME);
    if (result.success && result.data) {
      setUser(result.data);
      return { success: true, data: result.data };
    }
    return { success: false, error: result.error };
  };

  useEffect(() => {
    // Check for existing authentication
    const checkAuth = async () => {
      try {
        const token = localStorage.getItem('authToken');
        if (token) {
          // Verify token with backend
          const me = await fetchCurrentUser();
          if (!me.success) {
            // Token is invalid, remove it
            localStorage.removeItem('authToken');
          }
        }
      } catch (error) {
        console.error('Auth check failed:', error);
        localStorage.removeItem('authToken');
      } finally {
        setIsLoading(false);
      }
    };

    checkAuth();
  }, []);

  const login = async (email, password) => {
    try {
      const result = await apiRequest(API_ENDPOINTS.LOGIN, {
        method: 'POST',
        body: JSON.stringify({ email, password }),
      });

      if (result.success && result.data) {
        const { user: userData, token } = result.data;
        localStorage.setItem('authToken', token);
        setUser(userData);
        return { success: true };
      } else {
        return { success: false, error: result.error || 'Login failed. Please try again.' };
      }
    } catch (error) {
      // console.error('Login failed:', error);
      return { success: false, error: 'Login failed. Please try again.' };
    }
  };

  const signup = async (name, email, password) => {
    try {
      const result = await apiRequest(API_ENDPOINTS.REGISTER, {
        method: 'POST',
        body: JSON.stringify({ name, email, password }),
      });

      if (result.success && result.data) {
        const { user: userData, token } = result.data;
        localStorage.setItem('authToken', token);
        setUser(userData);
        return { success: true };
      } else {
        return { success: false, error: result.error || 'Signup failed. Please try again.' };
      }
    } catch (error) {
      console.error('Signup failed:', error);
      return { success: false, error: 'Signup failed. Please try again.' };
    }
  };

  const googleAuth = () => {
    // Redirect to backend Google OAuth endpoint
    window.location.href = API_ENDPOINTS.GOOGLE_AUTH;
  };

  const logout = async () => {
    try {
      // Call logout endpoint to invalidate token on server
      await apiRequest(API_ENDPOINTS.LOGOUT, {
        method: 'POST',
      });
    } catch (error) {
      console.error('Logout API call failed:', error);
    } finally {
      // Always clear local storage and user state
      localStorage.removeItem('authToken');
      setUser(null);
    }
  };

  const value = {
    user,
    isLoading,
    login,
    signup,
    googleAuth,
    logout,
    fetchCurrentUser,
    isAuthenticated: !!user
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};
