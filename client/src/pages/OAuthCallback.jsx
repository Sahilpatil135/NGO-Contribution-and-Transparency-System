import { useOAuthCallback } from '../hooks/useOAuthCallback';
import './OAuthCallback.css';

const OAuthCallback = () => {
  const { isProcessing } = useOAuthCallback();

  return (
    <div className="oauth-callback">
      <div className="oauth-callback-content">
        <div className="loading-spinner"></div>
        <h2>Processing Authentication</h2>
        <p>Please wait while we complete your authentication...</p>
      </div>
    </div>
  );
};

export default OAuthCallback;
