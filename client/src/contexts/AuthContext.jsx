import { createContext, useContext, useState, useEffect } from 'react';

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

  useEffect(() => {
    // Check for existing authentication
    const checkAuth = async () => {
      try {
        const token = localStorage.getItem('authToken');
        if (token) {
          // TODO: Verify token with backend
          // For now, just set a mock user
          setUser({
            id: '1',
            name: 'John Doe',
            email: 'john@example.com',
            provider: 'email' // or 'google'
          });
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
      // TODO: Implement actual login API call
      console.log('Login attempt:', { email, password });
      
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Mock successful login
      const mockUser = {
        id: '1',
        name: 'John Doe',
        email: email,
        provider: 'email'
      };
      
      const mockToken = 'mock-jwt-token';
      localStorage.setItem('authToken', mockToken);
      setUser(mockUser);
      
      return { success: true };
    } catch (error) {
      console.error('Login failed:', error);
      return { success: false, error: 'Login failed. Please try again.' };
    }
  };

  const signup = async (name, email, password) => {
    try {
      // TODO: Implement actual signup API call
      console.log('Signup attempt:', { name, email, password });
      
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Mock successful signup
      const mockUser = {
        id: '1',
        name: name,
        email: email,
        provider: 'email'
      };
      
      const mockToken = 'mock-jwt-token';
      localStorage.setItem('authToken', mockToken);
      setUser(mockUser);
      
      return { success: true };
    } catch (error) {
      console.error('Signup failed:', error);
      return { success: false, error: 'Signup failed. Please try again.' };
    }
  };

  const googleAuth = () => {
    // TODO: Implement Google OAuth redirect
    console.log('Google auth redirect');
    window.location.href = '/api/auth/google';
  };

  const logout = () => {
    localStorage.removeItem('authToken');
    setUser(null);
  };

  const value = {
    user,
    isLoading,
    login,
    signup,
    googleAuth,
    logout,
    isAuthenticated: !!user
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};
