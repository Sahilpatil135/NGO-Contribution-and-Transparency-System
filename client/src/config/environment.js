// Environment configuration
export const ENV = {
  // Base API
  API_BASE_URL:
    import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',

  // Server details (NEW)
  SERVER_IP: import.meta.env.VITE_SERVER_IP || 'localhost',
  BACKEND_PORT: import.meta.env.VITE_BACKEND_PORT || '8080',
  FRONTEND_PORT: import.meta.env.VITE_FRONTEND_PORT || '3000',

  // Derived URLs (NEW)
  WS_BASE_URL: `ws://${import.meta.env.VITE_SERVER_IP || 'localhost'}:${
    import.meta.env.VITE_BACKEND_PORT || '8080'
  }`,

  FRONTEND_BASE_URL: `http://${import.meta.env.VITE_SERVER_IP || 'localhost'}:${
    import.meta.env.VITE_FRONTEND_PORT || '3000'
  }`,

  // Environment
  NODE_ENV: import.meta.env.NODE_ENV || 'development',
  IS_DEVELOPMENT: import.meta.env.NODE_ENV === 'development',
  IS_PRODUCTION: import.meta.env.NODE_ENV === 'production',
};

// Dev log
if (ENV.IS_DEVELOPMENT) {
  console.log('Environment configuration:', ENV);
}

