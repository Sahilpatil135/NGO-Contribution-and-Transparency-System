import React, { useState, useEffect } from "react";
import { API_ENDPOINTS, apiRequest } from "../config/api";

const CreateCampaign = () => {
  const [domains, setDomains] = useState([]);
  const [aidTypes, setAidTypes] = useState([]);
  const [formData, setFormData] = useState({
    title: "",
    description: "",
    domain_id: "",
    aid_type_id: "",
    goal_amount: "",
    deadline: "",
    execution_lat: "",
    execution_lng: "",
    execution_radius_meters: "",
    execution_start_time: "",
    execution_end_time: "",
    funding_status: "PENDING",
  });

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchOptions = async () => {
      const [domRes, aidRes] = await Promise.all([
        apiRequest(API_ENDPOINTS.GET_ALL_DOMAINS),
        apiRequest(API_ENDPOINTS.GET_ALL_AID_TYPES),
      ]);
      if (domRes.success && domRes.data) setDomains(Array.isArray(domRes.data) ? domRes.data : []);
      if (aidRes.success && aidRes.data) setAidTypes(Array.isArray(aidRes.data) ? aidRes.data : []);
    };
    fetchOptions();
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");
    if (!formData.title || !formData.domain_id || !formData.aid_type_id) {
      setError("Please fill Title, Domain and Aid Type.");
      return;
    }

    // Build payload; use RFC3339 for dates so backend can parse
    const payload = {
      title: formData.title.trim(),
      description: formData.description?.trim() || null,
      domain_id: formData.domain_id?.trim() || null,
      aid_type_id: formData.aid_type_id?.trim() || null,
      goal_amount: formData.goal_amount ? parseFloat(formData.goal_amount) : null,
      deadline: formData.deadline ? new Date(formData.deadline + "T00:00:00Z").toISOString() : null,
      execution_lat: formData.execution_lat ? parseFloat(formData.execution_lat) : null,
      execution_lng: formData.execution_lng ? parseFloat(formData.execution_lng) : null,
      execution_radius_meters: formData.execution_radius_meters ? parseInt(formData.execution_radius_meters, 10) : null,
      execution_start_time: formData.execution_start_time ? new Date(formData.execution_start_time).toISOString() : null,
      execution_end_time: formData.execution_end_time ? new Date(formData.execution_end_time).toISOString() : null,
      funding_status: formData.funding_status || null,
    };
    // Remove null, undefined, or empty string so backend never gets invalid UUID/date
    Object.keys(payload).forEach((k) => {
      const v = payload[k];
      if (v === null || v === undefined || v === "") delete payload[k];
    });

    setLoading(true);
    try {
      const result = await apiRequest(API_ENDPOINTS.CREATE_CAUSE, {
        method: "POST",
        body: JSON.stringify(payload),
      });
      if (result.success) {
        alert("Campaign created successfully!");
        setFormData({
          title: "",
          description: "",
          domain_id: "",
          aid_type_id: "",
          goal_amount: "",
          deadline: "",
          execution_lat: "",
          execution_lng: "",
          execution_radius_meters: "",
          execution_start_time: "",
          execution_end_time: "",
          funding_status: "PENDING",
        });
      } else {
        setError(result.error || "Error creating campaign");
      }
    } catch (err) {
      console.error(err);
      setError("Error connecting to server");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bg-gray-100 py-12">
      <div className="max-w-4xl mx-auto bg-white rounded-lg shadow-md p-10">
        <h1 className="text-3xl font-bold text-[#3a0b2e] mb-8 text-center">
          Create a New Campaign
        </h1>

        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg text-red-700">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-8">

          {/* --- Campaign Title --- */}
          <div>
            <label className="block text-sm font-medium mb-2">
              Campaign Title <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              name="title"
              value={formData.title}
              onChange={handleChange}
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              placeholder="e.g. Help Dr Amruta provide cancer treatment"
            />
          </div>

          {/* --- Description --- */}
          <div>
            <label className="block text-sm font-medium mb-2">
              Description <span className="text-red-500">*</span>
            </label>
            <textarea
              name="description"
              value={formData.description}
              onChange={handleChange}
              rows="5"
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              placeholder="Describe your campaign purpose, who it helps, and how donations will be used."
            ></textarea>
          </div>

          {/* --- Domain and Aid Type --- */}
          <div className="grid md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium mb-2">Domain <span className="text-red-500">*</span></label>
              <select
                name="domain_id"
                value={formData.domain_id}
                onChange={handleChange}
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              >
                <option value="">Select domain</option>
                {domains.map((d) => (
                  <option key={d.id} value={d.id}>{d.name}</option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium mb-2">Aid Type <span className="text-red-500">*</span></label>
              <select
                name="aid_type_id"
                value={formData.aid_type_id}
                onChange={handleChange}
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              >
                <option value="">Select type</option>
                {aidTypes.map((a) => (
                  <option key={a.id} value={a.id}>{a.name}</option>
                ))}
              </select>
            </div>
          </div>

          {/* --- Goal Amount & Deadline --- */}
          <div className="grid md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium mb-2">
                Goal Amount (₹)
              </label>
              <input
                type="number"
                name="goal_amount"
                value={formData.goal_amount}
                onChange={handleChange}
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none"
                placeholder="e.g. 500000"
              />
            </div>

            <div>
              <label className="block text-sm font-medium mb-2">Deadline</label>
              <input
                type="date"
                name="deadline"
                value={formData.deadline}
                onChange={handleChange}
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              />
            </div>
          </div>

          {/* Cover Image Pending Implementation */}

          {/* --- Execution / Proof of work approach pending --- */}
          {/* Deadline will tell till when funds needed for the campaign */}
          {/* Execution window will tell when the campaign will be executed. We can put schedule (like 3 days drive) in place of 
          execution time.*/}
          {/* Also execution time can be updated after funds are fullfilled. */}

          <section>
            <h2 className="text-xl font-semibold text-[#ff6200] mb-4">
              Execution Window & Location (for proof of work)
            </h2>
            <p className="text-sm text-gray-600 mb-4">
              Optional. Set when and where this cause will be executed so proof captures can be validated.
            </p>
            <div className="grid md:grid-cols-2 gap-6">
              <div>
                <label className="block text-sm font-medium mb-2">Execution start time</label>
                <input
                  type="datetime-local"
                  name="execution_start_time"
                  value={formData.execution_start_time}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">Execution end time</label>
                <input
                  type="datetime-local"
                  name="execution_end_time"
                  value={formData.execution_end_time}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">Execution latitude</label>
                <input
                  type="number"
                  step="any"
                  name="execution_lat"
                  value={formData.execution_lat}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="e.g. 19.0760"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">Execution longitude</label>
                <input
                  type="number"
                  step="any"
                  name="execution_lng"
                  value={formData.execution_lng}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="e.g. 72.8777"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">Radius (meters)</label>
                <input
                  type="number"
                  name="execution_radius_meters"
                  value={formData.execution_radius_meters}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                  placeholder="200"
                />
              </div>
              <div>
                <label className="block text-sm font-medium mb-2">Funding status</label>
                <select
                  name="funding_status"
                  value={formData.funding_status}
                  onChange={handleChange}
                  className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
                >
                  <option value="PENDING">PENDING</option>
                  <option value="FUNDED">FUNDED</option>
                  <option value="COMPLETED">COMPLETED</option>
                </select>
              </div>
            </div>
          </section>

          {/* --- Submit --- */}
          <div className="text-center pt-4">
            <button
              type="submit"
              disabled={loading}
              className="bg-[#ff6200] text-white font-semibold px-8 py-3 rounded-lg hover:bg-[#e45a00] transition disabled:opacity-50 cursor-pointer"
            >
              {loading ? "Submitting..." : "Create Campaign"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateCampaign;
