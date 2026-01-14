import { ENV } from './environment';

// API Configuration
export const API_BASE_URL = ENV.API_BASE_URL;

// WebSocket Base URL (NEW)
export const WS_BASE_URL = ENV.WS_BASE_URL;

export const API_ENDPOINTS = {
  // Authentication endpoints
  REGISTER: `${API_BASE_URL}/api/auth/register`,
  LOGIN: `${API_BASE_URL}/api/auth/login`,
  LOGOUT: `${API_BASE_URL}/api/auth/logout`,
  ME: `${API_BASE_URL}/api/auth/me`,
  GOOGLE_AUTH: `${API_BASE_URL}/api/auth/google`,
  GOOGLE_CALLBACK: `${API_BASE_URL}/api/auth/google/callback`,

  // Payments
  CREATE_ORDER: `${API_BASE_URL}/api/payment/create-order`,
  VERIFY_PAYMENT: `${API_BASE_URL}/api/payment/verify`,

  // Data
  GET_ALL_AID_TYPES: `${API_BASE_URL}/api/aids`,
  GET_ALL_DOMAINS: `${API_BASE_URL}/api/domains`,

  // Causes
  GET_CAUSES: `${API_BASE_URL}/api/causes`,
  GET_CAUSES_BY_DOMAIN: `${API_BASE_URL}/api/causes/domain`,
  GET_CAUSES_BY_AID_TYPE: `${API_BASE_URL}/api/causes/aid`,
  GET_CAUSES_BY_ORGANIZATION: `${API_BASE_URL}/api/causes/organization`,

  // PROOF OF WORK (NEW)
  CREATE_PROOF_SESSION: `${API_BASE_URL}/api/proof/session`,
  UPLOAD_PROOF_IMAGE: (sessionId) =>
    `${API_BASE_URL}/api/proof/upload/${sessionId}`,
};

// API utility functions
export const apiRequest = async (url, options = {}) => {
  const defaultOptions = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  // Add authorization header if token exists
  const token = localStorage.getItem('authToken');
  if (token) {
    defaultOptions.headers.Authorization = `Bearer ${token}`;
  }

  const config = {
    ...defaultOptions,
    ...options,
    headers: {
      ...defaultOptions.headers,
      ...options.headers,
    },
  };

  try {
    const response = await fetch(url, config);

    // Handle non-JSON responses (like redirects)
    const contentType = response.headers.get('content-type');
    if (!contentType || !contentType.includes('application/json')) {
      if (response.ok) {
        return { success: true, data: null };
      }

      throw new Error(`${await response.text()}`);
    }


    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || `HTTP ${response.status}: ${response.statusText}`);
    }

    return { success: true, data };
  } catch (error) {
    console.error('API request failed:', error);
    return { success: false, error: error.message };
  }
};
