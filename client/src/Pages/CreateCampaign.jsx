import React, { useState } from "react";

const CreateCampaign = () => {
  const [formData, setFormData] = useState({
    organization_id: "",
    title: "",
    description: "",
    domain_id: "",
    aid_type_id: "",
    goal_amount: "",
    deadline: "",
  });

  const [coverImage, setCoverImage] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleFileChange = (e) => {
    setCoverImage(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!formData.title || !formData.organization_id) {
      alert("Please fill all required fields");
      return;
    }

    const data = new FormData();
    Object.entries(formData).forEach(([key, value]) => data.append(key, value));
    if (coverImage) data.append("cover_image", coverImage);

    setLoading(true);
    try {
      const res = await fetch("http://localhost:8080/api/cause/create", {
        method: "POST",
        body: data,
      });
      if (res.ok) {
        alert("✅ Campaign created successfully!");
        setFormData({
          organization_id: "",
          title: "",
          description: "",
          domain_id: "",
          aid_type_id: "",
          goal_amount: "",
          deadline: "",
        });
        setCoverImage(null);
      } else {
        alert("❌ Error creating campaign");
      }
    } catch (error) {
      console.error(error);
      alert("Error connecting to server");
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

        <form onSubmit={handleSubmit} className="space-y-8">

          {/* --- Organization ID --- */}
          <div>
            <label className="block text-sm font-medium mb-2">
              Organization ID <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              name="organization_id"
              value={formData.organization_id}
              onChange={handleChange}
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              placeholder="Enter your organization ID"
            />
          </div>

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
              <label className="block text-sm font-medium mb-2">Domain</label>
              <select
                name="domain_id"
                value={formData.domain_id}
                onChange={handleChange}
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              >
                <option value="">Select domain</option>
                <option value="1">Health</option>
                <option value="2">Education</option>
                <option value="3">Environment</option>
                <option value="4">Disaster Relief</option>
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium mb-2">Aid Type</label>
              <select
                name="aid_type_id"
                value={formData.aid_type_id}
                onChange={handleChange}
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              >
                <option value="">Select type</option>
                <option value="1">Medical Aid</option>
                <option value="2">Food Aid</option>
                <option value="3">Shelter Aid</option>
                <option value="4">Educational Aid</option>
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
                placeholder="500000"
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

          {/* --- Cover Image --- */}
          <div>
            <label className="block text-sm font-medium mb-2">Cover Image</label>
            <input
              type="file"
              accept="image/*"
              onChange={handleFileChange}
              className="block w-full text-sm text-gray-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-white file:bg-[#ff6200] hover:file:bg-[#e45a00] file:cursor-pointer cursor-pointer"
            />
          </div>

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
