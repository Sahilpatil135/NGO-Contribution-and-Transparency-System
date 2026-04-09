import React, { useState, useEffect } from "react";
import img1 from "../../public/domains/domain_example.png";
import { LuLock } from "react-icons/lu";
import { useLocation } from "react-router-dom";
import DonateButton from "../components/DonateButton";
import { apiRequest, API_ENDPOINTS } from "../config/api";
import { getCauseImage } from "../utils/imageHelper";
import { formatGoal, formatCollected, getCollectedLabel } from "../utils/goalHelper";
import { ValidationRules, FormField, ErrorMessage } from "../components/FormValidation";

const CheckoutPage = () => {
  const location = useLocation();
  const causeId = location.state?.causeID;

  const [amount, setAmount] = useState(500);
  const [isAnonymous, setIsAnonymous] = useState(false);
  const [isIndianCitizen, setIsIndianCitizen] = useState(false);
  const [cause, setCause] = useState(null);
  const [loadingCause, setLoadingCause] = useState(true);
  const [form, setForm] = useState({
    name: "",
    mobile: "",
    email: "",
    address: "",
    pincode: "",
    pan: "",
  });
  const [errors, setErrors] = useState({});
  const [touched, setTouched] = useState({});

  const suggestedAmounts = [100, 500, 1000, 2500, 5000];

  // Fetch cause details
  useEffect(() => {
    const fetchCause = async () => {
      if (!causeId) {
        setLoadingCause(false);
        return;
      }

      try {
        const res = await apiRequest(`${API_ENDPOINTS.GET_CAUSES}/${causeId}`);
        if (res.success && res.data) {
          setCause(res.data);
        }
      } catch (error) {
        console.error("Failed to load cause:", error);
      } finally {
        setLoadingCause(false);
      }
    };

    fetchCause();
  }, [causeId]);

  // Calculate progress
  const aidTypeName = cause?.aid_type?.name || "";
  const collected = Number(cause?.collected_amount) || 0;
  const goal = Number(cause?.goal_amount) || 0;
  const progressPercent = goal > 0 ? Math.min((collected / goal) * 100, 100) : 0;

  // Prevent scroll-based value change
  const disableScroll = (e) => e.target.blur();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
    
    // Clear error when user starts typing
    if (errors[name]) {
      setErrors({ ...errors, [name]: "" });
    }
  };

  const handleBlur = (field) => {
    setTouched({ ...touched, [field]: true });
    validateField(field, form[field]);
  };

  const validateField = (field, value) => {
    let error = "";
    
    switch (field) {
      case "name":
        error = ValidationRules.required(value) || ValidationRules.minLength(3)(value);
        break;
      case "mobile":
        error = ValidationRules.required(value) || ValidationRules.phone(value);
        break;
      case "email":
        error = ValidationRules.required(value) || ValidationRules.email(value);
        break;
      case "address":
        error = ValidationRules.required(value) || ValidationRules.minLength(10)(value);
        break;
      case "pincode":
        const pincodeRegex = /^[1-9][0-9]{5}$/;
        error = ValidationRules.required(value) || 
                (!pincodeRegex.test(value) ? "Please enter a valid 6-digit Indian pincode" : "");
        break;
      case "pan":
        const panRegex = /^[A-Z]{5}[0-9]{4}[A-Z]{1}$/;
        error = ValidationRules.required(value) || 
                (!panRegex.test(value.toUpperCase()) ? "Please enter a valid PAN (e.g., ABCDE1234F)" : "");
        break;
      default:
        break;
    }
    
    setErrors({ ...errors, [field]: error });
    return error;
  };

  // Validation
  const validateForm = () => {
    const newErrors = {};
    let isValid = true;

    // Validate all fields
    Object.keys(form).forEach(field => {
      const error = validateField(field, form[field]);
      if (error) {
        newErrors[field] = error;
        isValid = false;
      }
    });

    // Check Indian citizen confirmation
    if (!isIndianCitizen) {
      newErrors.citizenship = "Please confirm that you are an Indian citizen.";
      isValid = false;
    }

    setErrors(newErrors);
    setTouched({
      name: true,
      mobile: true,
      email: true,
      address: true,
      pincode: true,
      pan: true,
    });

    return isValid;
  };

  // const handleSubmit = () => {
  //   if (validateForm()) {
  //     alert("Form submitted successfully!");
  //   }
  // };
  const handleProceedToPay = () => {
    if (validateForm()) {
      document.getElementById("rzp-trigger").click();
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
                  className={`py-2 rounded-lg border cursor-pointer transition ${amount === val
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
                autoComplete="transaction-amount"
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
                  autoComplete="name"
                  type="text"
                  name="name"
                  value={form.name}
                  onChange={handleChange}
                  onBlur={() => handleBlur("name")}
                  placeholder="Enter your full name"
                  className={`w-full border rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400 ${
                    touched.name && errors.name ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                {touched.name && errors.name && (
                  <p className="text-red-600 text-sm mt-1">{errors.name}</p>
                )}
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
                  autoComplete="mobile tel"
                  type="tel"
                  name="mobile"
                  value={form.mobile}
                  onChange={handleChange}
                  onBlur={() => handleBlur("mobile")}
                  placeholder="10-digit Indian mobile number (e.g., 9876543210)"
                  maxLength="10"
                  className={`w-full border rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400 ${
                    touched.mobile && errors.mobile ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                {touched.mobile && errors.mobile && (
                  <p className="text-red-600 text-sm mt-1">{errors.mobile}</p>
                )}
              </div>

              {/* Email Address */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Email Address
                </label>
                <input
                  autoComplete="email"
                  type="email"
                  name="email"
                  value={form.email}
                  onChange={handleChange}
                  onBlur={() => handleBlur("email")}
                  placeholder="Enter your email address"
                  className={`w-full border rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400 ${
                    touched.email && errors.email ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                {touched.email && errors.email && (
                  <p className="text-red-600 text-sm mt-1">{errors.email}</p>
                )}
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
                  onBlur={() => handleBlur("address")}
                  placeholder="Enter your complete address"
                  rows="2"
                  className={`w-full border rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400 ${
                    touched.address && errors.address ? 'border-red-500' : 'border-gray-300'
                  }`}
                ></textarea>
                {touched.address && errors.address && (
                  <p className="text-red-600 text-sm mt-1">{errors.address}</p>
                )}
              </div>

              {/* Pincode */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  Pincode
                </label>
                <input
                  autoComplete="postal-code"
                  type="text"
                  name="pincode"
                  value={form.pincode}
                  onChange={handleChange}
                  onBlur={() => handleBlur("pincode")}
                  placeholder="6-digit Indian pincode (e.g., 400001)"
                  maxLength="6"
                  className={`w-full border rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400 ${
                    touched.pincode && errors.pincode ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                {touched.pincode && errors.pincode && (
                  <p className="text-red-600 text-sm mt-1">{errors.pincode}</p>
                )}
              </div>

              {/* PAN */}
              <div>
                <label className="block text-gray-700 text-sm font-semibold mb-2">
                  PAN Number
                </label>
                <input
                  autoComplete="pan-number"
                  type="text"
                  name="pan"
                  value={form.pan}
                  onChange={handleChange}
                  onBlur={() => handleBlur("pan")}
                  placeholder="ABCDE1234F"
                  maxLength="10"
                  style={{ textTransform: "uppercase" }}
                  className={`w-full border rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-400 ${
                    touched.pan && errors.pan ? 'border-red-500' : 'border-gray-300'
                  }`}
                />
                {touched.pan && errors.pan && (
                  <p className="text-red-600 text-sm mt-1">{errors.pan}</p>
                )}
              </div>

              {/* Indian Citizen Confirmation */}
              <div>
                <div className="flex items-center mt-2">
                  <input
                    type="checkbox"
                    id="citizen"
                    checked={isIndianCitizen}
                    onChange={() => setIsIndianCitizen(!isIndianCitizen)}
                    className="mr-2 accent-[#ff6200] cursor-pointer"
                  />
                  <label htmlFor="citizen" className="text-sm text-gray-600">
                    I confirm that I am an Indian citizen <span className="text-red-500">*</span>
                  </label>
                </div>
                {errors.citizenship && Object.keys(touched).length > 0 && (
                  <p className="text-red-600 text-sm mt-1">{errors.citizenship}</p>
                )}
              </div>

              {/* <button
                onClick={handleSubmit}
                className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 my-2 rounded-lg transition cursor-pointer"
              >
                Proceed to Pay ₹{amount}
              </button> */}

              {/* Proceed to Pay */}
              <button
                onClick={handleProceedToPay}
                className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 my-2 rounded-lg transition cursor-pointer"
              >
                Proceed to Pay ₹{amount}
              </button>

              {/* Hidden DonateButton trigger */}
              <div className="hidden">
                <DonateButton
                  id="rzp-trigger"
                  amount={parseInt(amount)}
                  donorInfo={{
                    name: form.name,
                    email: form.email,
                    mobile: form.mobile,
                    address: form.address,
                    pincode: form.pincode,
                    pan: form.pan,
                  }}
                  causeId={causeId}
                />
              </div>

              <p className="flex justify-center text-center text-sm text-gray-500 mt-1">
                <LuLock className="ml-4 mr-1 h-5" />All payments go through a secure gateway
              </p>
            </div>
          </div>
        </div>

        {/* RIGHT SIDE */}
        <div className="bg-white p-6 rounded-xl shadow-md h-fit sticky top-24">
          {loadingCause ? (
            <div className="text-center py-8">
              <p className="text-gray-600">Loading cause details...</p>
            </div>
          ) : cause ? (
            <>
              <img
                src={getCauseImage(cause.id)}
                alt={cause.title}
                className="w-full h-48 object-cover rounded-lg mb-4"
                onError={(e) => {
                  e.target.src = img1;
                }}
              />
              <h3 className="text-lg font-bold text-[#3a0b2e] mb-1">
                {cause.title}
              </h3>
              <p className="text-gray-600 text-sm mb-4">
                Campaign by {cause.organization?.name || "Organization"}
              </p>

              {goal > 0 && (
                <>
                  <div className="w-full bg-gray-300 rounded-full h-3 mb-4">
                    <div
                      className="h-full bg-[#ff6200] rounded-full transition-all"
                      style={{ width: `${progressPercent}%` }}
                    ></div>
                  </div>

                  <div className="flex justify-between text-sm mb-6">
                    <p className="text-[#3a0b2e] font-semibold">
                      {formatCollected(collected, aidTypeName)} {getCollectedLabel(aidTypeName).toLowerCase()}
                    </p>
                    <p className="text-gray-600">
                      of {formatGoal(goal, aidTypeName)} goal
                    </p>
                  </div>
                </>
              )}

              {/* Mirror button for right side */}
              <button
                onClick={handleProceedToPay}
                className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition cursor-pointer"
              >
                Proceed to Pay ₹{amount}
              </button>
            </>
          ) : (
            <div className="text-center py-8">
              <p className="text-gray-600">Cause not found</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default CheckoutPage;
