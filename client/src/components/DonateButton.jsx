import React, { useState } from "react";
import { createOrder, verifyPayment } from "../services/paymentApi";

const DonateButton = ({ amount, donorInfo = {} }) => {
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
