import React, { useState } from "react";
import { createOrder, verifyPayment } from "../services/paymentApi";
import { apiRequest, API_ENDPOINTS } from "../config/api";
import { useAuth } from "../contexts/AuthContext";
import { useNavigate } from "react-router-dom";

const DonateButton = ({ amount, donorInfo = {}, causeId }) => {
  const { user } = useAuth();
  const navigate = useNavigate();

  const handleDonate = async () => {
    try {
      const order = await createOrder(amount);

      const options = {
        key: order.key,
        amount: order.amount,
        currency: order.currency,
        name: "CharityLight",
        description: "Donation Support",
        order_id: order.orderId,
        handler: async function(response) {
          const data = {
            order_id: response.razorpay_order_id,
            payment_id: response.razorpay_payment_id,
            signature: response.razorpay_signature,
          };

          const result = await verifyPayment(data);
          if (result.data.success) {
            // Record donation in backend
            try {
              if (!user) {
                console.warn("User not authenticated, skipping donation record.");
              } else if (!causeId) {
                console.warn("No causeId provided, skipping donation record.");
              } else {
                const payload = {
                  cause_id: causeId,
                  user_id: user.id,
                  name:
                    donorInfo.name && donorInfo.name.trim().length > 0
                      ? donorInfo.name
                      : user.name || "Donor",
                  phone: donorInfo.mobile,
                  billing_address: donorInfo.address || undefined,
                  pincode: donorInfo.pincode || undefined,
                  amount: Number(amount),
                  pan_number: donorInfo.pan || undefined,
                  payment_id: response.razorpay_payment_id,
                };

                const donationResult = await apiRequest(
                  API_ENDPOINTS.CREATE_DONATION,
                  {
                    method: "POST",
                    body: JSON.stringify(payload),
                  }
                );

                if (!donationResult.success) {
                  console.error(
                    "Failed to record donation:",
                    donationResult.error
                  );
                }
              }
            } catch (err) {
              console.error("Error while recording donation:", err);
            }

            navigate("/donation/success", {
              replace: true,
              state: {
                amount: Number(amount),
                causeId,
                paymentId: response.razorpay_payment_id,
              },
            });
          } else {
            alert("❌ Payment verification failed!");
          }
        },
        prefill: {
          name: donorInfo.name || "Donor",
          email: donorInfo.email || "donor@example.com",
          contact: donorInfo.mobile || "",
        },
        theme: {
          color: "#3399cc",
        },
      };

      const rzp = new window.Razorpay(options);
      rzp.open();
    } catch (error) {
      console.error("Payment error:", error);
      alert("Something went wrong during payment.");
    }
  };

  const [hover, setHover] = useState(false);

  const buttonStyle = {
    color: hover ? "white" : "black", // Change color on hover
  };

  return (
    <button
      id="rzp-trigger"
      onClick={handleDonate}
      className="bg-[#28a745] text-white border-none px-5 py-2 rounded-md cursor-pointer font-medium transition-colors duration-300 hover:bg-[#218838]"
      style={buttonStyle}
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
    >
      Donate ₹{amount}
    </button>
  );
};

export default DonateButton;
