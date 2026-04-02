import { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import { apiRequest, API_ENDPOINTS } from "../config/api";

const UUID_RE =
  /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i;

const initialForm = {
  cause_id: "",
  full_name: "",
  phone: "",
  email: "",
  village: "",
  city: "",
  district: "",
  state: "",
  skills: "",
  interests: "",
  availability_type: "",
  available_hours: "",
  experience: "",
  consent: false,
};

export default function VolunteerPage() {
  const location = useLocation();
  const [form, setForm] = useState(initialForm);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

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

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    setSuccess("");

    if (!form.full_name.trim() || !form.phone.trim() || !form.skills.trim()) {
      setError("Full name, phone and skills are required.");
      return;
    }

    if (!form.consent) {
      setError("Please provide consent before submitting.");
      return;
    }

    const causeIdTrimmed = form.cause_id.trim();
    if (causeIdTrimmed && !UUID_RE.test(causeIdTrimmed)) {
      setError("Cause ID must be a valid campaign ID (UUID), or leave it empty.");
      return;
    }

    let availableHoursNum = null;
    if (String(form.available_hours).trim() !== "") {
      const n = Number(form.available_hours);
      if (!Number.isFinite(n) || n < 0) {
        setError("Available hours must be a valid non-negative number.");
        return;
      }
      availableHoursNum = n > 0 ? Math.floor(n) : null;
    }

    const payload = {
      full_name: form.full_name.trim(),
      phone: form.phone.trim(),
      skills: form.skills.trim(),
      consent: form.consent,
    };

    if (causeIdTrimmed) payload.cause_id = causeIdTrimmed;
    if (form.email.trim()) payload.email = form.email.trim();
    if (form.village.trim()) payload.village = form.village.trim();
    if (form.city.trim()) payload.city = form.city.trim();
    if (form.district.trim()) payload.district = form.district.trim();
    if (form.state.trim()) payload.state = form.state.trim();
    if (form.interests.trim()) payload.interests = form.interests.trim();
    if (form.availability_type.trim()) {
      payload.availability_type = form.availability_type.trim();
    }
    if (availableHoursNum != null) payload.available_hours = availableHoursNum;
    if (form.experience.trim()) payload.experience = form.experience.trim();

    setSubmitting(true);
    const result = await apiRequest(API_ENDPOINTS.CREATE_VOLUNTEER, {
      method: "POST",
      body: JSON.stringify(payload),
    });
    setSubmitting(false);

    if (!result.success) {
      setError(result.error || "Failed to submit volunteer information.");
      return;
    }

    setSuccess("Your volunteer details have been submitted successfully.");
    setForm(initialForm);
  };

  return (
    <div className="max-w-3xl mx-auto px-4 py-8">
      <div className="bg-white border border-gray-200 rounded-xl shadow-sm p-6">
        <h1 className="text-2xl font-bold text-[#3a0b2e] mb-2">Volunteer Form</h1>
        <p className="text-gray-600 mb-6">
          Share your skills and availability to support this campaign.
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
              <label className="block text-sm font-medium mb-2">Email <span className="text-red-500">*</span></label>
              <input
                name="email"
                value={form.email}
                onChange={handleChange}
                placeholder="Email"
                type="email"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                required
              />
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">Village <span className="text-red-500">*</span></label>
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
              <label className="block text-sm font-medium mb-2">City <span className="text-red-500">*</span></label>
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
              <label className="block text-sm font-medium mb-2">District <span className="text-red-500">*</span></label>
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
              <label className="block text-sm font-medium mb-2">State <span className="text-red-500">*</span></label>
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

          <label className="block text-sm font-medium mb-2">
            Skills <span className="text-red-500">*</span>
          </label>
          <textarea
            name="skills"
            value={form.skills}
            onChange={handleChange}
            placeholder="Describe relevant skills (e.g. teaching, logistics, first aid)"
            rows="3"
            className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
            required
          />

          <label className="block text-sm font-medium mb-2">Interests (optional)</label>
          <textarea
            name="interests"
            value={form.interests}
            onChange={handleChange}
            placeholder="Areas you are most interested in"
            rows="2"
            className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
          />

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Availability type
              </label>
              <select
                name="availability_type"
                value={form.availability_type}
                onChange={handleChange}
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] bg-white"
              >
                <option value="">Select (optional)</option>
                <option value="full-time">Full-time</option>
                <option value="part-time">Part-time</option>
                <option value="weekends">Weekends</option>
              </select>
            </div>
            <div className="flex flex-col">
              <label className="block text-sm font-medium mb-2">
                Available hours per week (optional)
              </label>
              <input
                name="available_hours"
                value={form.available_hours}
                onChange={handleChange}
                placeholder="e.g. 10"
                type="number"
                min="0"
                className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
              />
            </div>
          </div>

          <label className="block text-sm font-medium mb-2">
            Prior experience (optional)
          </label>
          <textarea
            name="experience"
            value={form.experience}
            onChange={handleChange}
            placeholder="Relevant volunteering or work experience"
            rows="3"
            className="w-full border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
          />

          <label className="flex items-center gap-2 text-sm text-gray-700">
            <input
              name="consent"
              checked={form.consent}
              onChange={handleChange}
              type="checkbox"
              required
            />
            I consent to share these details for volunteer coordination *
          </label>

          <button
            type="submit"
            disabled={submitting}
            className="bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-5 py-2 rounded-lg transition cursor-pointer disabled:opacity-60"
          >
            {submitting ? "Submitting..." : "Submit"}
          </button>
        </form>
      </div>
    </div>
  );
}
