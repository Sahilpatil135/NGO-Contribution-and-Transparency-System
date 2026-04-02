import { useState, useEffect, useCallback } from "react";
import { useLocation } from "react-router-dom";
import { apiRequest, API_ENDPOINTS } from "../config/api";

const UUID_RE =
  /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i;

const initialForm = {
  cause_id: "",
  full_name: "",
  age: "",
  blood_group: "",
  phone: "",
  email: "",
  village: "",
  city: "",
  district: "",
  state: "",
  last_donation_date: "",
  availability: true,
  medical_conditions: "",
  consent: false,
};

export default function BloodDonationPage() {
  const location = useLocation();
  const [form, setForm] = useState(initialForm);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const [eligibility, setEligibility] = useState(null);
  const [eligibilityLoading, setEligibilityLoading] = useState(false);

  useEffect(() => {
    const id = location.state?.causeID;
    if (id != null && String(id).trim() !== "") {
      setForm((prev) => ({ ...prev, cause_id: String(id) }));
    }
  }, [location.state]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const loadEligibility = useCallback(async () => {
    setEligibilityLoading(true);
    const result = await apiRequest(API_ENDPOINTS.CHECK_BLOOD_DONATION_ELIGIBILITY);
    setEligibilityLoading(false);

    if (!result.success) {
      setEligibility(null);
      setError(result.error || "Unable to verify blood donation eligibility.");
      return null;
    }

    setError("");
    setEligibility(result.data);
    return result.data;
  }, []);

  useEffect(() => {
    loadEligibility();
  }, [loadEligibility]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (!form.full_name.trim() || !form.phone.trim() || !form.blood_group.trim()) {
      setError("Full name, phone and blood group are required.");
      return;
    }

    if (!form.consent) {
      setError("Please provide consent before submitting.");
      return;
    }

    const latestEligibility = await loadEligibility();
    if (!latestEligibility) {
      return;
    }
    if (!latestEligibility.eligible) {
      setError(
        latestEligibility.eligibility_message ||
          "You are not eligible to donate right now."
      );
      return;
    }

    const causeIdTrimmed = form.cause_id.trim();
    if (causeIdTrimmed && !UUID_RE.test(causeIdTrimmed)) {
      setError("Cause ID must be a valid campaign ID (UUID), or leave it empty.");
      return;
    }

    const ageNumber = Number(form.age);
    if (!Number.isFinite(ageNumber) || ageNumber <= 0) {
      setError("Age must be a valid number greater than zero.");
      return;
    }

    const payload = {
      full_name: form.full_name.trim(),
      age: ageNumber,
      blood_group: form.blood_group.trim().toUpperCase(),
      phone: form.phone.trim(),
      consent: form.consent,
      availability: form.availability,
    };

    if (causeIdTrimmed) payload.cause_id = causeIdTrimmed;
    if (form.email.trim()) payload.email = form.email.trim();
    if (form.village.trim()) payload.village = form.village.trim();
    if (form.city.trim()) payload.city = form.city.trim();
    if (form.district.trim()) payload.district = form.district.trim();
    if (form.state.trim()) payload.state = form.state.trim();
    if (form.medical_conditions.trim()) {
      payload.medical_conditions = form.medical_conditions.trim();
    }
    if (!latestEligibility.has_verified_record && form.last_donation_date) {
      payload.last_donation_date = form.last_donation_date;
    }

    setSubmitting(true);
    const result = await apiRequest(API_ENDPOINTS.CREATE_BLOOD_DONOR, {
      method: "POST",
      body: JSON.stringify(payload),
    });
    setSubmitting(false);

    if (!result.success) {
      setError(result.error || "Failed to submit blood donor information.");
      return;
    }

    setSuccess("Your donor details have been submitted successfully.");
    setForm(initialForm);
    loadEligibility();
  };

  return (
    <div className="max-w-3xl mx-auto px-4 py-8">
      <div className="bg-white border border-gray-200 rounded-xl shadow-sm p-6">
        <h1 className="text-2xl font-bold text-[#3a0b2e] mb-2">Blood Donor Form</h1>
        <p className="text-gray-600 mb-6">
          Enter your details to register as a blood donor.
        </p>

        {error && <p className="mb-4 text-sm text-red-600">{error}</p>}
        {success && <p className="mb-4 text-sm text-green-700">{success}</p>}

        <form onSubmit={handleSubmit} className="space-y-4">
          <label className="block text-sm font-medium mb-2">
            Cause ID <span className="text-red-500">*</span>
          </label>
          <input
            name="cause_id"
            value={form.cause_id}
            onChange={handleChange}
            placeholder="Cause ID"
            className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
            required
          />
          <label className="block text-sm font-medium mb-2">
            Full Name <span className="text-red-500">*</span>
          </label>
          <input
            name="full_name"
            value={form.full_name}
            onChange={handleChange}
            placeholder="Full Name"
            className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
            required
          />

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Age <span className="text-red-500">*</span>
              </label>
              <input
                name="age"
                value={form.age}
                onChange={handleChange}
                placeholder="Age"
                type="number"
                min="1"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Blood Group <span className="text-red-500">*</span>
              </label>
              <input
                name="blood_group"
                value={form.blood_group}
                onChange={handleChange}
                placeholder="Blood Group (e.g. O+, AB-)"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Phone <span className="text-red-500">*</span>
              </label>
              <input
                name="phone"
                value={form.phone}
                onChange={handleChange}
                placeholder="Phone"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Email <span className="text-red-500">*</span>
              </label>
              <input
                name="email"
                value={form.email}
                onChange={handleChange}
                placeholder="Email"
                type="email"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
              />
            </div>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Village <span className="text-red-500">*</span>
              </label>
              <input
                name="village"
                value={form.village}
                onChange={handleChange}
                placeholder="Village"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                City <span className="text-red-500">*</span>
              </label>
              <input
                name="city"
                value={form.city}
                onChange={handleChange}
                placeholder="City"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                District <span className="text-red-500">*</span>
              </label>
              <input
                name="district"
                value={form.district}
                onChange={handleChange}
                placeholder="District"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                State <span className="text-red-500">*</span>
              </label>
              <input
                name="state"
                value={form.state}
                onChange={handleChange}
                placeholder="State"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
          </div>
          {eligibilityLoading && (
            <p className="text-sm text-gray-600">
              Checking donation eligibility...
            </p>
          )}

          {eligibility?.has_verified_record ? (
            <div className="rounded-lg border border-gray-200 bg-gray-50 p-3 text-sm">
              <p className="text-gray-700">
                Latest verified donation date:{" "}
                <span className="font-semibold">
                  {eligibility.latest_verified_date}
                </span>
              </p>
              <p
                className={`mt-1 ${
                  eligibility.eligible ? "text-green-700" : "text-red-600"
                }`}
              >
                {eligibility.eligibility_message}
              </p>
            </div>
          ) : (
            <>
              <label className="block text-sm font-medium mb-2">
                Last Donation Date (optional)
              </label>
              <input
                name="last_donation_date"
                value={form.last_donation_date}
                onChange={handleChange}
                type="date"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
              />
            </>
          )}
          <label className="block text-sm font-medium mb-2">
            Medical Conditions (optional)
          </label>
          <textarea
            name="medical_conditions"
            value={form.medical_conditions}
            onChange={handleChange}
            placeholder="Mention any known conditions (e.g., diabetes, BP, recent illness)"
            rows="3"
            className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
          />

          <label className="flex items-center gap-2 text-sm text-gray-700">
            <input
              name="availability"
              checked={form.availability}
              onChange={handleChange}
              type="checkbox"
            />
            I am currently available for blood donation
          </label>

          <label className="flex items-center gap-2 text-sm text-gray-700">
            <input
              name="consent"
              checked={form.consent}
              onChange={handleChange}
              type="checkbox"
              required
            />
            I consent to share these details for blood donation requests *
          </label>

          <button
            type="submit"
            disabled={submitting || (eligibility?.has_verified_record && !eligibility?.eligible)}
            className="bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-5 py-2 rounded-lg transition cursor-pointer disabled:opacity-60"
          >
            {submitting ? "Submitting..." : "Submit"}
          </button>
        </form>
      </div>
    </div>
  );
}
