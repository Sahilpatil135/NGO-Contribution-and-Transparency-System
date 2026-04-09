import { useEffect, useRef, useState } from "react";
import { useNavigate, useParams, Link, useLocation } from "react-router-dom";
import { apiRequest, API_ENDPOINTS, API_BASE_URL } from "@/config/api";

const UPDATE_TYPES = [
  { value: "Engagement", label: "Engagement (during fundraising)" },
  { value: "Execution", label: "Execution (after funds released)" },
];

export default function UploadUpdate() {
  const { causeID } = useParams();
  const navigate = useNavigate();
  const location = useLocation();
  const [cause, setCause] = useState(null);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState("");

  const [form, setForm] = useState({
    title: "",
    description: "",
    update_type: "Engagement",
    funding_percentage: "",
    claimed_amount: "",
  });

  const proofSessionIdFromQuery = new URLSearchParams(location.search).get(
    "proofSessionId"
  );
  const [proofs, setProofs] = useState([]);
  const [proofsLoading, setProofsLoading] = useState(false);
  const [proofsError, setProofsError] = useState("");

  const [receiptUploadLoading, setReceiptUploadLoading] = useState(false);
  // new {
  // Each receipt has its own verification job started on upload.
  // { url, receiptJobId, status, receiptScore, errorMessage }
  // }
  const [receipts, setReceipts] = useState([]);
  // new {
  const pollingIntervalsRef = useRef({});
  const draftKey = `uploadUpdateDraft:${causeID}`;
  const [draftRestored, setDraftRestored] = useState(false);
  const saveDraft = (nextForm, nextReceipts) => {
    try {
      localStorage.setItem(
        draftKey,
        JSON.stringify({
          form: nextForm,
          receipts: nextReceipts,
        })
      );
    } catch {
      // Ignore localStorage write errors (e.g. privacy mode).
    }
  };

  const claimedAmountNumber = Number(String(form.claimed_amount ?? "").trim());
  const claimedAmountValid =
    Number.isFinite(claimedAmountNumber) && claimedAmountNumber > 0;

  const isReceiptVerificationInProgress =
    form.update_type === "Execution" &&
    receipts.length > 0 &&
    receipts.some((r) => {
      const s = String(r.status ?? "").toLowerCase();
      return s === "" || s === "pending" || s === "processing" || s === "error";
    });
  // }

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

  // new {
  // Restore any saved draft if the user navigated away to proof capture.
  useEffect(() => {
    try {
      const raw = localStorage.getItem(draftKey);
      if (!raw) return;
      const parsed = JSON.parse(raw);
      if (parsed?.form) {
        setForm((prev) => ({ ...prev, ...parsed.form }));
      }
      if (Array.isArray(parsed?.receipts)) {
        const normalized = parsed.receipts.map((r) => {
          if (typeof r === "string") {
            return {
              url: r,
              receiptJobId: null,
              status: "pending",
              receiptScore: null,
              errorMessage: null,
            };
          }
          return r;
        });
        setReceipts(normalized);
      }
    } catch (err) {
      // Ignore draft restore errors.
    }
    setDraftRestored(true);
  }, [draftKey]);
  // }
  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => {
      const next = { ...prev, [name]: value };
      let nextReceipts = receipts;
      // Reset execution-only fields when switching away from Execution
      if (name === "update_type" && value !== "Execution") {
        next.claimed_amount = "";
        // new {
        for (const intervalId of Object.values(pollingIntervalsRef.current)) {
          clearInterval(intervalId);
        }
        pollingIntervalsRef.current = {};
        // }
        setReceipts([]);
        nextReceipts = [];
      }
      // Save immediately so draft exists even if we navigate away quickly.
      saveDraft(next, nextReceipts);
      return next;
    });
  };

  // new {
  // Persist draft so we can return from proof capture without losing input.
  useEffect(() => {
    if (!draftRestored) return;
    try {
      localStorage.setItem(
        draftKey,
        JSON.stringify({
          form,
          receipts,
        })
      );
    } catch {
      // Ignore localStorage write errors (e.g. privacy mode).
    }
  }, [draftKey, form, receipts, draftRestored]);

  useEffect(() => {
    // Cleanup polling intervals when leaving this page.
    return () => {
      for (const intervalId of Object.values(pollingIntervalsRef.current)) {
        clearInterval(intervalId);
      }
      pollingIntervalsRef.current = {};
    };
  }, []);

  // Backend-first: fetch proof images using the proof session id returned
  // from UploadProof back to this page.
  useEffect(() => {
    const sessionId = proofSessionIdFromQuery;
    if (!sessionId) {
      setProofs([]);
      return;
    }

    const fetchProofImages = async () => {
      setProofsLoading(true);
      setProofsError("");
      try {
        const res = await apiRequest(
          API_ENDPOINTS.GET_PROOF_IMAGES_BY_SESSION(sessionId)
        );
        if (res.success) {
          setProofs(Array.isArray(res.data) ? res.data : []);
        } else {
          setProofsError(res.error || "Failed to load proofs");
          setProofs([]);
        }
      } catch (e) {
        setProofsError(e?.message || "Failed to load proofs");
        setProofs([]);
      } finally {
        setProofsLoading(false);
      }
    };

    fetchProofImages();
  }, [proofSessionIdFromQuery]);

  const getProofStatusStyle = (status) => {
    const s = String(status ?? "").toLowerCase();
    if (s === "verified") return "bg-green-100 text-green-700";
    if (s === "review") return "bg-amber-100 text-amber-700";
    if (s === "rejected") return "bg-red-100 text-red-700";
    return "bg-gray-100 text-gray-700";
  };

  const startPollingReceipt = (receiptJobId) => {
    if (!receiptJobId) return;
    if (pollingIntervalsRef.current[receiptJobId]) return;

    let attempts = 0;
    const maxAttempts = 60; // ~2-3 min hard cap

    pollingIntervalsRef.current[receiptJobId] = setInterval(async () => {
      attempts += 1;

      try {
        const res = await apiRequest(
          API_ENDPOINTS.GET_RECEIPT_STATUS(receiptJobId)
        );
        if (!res.success || !res.data) {
          throw new Error(res.error || "Failed to fetch receipt status");
        }

        const {
          status: nextStatus,
          receipt_score: nextReceiptScore,
          error_message: nextErrorMessage,
        } = res.data;

        setReceipts((prev) =>
          prev.map((r) => {
            if (r.receiptJobId !== receiptJobId) return r;
            return {
              ...r,
              status: nextStatus,
              receiptScore: nextReceiptScore ?? null,
              errorMessage: nextErrorMessage ?? null,
            };
          })
        );

        const s = String(nextStatus ?? "").toLowerCase();
        if (
          s === "verified" ||
          s === "review" ||
          s === "rejected" ||
          s === "error"
        ) {
          clearInterval(pollingIntervalsRef.current[receiptJobId]);
          delete pollingIntervalsRef.current[receiptJobId];
        }
      } catch (err) {
        if (attempts >= maxAttempts) {
          setReceipts((prev) =>
            prev.map((r) => {
              if (r.receiptJobId !== receiptJobId) return r;
              return {
                ...r,
                status: "error",
                errorMessage: err.message || "Receipt verification failed",
              };
            })
          );
          clearInterval(pollingIntervalsRef.current[receiptJobId]);
          delete pollingIntervalsRef.current[receiptJobId];
        }
      }
    }, 2500);
  };

  // Resume polling for any pending receipts (e.g. after returning from UploadProof).
  useEffect(() => {
    if (form.update_type !== "Execution") return;
    receipts.forEach((r) => {
      const s = String(r.status ?? "").toLowerCase();
      if (
        r.receiptJobId &&
        (s === "" || s === "pending" || s === "processing")
      ) {
        startPollingReceipt(r.receiptJobId);
      }
    });
  }, [receipts, form.update_type]);
  // }

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
    // new {
    if (!claimedAmountValid) {
      setError("Enter a valid claimed amount before uploading receipts.");
      return;
    }
    // }
    const formData = new FormData();
    formData.append("file", file);
    // new {
    formData.append("claimed_amount", String(claimedAmountNumber));
    // }
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
      // if (!data.url) {
      // new {
      if (!data.url || !data.receipt_id) {  //}
        throw new Error("Upload did not return receipt URL");
      }

      // setReceipts((prev) => [...prev, data.url]);
      // new {
      const receiptJobId = data.receipt_id;
      const newReceipt = {
        url: data.url,
        receiptJobId,
        status: data.status || "pending",
        receiptScore: data.receipt_score ?? null,
        errorMessage: data.error_message ?? null,
      };
      setReceipts((prev) => {
        const nextReceipts = [...prev, newReceipt];
        saveDraft(form, nextReceipts);
        return nextReceipts;
      });
      startPollingReceipt(receiptJobId);
      // }
    } catch (err) {
      console.error(err);
      setError(err.message || "Failed to upload receipt");
    } finally {
      setReceiptUploadLoading(false);
      e.target.value = "";
    }
  };

  const handleRemoveReceipt = (indexToRemove) => {
    // setReceipts((prev) => prev.filter((_, i) => i !== indexToRemove));
    // new {
    setReceipts((prev) => {
      const removed = prev[indexToRemove];
      if (removed?.receiptJobId && pollingIntervalsRef.current[removed.receiptJobId]) {
        clearInterval(pollingIntervalsRef.current[removed.receiptJobId]);
        delete pollingIntervalsRef.current[removed.receiptJobId];
      }
      const nextReceipts = prev.filter((_, i) => i !== indexToRemove);
      saveDraft(form, nextReceipts);
      return nextReceipts;
    });
    // }
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

    if (form.update_type === "Execution") {
      const raw = String(form.claimed_amount ?? "").trim();
      if (!raw) {
        setError("Claimed amount is required for Execution updates.");
        return;
      }
      const amount = Number(raw);
      if (!Number.isFinite(amount) || amount <= 0) {
        setError("Claimed amount must be a positive number.");
        return;
      }
      payload.claimed_amount = amount;
    }

    if (form.update_type === "Execution" && receipts.length > 0) {
      // payload.receipt_urls = receipts;
      // new {
      const missingJobId = receipts.some((r) => !r.receiptJobId);
      if (missingJobId) {
        setError(
          "Receipt verification is required. Please remove and re-upload receipts."
        );
        return;
      }

      const pending = receipts.filter((r) => {
        const s = String(r.status ?? "").toLowerCase();
        return s === "" || s === "pending" || s === "processing";
      });
      if (pending.length > 0) {
        setError("Please wait until all receipts are verified before posting.");
        return;
      }

      const errored = receipts.filter((r) =>
        String(r.status ?? "").toLowerCase() === "error"
      );
      if (errored.length > 0) {
        setError("Receipt verification failed. Please remove and re-upload receipts.");
        return;
      }

      payload.receipt_urls = receipts.map((r) => r.url);
      payload.receipt_job_ids = receipts.map((r) => r.receiptJobId);
      // }
    }

    if (form.update_type === "Execution" && proofSessionIdFromQuery) {
      payload.proof_session_id = proofSessionIdFromQuery;
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

    // new {
    try {
      localStorage.removeItem(draftKey);
    } catch {
      // ignore
    }
    // }
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
              Claimed Amount (Execution updates only){" "}
              <span className="text-red-500">*</span>
            </label>
            <input
              type="number"
              name="claimed_amount"
              value={form.claimed_amount}
              onChange={handleChange}
              onWheel={(e) => e.target.blur()}
              min="0"
              step="0.01"
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
              placeholder="e.g. 2500.00"
            />
            <p className="text-xs text-gray-500 mt-1">
              This will be used to verify receipts against the claimed spend.
            </p>
          </div>
        )}

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
              // disabled={form.update_type !== "Execution"}
              // new {
              disabled={form.update_type !== "Execution" || !claimedAmountValid}
              // }
              className="w-full border rounded-md p-2 focus:outline-none focus:border-[#ff6200] cursor-pointer"
            />
            {/* new  */}
            {form.update_type === "Execution" && !claimedAmountValid && (
              <p className="text-xs text-red-500 mt-1">
                Enter a valid claimed amount to upload receipts.
              </p>
            )}
            {/* } */}
            {receiptUploadLoading && (
              <p className="text-xs text-gray-500 mt-1">
                Uploading receipt...
              </p>
            )}
            {receipts.length > 0 && (
              <div className="mt-3 grid grid-cols-2 md:grid-cols-3 gap-3">
                {/* {receipts.map((url, idx) => {
                  const src = url.startsWith("http")
                    ? url
                    : `${API_BASE_URL}${url}`; */}
                {/* new { */}
                {receipts.map((r, idx) => {
                  const src = r.url.startsWith("http")
                    ? r.url
                    : `${API_BASE_URL}${r.url}`;
                  const s = String(r.status ?? "").toLowerCase();
                  // }
                  return (
                    <div key={idx} className="relative group">
                      <img
                        src={src}
                        alt={`Receipt ${idx + 1}`}
                        className="w-full h-24 object-cover rounded border"
                      />
                      {/* new { */}
                      <div className="mt-2 text-xs">
                        <p className="font-medium text-gray-800">Verification:</p>
                        {s === "verified" && (
                          <span className="text-green-700 font-semibold">
                            Verified
                          </span>
                        )}
                        {s === "review" && (
                          <span className="text-yellow-700 font-semibold">
                            Needs Review
                          </span>
                        )}
                        {s === "rejected" && (
                          <span className="text-red-700 font-semibold">
                            Rejected
                          </span>
                        )}
                        {(s === "" || s === "pending" || s === "processing") && (
                          <span className="text-gray-600 font-semibold">
                            Pending...
                          </span>
                        )}
                        {s === "error" && (
                          <span className="text-red-700 font-semibold">
                            Error
                          </span>
                        )}

                        {r.receiptScore != null && (
                          <p className="text-gray-600 mt-1">
                            Score:{" "}
                            {Number(r.receiptScore).toFixed(2)}
                          </p>
                        )}
                        {r.errorMessage && (
                          <p className="text-red-600 mt-1">
                            {r.errorMessage}
                          </p>
                        )}
                      </div>
                      {/* } */}
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

            {proofsLoading && (
              <p className="text-xs text-gray-500 mt-3">Loading proofs...</p>
            )}

            {proofsError && (
              <p className="text-xs text-red-500 mt-3">{proofsError}</p>
            )}

            {Array.isArray(proofs) && proofs.length > 0 && (
              <div className="mt-4 grid grid-cols-2 md:grid-cols-3 gap-3">
                {proofs.map((p, idx) => {
                  const src = p.image?.startsWith("http")
                    ? p.image
                    : p.image
                      ? `${API_BASE_URL}/uploads/${p.image}`
                      : null;
                  return (
                    <div
                      key={p.image ? p.image : idx}
                      className="border rounded-lg p-2 bg-white"
                    >
                      {src ? (
                        <img
                          src={src}
                          alt={`Proof ${idx + 1}`}
                          className="w-full h-24 object-cover rounded border"
                        />
                      ) : (
                        <div className="w-full h-24 bg-gray-100 flex items-center justify-center text-xs text-gray-400 rounded border">
                          Image not found
                        </div>
                      )}

                      <div className="mt-2 text-xs">
                        <p className="font-medium text-gray-800">AI Status</p>
                        <span
                          className={`inline-block mt-1 px-2 py-0.5 rounded font-semibold ${getProofStatusStyle(
                            p.aiStatus
                          )}`}
                        >
                          {p.aiStatus || "unknown"}
                        </span>
                      </div>
                    </div>
                  );
                })}
              </div>
            )}
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
            // disabled={submitting}
            // new {
            disabled={submitting || isReceiptVerificationInProgress}
            // }
            className="px-6 py-2 rounded-lg bg-[#ff6200] hover:bg-[#e45a00] text-white text-sm font-semibold disabled:opacity-60 cursor-pointer"
          >
            {/* {submitting ? "Posting..." : "Post Update"} */}
            {/* new { */}
            {submitting
              ? "Posting..."
              : isReceiptVerificationInProgress
                ? "Verifying receipts..."
                : "Post Update"}
            {/* } */}
          </button>
        </div>
      </form>
    </div>
  );
}

