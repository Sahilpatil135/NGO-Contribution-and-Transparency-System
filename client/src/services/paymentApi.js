import axios from "axios";

const API_BASE = "http://localhost:8080/api/payment"; // adjust your backend port

export const createOrder = async (amount) => {
  const receipt = "rcpt_" + Math.floor(Math.random() * 100000);
  const response = await axios.post(`${API_BASE}/create-order`, {
    amount,
    receipt,
  });
  return response.data;
};

export const verifyPayment = async (data) => {
  return axios.post(`${API_BASE}/verify`, data);
};
