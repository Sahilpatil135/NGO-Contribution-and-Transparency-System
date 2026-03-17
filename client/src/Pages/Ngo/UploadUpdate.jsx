import { useEffect, useState } from "react";
import { useNavigate, useParams, Link } from "react-router-dom";
import { apiRequest, API_ENDPOINTS, API_BASE_URL } from "@/config/api";

const UPDATE_TYPES = [
  { value: "Engagement", label: "Engagement (during fundraising)" },
  { value: "Execution", label: "Execution (after funds released)" },
];

export default function UploadUpdate() {
  const { causeID } = useParams();
  const navigate = useNavigate();
  const [cause, setCause] = useState(null);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");

  const [form, setForm] = useState({
    title: "",
    description: "",
    update_type: "Engagement",
    funding_percentage: "",
  });

  const [receiptUploadLoading, setReceiptUploadLoading] = useState(false);
  const [receipts, setReceipts] = useState([]);

  useEffect(() => {
    const fetchCause = async () => {
      const res = await apiRequest(`${API_ENDPOINTS.GET_CAUSES}/${causeID}`);
      if (res.success && res.data) {
        setCause(res.data);
      }
      setLoading(false);
    };
    fetchCause();
  }, [causeID]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleReceiptChange = async (e) => {
    const file = e.target.files?.[0];
    // const files = Array.from(e.target.files);
    if (!file) return;

    setError("");
    if (!file.type.startsWith("image/")) {
      setError("Receipt must be an image file.");
      return;
    }

    // Add if receipt accepts pdf format 
    // if (
    //   !file.type.startsWith("image/") &&
    //   file.type !== "application/pdf"
    // ) {
    //   setError("Only JPG, PNG, or PDF files are allowed.");
    //   return;
    // }

    const token = localStorage.getItem("authToken");
    if (!token) {
      setError("You must be logged in as an NGO to upload receipts.");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    setReceiptUploadLoading(true);
    try {
      const response = await fetch(API_ENDPOINTS.UPLOAD_UPDATE_RECEIPT, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formData,
      });

      if (!response.ok) {
        const text = await response.text();
        throw new Error(text || "Failed to upload receipt");
      }

      const data = await response.json();
      if (!data.url) {
        throw new Error("Upload did not return receipt URL");
      }

      setReceipts((prev) => [...prev, data.url]);
    } catch (err) {
      console.error(err);
      setError(err.message || "Failed to upload receipt");
    } finally {
      setReceiptUploadLoading(false);
      e.target.value = "";
    }
  };

  const handleRemoveReceipt = (indexToRemove) => {
    setReceipts((prev) => prev.filter((_, i) => i !== indexToRemove));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    const title = form.title.trim();
    const description = form.description.trim();
    if (!title || !description) {
      setError("Title and description are required.");
      return;
    }

    const payload = {
      title,
      description,
      update_type: form.update_type,
    };

    if (form.funding_percentage) {
      const fp = parseInt(form.funding_percentage, 10);
      if (!Number.isNaN(fp)) {
        payload.funding_percentage = fp;
      }
    }

    if (form.update_type === "Execution" && receipts.length > 0) {
      payload.receipt_urls = receipts;
    }

    setSubmitting(true);
    const result = await apiRequest(API_ENDPOINTS.CREATE_CAUSE_UPDATE(causeID), {
      method: "POST",
      body: JSON.stringify(payload),
    });
    setSubmitting(false);

    if (!result.success) {
      setError(result.error || "Failed to post update");
      return;
    }

    navigate(`/campaign/${causeID}`);
  };

  if (loading) {
    return (
      <div className="max-w-4xl mx-auto px-4 py-8">
        <p className="text-gray-600">Loading campaign...</p>
      </div>
    );
  }

  if (!cause) {
    return (
      <div className="max-w-4xl mx-auto px-4 py-8">
        <p className="text-red-500">Campaign not found.</p>
      </div>
    );
  }


  return (
    <div className="max-w-5xl mx-auto px-4 py-10">
      <h1 className="text-3xl md:text-4xl mb-6 font-bold text-[#3a0b2e]">
        {cause.title}
      </h1>
      <h2 className="text-xl font-semibold text-gray-800 mb-6">
        Post Campaign Update
      </h2>

      <div className="mb-6 text-sm text-gray-600">
        <p>
          Use <span className="font-semibold">Engagement</span> updates during
          fundraising to share stories, milestones and progress.
        </p>
        <p className="mt-1">
          Use <span className="font-semibold">Execution</span> updates after
          funds are released to attach receipts and real-world proof of work.
        </p>
      </div>

      {error && (
        <div className="mb-4 p-3 rounded-md bg-red-50 border border-red-200 text-sm text-red-700">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="space-y-6 bg-white p-6 rounded-xl shadow-sm border border-gray-200">
        <div>
          <label className="block text-sm font-medium mb-1">Update Type</label>
          <select
            name="update_type"
            value={form.update_type}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200]"
          >
            {UPDATE_TYPES.map((t) => (
              <option key={t.value} value={t.value}>
                {t.label}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label className="block text-sm font-medium mb-1">
            Title <span className="text-red-500">*</span>
          </label>
          <input
            type="text"
            name="title"
            value={form.title}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
            placeholder="e.g. Field visit with beneficiaries"
          />
        </div>

        <div>
          <label className="block text-sm font-medium mb-1">
            Description <span className="text-red-500">*</span>
          </label>
          <textarea
            name="description"
            value={form.description}
            onChange={handleChange}
            rows={5}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
            placeholder="Share what happened, who benefited and any key outcomes."
          />
        </div>

        {/* <div>
          <label className="block text-sm font-medium mb-1">
            Funding Percentage (optional)
          </label>
          <input
            type="number"
            name="funding_percentage"
            value={form.funding_percentage}
            onChange={handleChange}
            className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
            placeholder="e.g. 50 for 50% funded"
          />
        </div> */}

        {form.update_type === "Execution" && (
          <div>
            <label className="block text-sm font-medium mb-1">
              Receipts (Execution updates only)
            </label>
            {/* <p className="text-xs text-gray-500 mb-2">
            Attach financial or purchase receipts as image files in <span className="text-black">jpg, png and pdf format</span>. These will be
            visible to donors under the Updates tab.
          </p> */}
            <p className="text-xs text-gray-500">
              Attach financial or purchase receipts. Supported formats:
              <div className="flex gap-2 mb-1">
                <span className="px-2 py-1 text-xs font-medium bg-blue-100 text-blue-700 rounded">
                  JPG
                </span>
                <span className="px-2 py-1 text-xs font-medium bg-green-100 text-green-700 rounded">
                  PNG
                </span>
                <span className="px-2 py-1 text-xs font-medium bg-red-100 text-red-700 rounded">
                  PDF
                </span>
              </div>
            </p>
            <p className="text-xs text-gray-500">
              These will be visible to donors under the Updates tab.
            </p>
            <input
              type="file"
              accept="image/*"
              onChange={handleReceiptChange}
              disabled={form.update_type !== "Execution"}
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] cursor-pointer"
            />
            {receiptUploadLoading && (
              <p className="text-xs text-gray-500 mt-1">
                Uploading receipt...
              </p>
            )}
            {receipts.length > 0 && (
              <div className="mt-3 grid grid-cols-2 md:grid-cols-3 gap-3">
                {receipts.map((url, idx) => {
                  const src = url.startsWith("http")
                    ? url
                    : `${API_BASE_URL}${url}`;
                  return (
                    <div key={idx} className="relative group">
                      <img
                        // key={idx}
                        src={src}
                        alt={`Receipt ${idx + 1}`}
                        className="w-full h-24 object-cover rounded border"
                      />
                      {/* <span className="absolute bottom-1 left-1 bg-white text-[10px] px-1 rounded">
                        IMG
                      </span>
                      {/* Remove button */}
                      <button
                        type="button"
                        onClick={() => handleRemoveReceipt(idx)}
                        className="absolute top-1 right-1 bg-black/70 text-white text-xs rounded-full w-6 h-6 flex items-center justify-center opacity-0 group-hover:opacity-100 transition cursor-pointer"
                      >
                        ✕
                      </button>
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        )}

        {form.update_type === "Execution" && (
          <div className="pt-2 border-t border-dashed border-gray-200">
            <h3 className="text-sm font-semibold text-gray-800 mb-1">
              Proof of Work Images
            </h3>
            <p className="text-xs text-gray-500 mb-3">
              Capture geo-tagged images from the field to prove real-world
              execution. This opens the QR-based proof capture flow.
            </p>
            <Link to={`/uploadProof/${causeID}`} onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
              <button
                type="button"
                className="px-4 py-2 rounded-lg bg-[#3a0b2e] hover:bg-[#6d1f57] text-white text-sm font-semibold cursor-pointer"
              >
                Open Proof Capture
              </button>
            </Link>
          </div>
        )}

        <div className="flex items-center justify-between pt-4">
          <button
            type="button"
            className="px-4 py-2 rounded-lg border border-gray-300 text-gray-700 text-sm cursor-pointer"
            onClick={() => navigate(-1)}
          >
            Cancel
          </button>
          <button
            type="submit"
            disabled={submitting}
            className="px-6 py-2 rounded-lg bg-[#ff6200] hover:bg-[#e45a00] text-white text-sm font-semibold disabled:opacity-60 cursor-pointer"
          >
            {submitting ? "Posting..." : "Post Update"}
          </button>
        </div>
      </form>
    </div>
  );
}

