import { useEffect, useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

export const useOAuthCallback = () => {
  const { user, isLoading, fetchCurrentUser } = useAuth();
  const navigate = useNavigate();
  const [isProcessing, setIsProcessing] = useState(true);

  useEffect(() => {
    const handleOAuthCallback = async () => {
      try {
        const urlParams = new URLSearchParams(window.location.search);
        const token = urlParams.get('token');
        const error = urlParams.get('error');

        if (error) {
          navigate('/login?error=oauth_failed');
          return;
        }

        if (token) {
          // Persist token and fetch current user
          localStorage.setItem('authToken', token);
          const me = await fetchCurrentUser();
          if (me.success) {
            // Redirect admin users to the admin dashboard
            const destination = me.data?.role === 'admin' ? '/admin' : '/';
            navigate(destination, { replace: true });
          } else {
            navigate('/login?error=oauth_failed');
          }
          return;
        }

        // Fallback: if already authenticated, respect role
        if (user) {
          const destination = user.role === 'admin' ? '/admin' : '/';
          navigate(destination, { replace: true });
          return;
        }

        navigate('/login');
      } catch (error) {
        console.error('OAuth callback error:', error);
        navigate('/login?error=oauth_failed');
      } finally {
        setIsProcessing(false);
      }
    };

    if (!isLoading) {
      handleOAuthCallback();
    }
  }, [user, isLoading, navigate]);

  return { isProcessing };
};
