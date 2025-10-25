import React, { useState } from "react";
import img1 from "../assets/img1.png";

const CheckoutPage = () => {
  const [amount, setAmount] = useState(500);
  const [isAnonymous, setIsAnonymous] = useState(false);

  const suggestedAmounts = [100, 500, 1000, 2500, 5000];

  const handleAmountClick = (val) => setAmount(val);
  const handleChange = (e) => setAmount(e.target.value);

  return (
    <div className="bg-gray-100 py-12">
      <div className="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-[70%_30%] gap-10 px-4 md:px-8">
        
        {/* LEFT SIDE */}
        <div className="flex flex-col space-y-6">
          {/* Donation Amount Card */}
          <div className="bg-white p-6 rounded-xl shadow-md">
            <h2 className="text-2xl font-semibold text-[#3a0b2e] mb-4">
              Choose Donation Amount
            </h2>
            <div className="grid grid-cols-3 sm:grid-cols-5 gap-3 mb-6">
              {suggestedAmounts.map((val) => (
                <button
                  key={val}
                  onClick={() => handleAmountClick(val)}
                  className={`py-2 rounded-lg border transition ${
                    amount === val
                      ? "bg-[#ff6200] text-white border-[#ff6200]"
                      : "bg-white text-[#3a0b2e] border-gray-300 hover:border-[#ff6200]"
                  }`}
                >
                  ₹{val}
                </button>
              ))}
            </div>

            <div className="mb-6">
              <label className="block text-gray-700 text-sm font-semibold mb-2">
                Enter custom amount (₹)
              </label>
              <input
                type="number"
                value={amount}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
              />
            </div>
          </div>

          {/* Donor Details Card */}
          <div className="bg-white p-6 rounded-xl shadow-md">
            <h2 className="text-2xl font-semibold text-[#3a0b2e] mb-4">
              Donor Details
            </h2>

            <div className="space-y-4">
              {/* Full Name + Anonymous Option */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Full Name
                </label>
                <input
                  type="text"
                  placeholder="Enter full name"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                />
                <div className="flex items-center mt-2">
                  <input
                    type="checkbox"
                    id="anonymous"
                    checked={isAnonymous}
                    onChange={() => setIsAnonymous(!isAnonymous)}
                    className="mr-2 accent-[#ff6200]"
                  />
                  <label htmlFor="anonymous" className="text-sm text-gray-600">
                    Make this donation anonymous
                  </label>
                </div>
              </div>

              {/* Indian Mobile Number */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Mobile Number
                </label>
                <input
                  type="tel"
                  placeholder="+91 9876543210"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                />
              </div>

              {/* Email Address */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Email Address
                </label>
                <input
                  type="email"
                  placeholder="Enter email address"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                />
              </div>

              {/* Billing / Home Address */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Billing / Home Address
                </label>
                <textarea
                  placeholder="Enter your full address"
                  rows="2"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                ></textarea>
              </div>

              {/* Pincode */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Pincode
                </label>
                <input
                  type="text"
                  placeholder="e.g. 400001"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                />
              </div>

              {/* PAN Number */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  PAN Number
                </label>
                <input
                  type="text"
                  placeholder="ABCDE1234F"
                  maxLength="10"
                  className="w-full uppercase border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                />
              </div>

              <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition">
                Proceed to Pay ₹{amount}
              </button>
            </div>
          </div>
        </div>

        {/* RIGHT SIDE */}
        <div className="bg-white p-6 rounded-xl shadow-md h-fit sticky top-24">
          <img
            src={img1}
            alt="Campaign"
            className="w-full h-48 object-cover rounded-lg mb-4"
          />
          <h3 className="text-lg font-bold text-[#3a0b2e] mb-1">
            Help Dr Amruta Provide Life-Saving Cancer Treatment
          </h3>
          <p className="text-gray-600 text-sm mb-4">Campaign by Dr. Amruta Tripathi</p>

          {/* Progress Bar */}
          <div className="w-full bg-gray-300 rounded-full h-3 mb-4">
            <div
              className="h-full bg-[#ff6200] rounded-full"
              style={{ width: "80%" }}
            ></div>
          </div>

          {/* Stats */}
          <div className="flex justify-between text-sm mb-6">
            <p className="text-[#3a0b2e] font-semibold">₹4,00,000 raised</p>
            <p className="text-gray-600">of ₹5,00,000 goal</p>
          </div>

          <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition">
            Proceed to Pay ₹{amount}
          </button>
        </div>
      </div>
    </div>
  );
};

export default CheckoutPage;
