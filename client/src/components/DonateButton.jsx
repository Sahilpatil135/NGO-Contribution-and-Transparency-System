import React, { useState } from "react";
import { createOrder, verifyPayment } from "../services/paymentApi";
import "./DonateButton.css";

const DonateButton = ({ amount }) => {
  const handleDonate = async () => {
    try {
      const order = await createOrder(amount);

      const options = {
        key: order.key,
        amount: order.amount,
        currency: order.currency,
        name: "Your NGO Name",
        description: "Donation Support",
        order_id: order.orderId,
        handler: async function (response) {
          const data = {
            order_id: response.razorpay_order_id,
            payment_id: response.razorpay_payment_id,
            signature: response.razorpay_signature,
          };

          const result = await verifyPayment(data);
          if (result.data.success) {
            alert("🎉 Donation Successful! Thank you for your support!");
          } else {
            alert("❌ Payment verification failed!");
          }
        },
        prefill: {
          name: "Your Donor",
          email: "donor@example.com",
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
      onClick={handleDonate}
      className="donate-button"  
      style={buttonStyle}
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
    >
      Donate ₹{amount}
    </button>
  );
};

export default DonateButton;
