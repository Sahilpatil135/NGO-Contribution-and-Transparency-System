import React, { useState, useEffect } from "react";
import { API_ENDPOINTS, apiRequest } from "../config/api";

const PRODUCT_AID_TYPE_NAMES = [
  "Goods & Resources",
  "Educational Support",
  "Medical Assistance",
  "Environmental Support",
  "Disaster Relief Assistance",
];

const CreateCampaign = () => {
  const [domains, setDomains] = useState([]);
  const [aidTypes, setAidTypes] = useState([]);
  const [formData, setFormData] = useState({
    // step1 fields
    title: "",
    description: "",
    domain_id: "",
    aid_type_id: "",
    goal_amount: "",
    deadline: "",
    cover_image_url: "",
    // step2 fields
    problem_statement: "",
    execution_plan: "",
    beneficiaries_count: "",
    execution_location: "",
    execution_start_time: "",
    execution_end_time: "",
    impact_goal: "",
    execution_lat: "",
    execution_lng: "",
    execution_radius_meters: "",
    funding_status: "Active",
  });

  const [products, setProducts] = useState([
    {
      name: "",
      description: "",
      price_per_unit: "",
      quantity_needed: "",
      image_url: "",
    },
  ]);

  const [currentStep, setCurrentStep] = useState(1);
  const [loading, setLoading] = useState(false);
  const [coverUploading, setCoverUploading] = useState(false);
  const [error, setError] = useState("");
  const [productUploadLoading, setProductUploadLoading] = useState({});

  useEffect(() => {
    const fetchOptions = async () => {
      const [domRes, aidRes] = await Promise.all([
        apiRequest(API_ENDPOINTS.GET_ALL_DOMAINS),
        apiRequest(API_ENDPOINTS.GET_ALL_AID_TYPES),
      ]);
      if (domRes.success && domRes.data) {
        setDomains(Array.isArray(domRes.data) ? domRes.data : []);
      }
      if (aidRes.success && aidRes.data) {
        setAidTypes(Array.isArray(aidRes.data) ? aidRes.data : []);
      }
    };
    fetchOptions();
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleProductChange = (index, field, value) => {
    setProducts((prev) =>
      prev.map((p, i) => (i === index ? { ...p, [field]: value } : p))
    );
  };

  const addProductRow = () => {
    setProducts((prev) => [
      ...prev,
      {
        name: "",
        description: "",
        price_per_unit: "",
        quantity_needed: "",
        image_url: "",
      },
    ]);
  };

  const removeProductRow = (index) => {
    setProducts((prev) => prev.filter((_, i) => i !== index));
  };

  const handleCoverChange = async (e) => {
    const file = e.target.files?.[0];
    if (!file) return;

    setError("");

    if (!file.type.startsWith("image/")) {
      setError("Cover image must be an image file (jpg, png, webp, etc).");
      return;
    }

    const token = localStorage.getItem("authToken");
    if (!token) {
      setError("You must be logged in as an NGO to upload a cover image.");
      return;
    }

    const formDataUpload = new FormData();
    formDataUpload.append("file", file);

    setCoverUploading(true);
    try {
      const response = await fetch(API_ENDPOINTS.UPLOAD_CAUSE_COVER, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formDataUpload,
      });

      if (!response.ok) {
        const text = await response.text();
        throw new Error(text || "Failed to upload cover image");
      }

      const data = await response.json();
      if (!data.url) {
        throw new Error("Upload did not return image URL");
      }

      setFormData((prev) => ({
        ...prev,
        cover_image_url: data.url,
      }));
    } catch (err) {
      console.error(err);
      setError(err.message || "Failed to upload cover image");
    } finally {
      setCoverUploading(false);
    }
  };

  const handleProductImageChange = async (index, file) => {
    if (!file) return;

    setError("");

    if (!file.type.startsWith("image/")) {
      setError("Product image must be an image file (jpg, png, webp, etc).");
      return;
    }

    const token = localStorage.getItem("authToken");
    if (!token) {
      setError("You must be logged in as an NGO to upload a product image.");
      return;
    }

    const formDataUpload = new FormData();
    formDataUpload.append("file", file);

    setProductUploadLoading((prev) => ({ ...prev, [index]: true }));
    try {
      const response = await fetch(API_ENDPOINTS.UPLOAD_PRODUCT_IMAGE, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formDataUpload,
      });

      if (!response.ok) {
        const text = await response.text();
        throw new Error(text || "Failed to upload product image");
      }

      const data = await response.json();
      if (!data.url) {
        throw new Error("Upload did not return image URL");
      }

      setProducts((prev) =>
        prev.map((p, i) =>
          i === index ? { ...p, image_url: data.url } : p
        )
      );
    } catch (err) {
      console.error(err);
      setError(err.message || "Failed to upload product image");
    } finally {
      setProductUploadLoading((prev) => ({ ...prev, [index]: false }));
    }
  };

  const selectedAidType = aidTypes.find(
    (a) => a.id === formData.aid_type_id
  );

  const requiresProducts = selectedAidType
    ? PRODUCT_AID_TYPE_NAMES.includes(selectedAidType.name)
    : false;

  const validateStep1 = () => {
    const title = formData.title.trim();
    const desc = formData.description.trim();
    const goal = parseFloat(formData.goal_amount || "0");
    const deadlineDate = formData.deadline ? new Date(formData.deadline) : null;
    const today = new Date();

    if (!title || title.length < 10 || title.length > 120) {
      return "Title must be between 10 and 120 characters.";
    }
    if (!formData.domain_id || !formData.aid_type_id) {
      return "Domain and Aid Type are required.";
    }
    if (!desc || desc.length < 100) {
      return "Full description must be at least 100 characters.";
    }
    if (isNaN(goal) || goal < 1000) {
      return "Goal amount must be at least ₹1000.";
    }
    if (!deadlineDate) {
      return "Deadline is required.";
    }
    const minDeadline = new Date(
      today.getFullYear(),
      today.getMonth(),
      today.getDate() + 7
    );
    if (deadlineDate < minDeadline) {
      return "Deadline must be at least 7 days from today.";
    }
    if (!formData.cover_image_url) {
      return "Cover image is required.";
    }
    return "";
  };

  const validateStep2 = () => {
    const problem = formData.problem_statement.trim();
    const plan = formData.execution_plan.trim();
    const impact = formData.impact_goal.trim();
    const ben = parseInt(formData.beneficiaries_count || "0", 10);

    if (!problem || problem.length < 100) {
      return "Problem statement must be at least 100 characters.";
    }
    if (!plan || plan.length < 150) {
      return "Execution plan must be at least 150 characters.";
    }
    if (isNaN(ben) || ben <= 0) {
      return "Beneficiaries count must be a positive integer.";
    }
    if (!formData.execution_location.trim()) {
      return "Execution location is required.";
    }
    if (!formData.execution_start_time || !formData.execution_end_time) {
      return "Execution start and end dates are required.";
    }
    const start = new Date(formData.execution_start_time);
    const end = new Date(formData.execution_end_time);
    if (!(end > start)) {
      return "Execution end date must be after start date.";
    }
    if (!impact || impact.length < 50) {
      return "Impact goal must be at least 50 characters.";
    }
    return "";
  };

  const validateProducts = () => {
    if (!requiresProducts) {
      return "";
    }
    if (!products.length) {
      return "At least one product is required for this aid type.";
    }

    let total = 0;
    for (const p of products) {
      if (
        !p.name.trim() ||
        !p.description.trim() ||
        !p.price_per_unit ||
        !p.quantity_needed
      ) {
        return "All product fields are required.";
      }
      const price = parseFloat(p.price_per_unit);
      const qty = parseInt(p.quantity_needed, 10);
      if (isNaN(price) || price <= 0 || isNaN(qty) || qty <= 0) {
        return "Product price and quantity must be greater than 0.";
      }
      total += price * qty;
    }

    const goal = parseFloat(formData.goal_amount || "0");
    if (Math.round(total) !== Math.round(goal)) {
      return "Sum of all product totals must equal the goal amount.";
    }

    return "";
  };

  const handleNext = () => {
    setError("");
    if (currentStep === 1) {
      const errMsg = validateStep1();
      if (errMsg) {
        setError(errMsg);
        return;
      }
      setCurrentStep(2);
    } else if (currentStep === 2) {
      const errMsg = validateStep2();
      if (errMsg) {
        setError(errMsg);
        return;
      }
      if (requiresProducts) {
        setCurrentStep(3);
      } else {
        setCurrentStep(4);
      }
    } else if (currentStep === 3) {
      const errMsg = validateProducts();
      if (errMsg) {
        setError(errMsg);
        return;
      }
      setCurrentStep(4);
    }
  };

  const handleBack = () => {
    setError("");
    if (currentStep === 4 && !requiresProducts) {
      setCurrentStep(2);
    } else {
      setCurrentStep((prev) => Math.max(1, prev - 1));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    const step1Error = validateStep1();
    if (step1Error) {
      setError(step1Error);
      setCurrentStep(1);
      return;
    }
    const step2Error = validateStep2();
    if (step2Error) {
      setError(step2Error);
      setCurrentStep(2);
      return;
    }
    const step3Error = validateProducts();
    if (step3Error) {
      setError(step3Error);
      if (requiresProducts) {
        setCurrentStep(3);
      }
      return;
    }

    const payload = {
      title: formData.title.trim(),
      description: formData.description.trim(),
      domain_id: formData.domain_id?.trim() || null,
      aid_type_id: formData.aid_type_id?.trim() || null,
      goal_amount: formData.goal_amount
        ? parseFloat(formData.goal_amount)
        : null,
      deadline: formData.deadline
        ? new Date(formData.deadline + "T00:00:00Z").toISOString()
        : null,
      cover_image_url: formData.cover_image_url || null,
      execution_lat: formData.execution_lat
        ? parseFloat(formData.execution_lat)
        : null,
      execution_lng: formData.execution_lng
        ? parseFloat(formData.execution_lng)
        : null,
      execution_radius_meters: formData.execution_radius_meters
        ? parseInt(formData.execution_radius_meters, 10)
        : null,
      execution_start_time: formData.execution_start_time
        ? new Date(formData.execution_start_time).toISOString()
        : null,
      execution_end_time: formData.execution_end_time
        ? new Date(formData.execution_end_time).toISOString()
        : null,
      funding_status: "Active",
      beneficiaries_count: formData.beneficiaries_count
        ? parseInt(formData.beneficiaries_count, 10)
        : null,
      execution_location: formData.execution_location.trim() || null,
      impact_goal: formData.impact_goal.trim() || null,
      problem_statement: formData.problem_statement.trim() || null,
      execution_plan: formData.execution_plan.trim() || null,
    };

    if (requiresProducts) {
      payload.products = products.map((p) => ({
        name: p.name.trim(),
        description: p.description.trim(),
        price_per_unit: parseFloat(p.price_per_unit),
        quantity_needed: parseInt(p.quantity_needed, 10),
        image_url: p.image_url?.trim() || undefined,
      }));
    }

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
          cover_image_url: "",
          problem_statement: "",
          execution_plan: "",
          beneficiaries_count: "",
          execution_location: "",
          execution_start_time: "",
          execution_end_time: "",
          impact_goal: "",
          execution_lat: "",
          execution_lng: "",
          execution_radius_meters: "",
          funding_status: "Active",
        });
        setProducts([
          {
            name: "",
            description: "",
            price_per_unit: "",
            quantity_needed: "",
            image_url: "",
          },
        ]);
        setCurrentStep(1);
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

  const renderStep1 = () => (
    <div className="space-y-6">
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
          placeholder="e.g. Support post-surgery care for elderly patients"
        />
        <p className="text-xs text-gray-500 mt-1">
          10–120 characters, clear and specific.
        </p>
      </div>

      <div>
        <label className="block text-sm font-medium mb-2">
          Full Description <span className="text-red-500">*</span>
        </label>
        <textarea
          name="description"
          value={formData.description}
          onChange={handleChange}
          rows="5"
          className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          placeholder="Describe your campaign purpose, who it helps, and how donations will be used (min 100 characters)."
        ></textarea>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium mb-2">
            Domain <span className="text-red-500">*</span>
          </label>
          <select
            name="domain_id"
            value={formData.domain_id}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          >
            <option value="">Select domain</option>
            {domains.map((d) => (
              <option key={d.id} value={d.id}>
                {d.name}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">
            Aid Type <span className="text-red-500">*</span>
          </label>
          <select
            name="aid_type_id"
            value={formData.aid_type_id}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          >
            <option value="">Select type</option>
            {aidTypes.map((a) => (
              <option key={a.id} value={a.id}>
                {a.name}
              </option>
            ))}
          </select>
        </div>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium mb-2">
            Goal Amount (₹) <span className="text-red-500">*</span>
          </label>
          <input
            type="number"
            name="goal_amount"
            value={formData.goal_amount}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none"
            placeholder="e.g. 500000"
          />
          <p className="text-xs text-gray-500 mt-1">
            Must be at least ₹1000.
          </p>
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">
            Deadline <span className="text-red-500">*</span>
          </label>
          <input
            type="date"
            name="deadline"
            value={formData.deadline}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          />
          <p className="text-xs text-gray-500 mt-1">
            Must be at least 7 days from today.
          </p>
        </div>
      </div>

      <div>
        <label className="block text-sm font-medium mb-2">
          Cover Image <span className="text-red-500">*</span>
        </label>
        <input
          type="file"
          accept="image/*"
          onChange={handleCoverChange}
          className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
        />
        {coverUploading && (
          <p className="text-xs text-gray-500 mt-1">
            Uploading cover image...
          </p>
        )}
        {formData.cover_image_url && !coverUploading && (
          <p className="text-xs text-green-600 mt-1">Cover image uploaded.</p>
        )}
      </div>
    </div>
  );

  const renderStep2 = () => (
    <div className="space-y-6">
      <div>
        <label className="block text-sm font-medium mb-2">
          Problem Statement <span className="text-red-500">*</span>
        </label>
        <textarea
          name="problem_statement"
          value={formData.problem_statement}
          onChange={handleChange}
          rows="4"
          className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          placeholder="Explain the core problem this campaign is addressing (min 100 characters)."
        ></textarea>
      </div>

      <div>
        <label className="block text-sm font-medium mb-2">
          Execution Plan <span className="text-red-500">*</span>
        </label>
        <textarea
          name="execution_plan"
          value={formData.execution_plan}
          onChange={handleChange}
          rows="5"
          className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          placeholder="Describe how you will execute this campaign (steps, partners, timelines) (min 150 characters)."
        ></textarea>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium mb-2">
            Beneficiaries Count <span className="text-red-500">*</span>
          </label>
          <input
            type="number"
            name="beneficiaries_count"
            value={formData.beneficiaries_count}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
            placeholder="e.g. 150"
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">
            Execution Location <span className="text-red-500">*</span>
          </label>
          <input
            type="text"
            name="execution_location"
            value={formData.execution_location}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
            placeholder="City / Area / Site"
          />
        </div>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <div>
          <label className="block text-sm font-medium mb-2">
            Execution Start Date <span className="text-red-500">*</span>
          </label>
          <input
            type="date"
            name="execution_start_time"
            value={formData.execution_start_time}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-2">
            Execution End Date <span className="text-red-500">*</span>
          </label>
          <input
            type="date"
            name="execution_end_time"
            value={formData.execution_end_time}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          />
        </div>
      </div>

      <div>
        <label className="block text-sm font-medium mb-2">
          Impact Goal <span className="text-red-500">*</span>
        </label>
        <textarea
          name="impact_goal"
          value={formData.impact_goal}
          onChange={handleChange}
          rows="3"
          className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          placeholder="What measurable impact do you aim to achieve? (min 50 characters)."
        ></textarea>
      </div>

      <section>
        <h2 className="text-lg font-semibold text-[#ff6200] mb-3">
          Execution Coordinates (optional)
        </h2>
        <p className="text-sm text-gray-600 mb-4">
          Helps validate proof-of-work captures around the execution location.
        </p>
        <div className="grid md:grid-cols-3 gap-6">
          <div>
            <label className="block text-sm font-medium mb-2">
              Latitude
            </label>
            <input
              type="number"
              step="any"
              name="execution_lat"
              value={formData.execution_lat}
              onChange={handleChange}
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
            />
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">
              Longitude
            </label>
            <input
              type="number"
              step="any"
              name="execution_lng"
              value={formData.execution_lng}
              onChange={handleChange}
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
            />
          </div>
          <div>
            <label className="block text-sm font-medium mb-2">
              Radius (meters)
            </label>
            <input
              type="number"
              name="execution_radius_meters"
              value={formData.execution_radius_meters}
              onChange={handleChange}
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
            />
          </div>
        </div>
      </section>
    </div>
  );

  const renderStep3 = () => (
    <div className="space-y-6">
      <p className="text-sm text-gray-700">
        This aid type requires structured products. Define exactly what donors
        are sponsoring. The sum of all products must equal the campaign goal.
      </p>

      {products.map((p, idx) => (
        <div
          key={idx}
          className="border rounded-lg p-4 space-y-4 bg-gray-50 relative"
        >
          {products.length > 1 && (
            <button
              type="button"
              onClick={() => removeProductRow(idx)}
              className="absolute top-2 right-2 text-xs text-red-500 underline cursor-pointer"
            >
              Remove
            </button>
          )}
          <div className="grid md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium mb-1">
                Product Name <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                value={p.name}
                onChange={(e) =>
                  handleProductChange(idx, "name", e.target.value)
                }
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">
                Price Per Unit (₹) <span className="text-red-500">*</span>
              </label>
              <input
                type="number"
                value={p.price_per_unit}
                onChange={(e) =>
                  handleProductChange(idx, "price_per_unit", e.target.value)
                }
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              />
            </div>
          </div>

          <div className="grid md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium mb-1">
                Quantity Needed <span className="text-red-500">*</span>
              </label>
              <input
                type="number"
                value={p.quantity_needed}
                onChange={(e) =>
                  handleProductChange(idx, "quantity_needed", e.target.value)
                }
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">
                Product Image (optional)
              </label>
              <input
                type="file"
                accept="image/*"
                onChange={(e) =>
                  handleProductImageChange(idx, e.target.files?.[0] || null)
                }
                className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
              />
              {productUploadLoading[idx] && (
                <p className="text-xs text-gray-500 mt-1">
                  Uploading product image...
                </p>
              )}
              {p.image_url && !productUploadLoading[idx] && (
                <p className="text-xs text-green-600 mt-1">
                  Product image uploaded.
                </p>
              )}
            </div>
          </div>

          <div>
            <label className="block text-sm font-medium mb-1">
              Product Description <span className="text-red-500">*</span>
            </label>
            <textarea
              rows="3"
              value={p.description}
              onChange={(e) =>
                handleProductChange(idx, "description", e.target.value)
              }
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
            ></textarea>
          </div>

          <p className="text-xs text-gray-600">
            Total for this product:{" "}
            <span className="font-semibold">
              ₹
              {(
                (parseFloat(p.price_per_unit || "0") || 0) *
                (parseInt(p.quantity_needed || "0", 10) || 0)
              ).toLocaleString()}
            </span>
          </p>
        </div>
      ))}

      <button
        type="button"
        onClick={addProductRow}
        className="text-sm text-[#ff6200] font-semibold underline cursor-pointer"
      >
        + Add another product
      </button>
    </div>
  );

  const renderStep4 = () => {
    const goal = parseFloat(formData.goal_amount || "0") || 0;
    const collected = 0;
    const percentage = goal > 0 ? Math.min((collected / goal) * 100, 100) : 0;

    return (
      <div className="space-y-6">
        <p className="text-sm text-gray-700">
          Review all details before publishing. Once donations begin, goal
          amount and deadline should not be changed.
        </p>

        <div className="border rounded-lg p-4 space-y-3 bg-gray-50">
          <h3 className="font-semibold text-[#3a0b2e]">Basic Details</h3>
          <p className="text-sm">
            <span className="font-semibold">Title:</span> {formData.title}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Domain:</span>{" "}
            {domains.find((d) => d.id === formData.domain_id)?.name || "-"}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Aid Type:</span>{" "}
            {selectedAidType?.name || "-"}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Goal Amount:</span> ₹
            {goal.toLocaleString()}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Deadline:</span>{" "}
            {formData.deadline || "-"}
          </p>
        </div>

        <div className="border rounded-lg p-4 space-y-3 bg-gray-50">
          <h3 className="font-semibold text-[#3a0b2e]">Project Details</h3>
          <p className="text-sm">
            <span className="font-semibold">Beneficiaries:</span>{" "}
            {formData.beneficiaries_count || "-"}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Location:</span>{" "}
            {formData.execution_location || "-"}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Execution Window:</span>{" "}
            {formData.execution_start_time || "-"} to{" "}
            {formData.execution_end_time || "-"}
          </p>
          <p className="text-sm">
            <span className="font-semibold">Impact Goal:</span>{" "}
            {formData.impact_goal || "-"}
          </p>
        </div>

        {requiresProducts && (
          <div className="border rounded-lg p-4 space-y-3 bg-gray-50">
            <h3 className="font-semibold text-[#3a0b2e]">Products Summary</h3>
            {products.map((p, idx) => (
              <p key={idx} className="text-sm">
                <span className="font-semibold">{p.name || "Unnamed"}:</span>{" "}
                {p.quantity_needed || 0} × ₹{p.price_per_unit || 0} = ₹
                {(
                  (parseFloat(p.price_per_unit || "0") || 0) *
                  (parseInt(p.quantity_needed || "0", 10) || 0)
                ).toLocaleString()}
              </p>
            ))}
          </div>
        )}

        <div className="border rounded-lg p-4 space-y-3 bg-gray-50">
          <h3 className="font-semibold text-[#3a0b2e]">Funding Preview</h3>
          <p className="text-sm">
            <span className="font-semibold">Funding Status:</span> Active
          </p>
          <div className="w-full bg-gray-200 rounded-full h-2 overflow-hidden">
            <div
              className="h-full bg-[#ff6200]"
              style={{ width: `${percentage}%` }}
            ></div>
          </div>
        </div>
      </div>
    );
  };

  return (
    <div className="bg-gray-100 py-12">
      <div className="max-w-5xl mx-auto bg-white rounded-lg shadow-md p-10">
        <h1 className="text-3xl font-bold text-[#3a0b2e] mb-6 text-center">
          Create a New Campaign
        </h1>

        <div className="flex justify-center mb-8">
          <div className="flex items-center space-x-4 text-sm font-medium">
            <span
              className={
                currentStep === 1 ? "text-[#ff6200]" : "text-gray-400"
              }
            >
              1. Basic Details
            </span>
            <span>›</span>
            <span
              className={
                currentStep === 2 ? "text-[#ff6200]" : "text-gray-400"
              }
            >
              2. Project
            </span>
            {requiresProducts && (
              <>
                <span>›</span>
                <span
                  className={
                    currentStep === 3 ? "text-[#ff6200]" : "text-gray-400"
                  }
                >
                  3. Products
                </span>
              </>
            )}
            <span>›</span>
            <span
              className={
                currentStep === 4 ? "text-[#ff6200]" : "text-gray-400"
              }
            >
              4. Review & Publish
            </span>
          </div>
        </div>

        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-200 rounded-lg text-red-700">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-8">
          {currentStep === 1 && renderStep1()}
          {currentStep === 2 && renderStep2()}
          {currentStep === 3 && requiresProducts && renderStep3()}
          {currentStep === 4 && renderStep4()}

          <div className="flex items-center justify-between pt-6">
            <button
              type="button"
              onClick={handleBack}
              disabled={currentStep === 1}
              className="px-4 py-2 rounded-lg border border-gray-300 text-gray-700 text-sm disabled:opacity-40 cursor-pointer"
            >
              Back
            </button>

            {currentStep < 4 && (
              <button
                type="button"
                onClick={handleNext}
                className="bg-[#ff6200] text-white font-semibold px-6 py-2 rounded-lg hover:bg-[#e45a00] transition text-sm cursor-pointer"
              >
                Next
              </button>
            )}

            {currentStep === 4 && (
              <button
                type="submit"
                disabled={loading}
                className="bg-[#ff6200] text-white font-semibold px-8 py-2 rounded-lg hover:bg-[#e45a00] transition disabled:opacity-50 cursor-pointer"
              >
                {loading ? "Submitting..." : "Publish Campaign"}
              </button>
            )}
          </div>
        </form>
      </div>
    </div>
  );
};

export default CreateCampaign;
