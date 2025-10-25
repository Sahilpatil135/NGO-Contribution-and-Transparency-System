import React, { useState } from "react";
import img1 from "../assets/img1.png";
import { LuLock } from "react-icons/lu";

const CheckoutPage = () => {
  const [amount, setAmount] = useState(500);
  const [isAnonymous, setIsAnonymous] = useState(false);
  const [isIndianCitizen, setIsIndianCitizen] = useState(false);
  const [form, setForm] = useState({
    name: "",
    mobile: "",
    email: "",
    address: "",
    pincode: "",
    pan: "",
  });

  const suggestedAmounts = [100, 500, 1000, 2500, 5000];

  // Prevent scroll-based value change
  const disableScroll = (e) => e.target.blur();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
  };

  // Validation
  const validateForm = () => {
    const mobileRegex = /^[6-9]\d{9}$/;
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    const pincodeRegex = /^[1-9][0-9]{5}$/;
    const panRegex = /^[A-Z]{5}[0-9]{4}[A-Z]{1}$/;

    if (!mobileRegex.test(form.mobile)) {
      alert("Please enter a valid 10-digit Indian mobile number.");
      return false;
    }
    if (!emailRegex.test(form.email)) {
      alert("Please enter a valid email address.");
      return false;
    }
    if (!pincodeRegex.test(form.pincode)) {
      alert("Please enter a valid 6-digit Indian pincode.");
      return false;
    }
    if (!panRegex.test(form.pan)) {
      alert("Please enter a valid PAN number (e.g., ABCDE1234F).");
      return false;
    }
    if (!isIndianCitizen) {
      alert("Please confirm that you are an Indian citizen.");
      return false;
    }
    return true;
  };

  const handleSubmit = () => {
    if (validateForm()) {
      alert("Form submitted successfully!");
    }
  };

  return (
    <div className="bg-gray-100 py-12">
      <div className="max-w-7xl mx-auto grid grid-cols-1 md:grid-cols-[70%_30%] gap-10 px-4 md:px-8">
        
        {/* LEFT SIDE */}
        <div className="flex flex-col space-y-10">
          
          {/* Donation Amount Card */}
          <div className="bg-white p-6 px-8 rounded-xl shadow-md">
            <h2 className="text-2xl font-semibold text-[#3a0b2e] mb-4">
              Choose Donation Amount
            </h2>
            <div className="grid grid-cols-3 sm:grid-cols-5 gap-3 mb-6">
              {suggestedAmounts.map((val) => (
                <button
                  key={val}
                  onClick={() => setAmount(val)}
                  className={`py-2 rounded-lg border cursor-pointer transition ${
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
                onChange={(e) => setAmount(e.target.value)}
                onWheel={disableScroll}
                min="1"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] no-spinner"
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
                  name="name"
                  value={form.name}
                  onChange={handleChange}
                  placeholder="Enter your full name"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400"
                />
                <div className="flex items-center mt-2">
                  <input
                    type="checkbox"
                    id="anonymous"
                    checked={isAnonymous}
                    onChange={() => setIsAnonymous(!isAnonymous)}
                    className="mr-2 accent-[#ff6200] cursor-pointer"
                  />
                  <label htmlFor="anonymous" className="text-sm text-gray-600">
                    Make this donation anonymous
                  </label>
                </div>
              </div>

              {/* Mobile Number */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Mobile Number
                </label>
                <input
                  type="tel"
                  name="mobile"
                  value={form.mobile}
                  onChange={handleChange}
                  placeholder="10-digit Indian mobile number (e.g., 9876543210)"
                  maxLength="10"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400"
                />
              </div>

              {/* Email Address */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Email Address
                </label>
                <input
                  type="email"
                  name="email"
                  value={form.email}
                  onChange={handleChange}
                  placeholder="Enter your email address"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400"
                />
              </div>

              {/* Address */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Billing / Home Address
                </label>
                <textarea
                  name="address"
                  value={form.address}
                  onChange={handleChange}
                  placeholder="Enter your complete address"
                  rows="2"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400"
                ></textarea>
              </div>

              {/* Pincode */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Pincode
                </label>
                <input
                  type="text"
                  name="pincode"
                  value={form.pincode}
                  onChange={handleChange}
                  placeholder="6-digit Indian pincode (e.g., 400001)"
                  maxLength="6"
                  className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400"
                />
              </div>

              {/* PAN */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  PAN Number
                </label>
                <input
                  type="text"
                  name="pan"
                  value={form.pan}
                  onChange={handleChange}
                  placeholder="ABCDE1234F"
                  maxLength="10"
                  className="w-full uppercase border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400"
                />
              </div>

              {/* Indian Citizen Confirmation */}
              <div className="flex items-center mt-2">
                <input
                  type="checkbox"
                  id="citizen"
                  checked={isIndianCitizen}
                  onChange={() => setIsIndianCitizen(!isIndianCitizen)}
                  className="mr-2 accent-[#ff6200] cursor-pointer"
                />
                <label htmlFor="citizen" className="text-sm text-gray-600">
                  I confirm that I am an Indian citizen
                </label>
              </div>

              <button
                onClick={handleSubmit}
                className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 my-2 rounded-lg transition cursor-pointer"
              >
                Proceed to Pay ₹{amount}
              </button>

              <p className="flex justify-center text-center text-sm text-gray-500 mt-1">
                <LuLock className="ml-4 mr-1 h-5"/>All payments go through a secure gateway
              </p>
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

          <div className="w-full bg-gray-300 rounded-full h-3 mb-4">
            <div
              className="h-full bg-[#ff6200] rounded-full"
              style={{ width: "80%" }}
            ></div>
          </div>

          <div className="flex justify-between text-sm mb-6">
            <p className="text-[#3a0b2e] font-semibold">₹4,00,000 raised</p>
            <p className="text-gray-600">of ₹5,00,000 goal</p>
          </div>

          <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition cursor-pointer">
            Proceed to Pay ₹{amount}
          </button>
        </div>
      </div>
    </div>
  );
};

export default CheckoutPage;
