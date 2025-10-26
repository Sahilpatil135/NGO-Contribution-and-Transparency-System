import axios from "axios";

const API_BASE = "http://localhost:8080/api/payment"; // adjust your backend port

export const createOrder = async (amount) => {
  const receipt = "rcpt_" + Math.floor(Math.random() * 100000);
  try {
    const response = await axios.post(`${API_BASE}/create-order`, {
      amount,
      receipt,
    });
    return response.data;
  } catch (error) {
    console.error("Error creating order:", error);
    throw new Error("Failed to create payment order");
  }
};

export const verifyPayment = async (data) => {
  try {
    const response = await axios.post(`${API_BASE}/verify`, data);
    return response;
  } catch (error) {
    console.error("Error verifying payment:", error);
    throw new Error("Failed to verify payment");
  }
};
