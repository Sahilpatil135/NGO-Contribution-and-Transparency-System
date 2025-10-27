import axios from "axios";
import { API_ENDPOINTS } from "../config/api";

export const createOrder = async (amount) => {
  const receipt = "rcpt_" + Math.floor(Math.random() * 100000);
  try {
    const response = await axios.post(API_ENDPOINTS.CREATE_ORDER, {
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
    const response = await axios.post(API_ENDPOINTS.VERIFY_PAYMENT, data);
    return response;
  } catch (error) {
    console.error("Error verifying payment:", error);
    throw new Error("Failed to verify payment");
  }
};
