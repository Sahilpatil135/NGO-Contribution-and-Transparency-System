import { ENV } from './environment';

// API Configuration
export const API_BASE_URL = ENV.API_BASE_URL;

// WebSocket Base URL (NEW)
export const WS_BASE_URL = ENV.WS_BASE_URL;

export const API_ENDPOINTS = {
  // Authentication endpoints
  REGISTER: `${API_BASE_URL}/api/auth/register`,
  REGISTER_ORGANIZATION: `${API_BASE_URL}/api/auth/register/organization`,
  LOGIN: `${API_BASE_URL}/api/auth/login`,
  LOGOUT: `${API_BASE_URL}/api/auth/logout`,
  ME: `${API_BASE_URL}/api/auth/me`,
  UPDATE_PROFILE: `${API_BASE_URL}/api/auth/me`,
  GOOGLE_AUTH: `${API_BASE_URL}/api/auth/google`,
  GOOGLE_CALLBACK: `${API_BASE_URL}/api/auth/google/callback`,

  // Payments
  CREATE_ORDER: `${API_BASE_URL}/api/payment/create-order`,
  VERIFY_PAYMENT: `${API_BASE_URL}/api/payment/verify`,

  // Donations
  CREATE_DONATION: `${API_BASE_URL}/api/donations`,
  GET_MY_DONATIONS: `${API_BASE_URL}/api/donations/user/me`,

  // Data
  GET_ALL_AID_TYPES: `${API_BASE_URL}/api/aids`,
  GET_ALL_DOMAINS: `${API_BASE_URL}/api/domains`,

  // Causes
  GET_CAUSES: `${API_BASE_URL}/api/causes`,
  CREATE_CAUSE: `${API_BASE_URL}/api/causes`,
  CREATE_BLOOD_DONOR: `${API_BASE_URL}/api/causes/blood`,
  CHECK_BLOOD_DONATION_ELIGIBILITY: `${API_BASE_URL}/api/causes/blood/eligibility`,
  CREATE_VOLUNTEER: `${API_BASE_URL}/api/causes/volunteer`,
  UPLOAD_CAUSE_COVER: `${API_BASE_URL}/api/causes/cover/upload`,
  UPLOAD_PRODUCT_IMAGE: `${API_BASE_URL}/api/causes/products/upload`,
  UPLOAD_UPDATE_RECEIPT: `${API_BASE_URL}/api/causes/updates/upload/receipt`,
  CREATE_CAUSE_UPDATE: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/updates`,
  GET_CAUSES_BY_DOMAIN: `${API_BASE_URL}/api/causes/domain`,
  GET_CAUSES_BY_AID_TYPE: `${API_BASE_URL}/api/causes/aid`,
  GET_CAUSES_BY_ORGANIZATION: `${API_BASE_URL}/api/causes/organization`,
  ME_ORGANIZATION: `${API_BASE_URL}/api/auth/me/organization`,

  // Disbursements
  GET_MY_ORGANIZATION_DISBURSEMENTS: `${API_BASE_URL}/api/disbursements/my-organization`,

  // Cause votes
  GET_CAUSE_VOTES: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/votes`,
  UPVOTE_CAUSE: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/upvote`,
  DOWNVOTE_CAUSE: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/downvote`,

  // Cause reviews
  GET_CAUSE_REVIEWS: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/reviews`,
  GET_CAUSE_REVIEW_COUNT: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/reviews/count`,
  CREATE_CAUSE_REVIEW: (causeId) =>
    `${API_BASE_URL}/api/causes/${causeId}/reviews`,

  // Donations (on-chain)
  GET_CAUSE_CHAIN_DONATIONS: (causeId) =>
    `${API_BASE_URL}/api/donations/chain/cause/${causeId}`,

  // PROOF OF WORK (NEW)
  CREATE_PROOF_SESSION: `${API_BASE_URL}/api/proof/session`,
  CREATE_CAUSE_PROOF_SESSION: `${API_BASE_URL}/api/proof/session/cause`,
  UPLOAD_PROOF_IMAGE: (sessionId) =>
    `${API_BASE_URL}/api/proof/upload/${sessionId}`,
  GET_PROOF_IMAGES_BY_SESSION: (sessionId) =>
    `${API_BASE_URL}/api/proof/sessions/${sessionId}/images`,
  // new {
  // Receipt verification polling (Execution updates)
  GET_RECEIPT_STATUS: (receiptJobId) =>
    `${API_BASE_URL}/api/causes/updates/receipt-status/${receiptJobId}`,
  // }
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
