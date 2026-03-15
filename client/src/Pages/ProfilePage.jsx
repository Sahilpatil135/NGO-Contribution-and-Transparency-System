import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";
import { apiRequest, API_ENDPOINTS, } from "../config/api";
import { getCauseImage } from "../utils/imageHelper";
import heroimg from "/default_user_avatar.jpg";
import causePlaceholder from "../../public/domains/domain_example.png";
import { LuPencil, LuHeart, LuLock, LuShieldCheck } from "react-icons/lu";

const ProfilePage = () => {
  const { user, fetchCurrentUser } = useAuth();
  const [donations, setDonations] = useState([]);
  const [causesMap, setCausesMap] = useState({});
  const [loadingDonations, setLoadingDonations] = useState(true);
  const [isEditing, setIsEditing] = useState(false);
  const [editForm, setEditForm] = useState({ name: "" });
  const [saving, setSaving] = useState(false);
  const [saveError, setSaveError] = useState(null);

  useEffect(() => {
    const fetchDonations = async () => {
      const result = await apiRequest(API_ENDPOINTS.GET_MY_DONATIONS);
      if (result.success && result.data) {
        setDonations(result.data);
        const uniqueCauseIds = [...new Set(result.data.map((d) => d.cause_id).filter(Boolean))];
        const causes = {};
        await Promise.all(
          uniqueCauseIds.map(async (id) => {
            const req = await apiRequest(`${API_ENDPOINTS.GET_CAUSES}/${id}`);
            if (req.success && req.data) causes[id] = req.data;
          })
        );
        setCausesMap(causes);
      }
      setLoadingDonations(false);
    };
    fetchDonations();
  }, []);

  useEffect(() => {
    if (user) {
      setEditForm({
        name: user.name || "",
      });
    }
  }, [user]);

  const handleEditChange = (e) => {
    const { name, value } = e.target;
    setEditForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSave = async () => {
    setSaving(true);
    setSaveError(null);
    const payload = {};
    if (editForm.name.trim()) payload.name = editForm.name.trim();

    const result = await apiRequest(API_ENDPOINTS.UPDATE_PROFILE, {
      method: "PATCH",
      body: JSON.stringify(payload),
    });

    if (result.success && result.data) {
      await fetchCurrentUser();
      setIsEditing(false);
    } else {
      setSaveError(result.error || "Failed to update profile");
    }
    setSaving(false);
  };

  const toTitleCase = (s = "") =>
    s ? s.charAt(0).toUpperCase() + s.slice(1).toLowerCase() : "";

  if (!user) {
    return (
      <div className="max-w-4xl mx-auto px-4 py-12 text-center">
        <p className="text-gray-600">Loading profile...</p>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto px-4 py-8">
      {/* Profile header */}
      <div className="bg-white rounded-xl shadow-lg border border-gray-200 overflow-hidden">
        <div className="h-24 bg-gradient-to-r from-[#ff6200]/20 via-amber-100/50 to-[#3a0b2e]/10" />
        <div className="px-6 pb-6 -mt-12 relative">
          <div className="flex flex-col sm:flex-row sm:items-end gap-4">
            <img
              src={user.avatar_url || heroimg}
              alt="Avatar"
              className="w-24 h-24 rounded-full border-4 border-white shadow-lg object-cover"
            />
            <div className="flex-1">
              {isEditing ? (
                <div className="space-y-3">
                  <input
                    type="text"
                    name="name"
                    value={editForm.name}
                    onChange={handleEditChange}
                    placeholder="Your name"
                    className="w-full max-w-md border border-gray-300 rounded-lg py-2 px-3 focus:outline-none focus:border-[#ff6200]"
                  />
                  <div className="flex gap-2">
                    <button
                      onClick={handleSave}
                      disabled={saving}
                      className="bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-4 py-2 rounded-lg transition cursor-pointer disabled:opacity-60"
                    >
                      {saving ? "Saving..." : "Save"}
                    </button>
                    <button
                      onClick={() => {
                        setIsEditing(false);
                        setEditForm({
                          name: user.name || "",
                        });
                      }}
                      disabled={saving}
                      className="bg-gray-200 hover:bg-gray-300 text-gray-800 font-semibold px-4 py-2 rounded-lg transition cursor-pointer"
                    >
                      Cancel
                    </button>
                  </div>
                  {saveError && (
                    <p className="text-red-500 text-sm">{saveError}</p>
                  )}
                </div>
              ) : (
                <>
                  <h1 className="text-2xl font-bold text-[#3a0b2e]">
                    {toTitleCase(user.name)}
                  </h1>
                  <p className="text-gray-600 mt-1">{user.email}</p>
                  <p className="text-sm text-gray-500 mt-1">
                    Role: {toTitleCase(user.role)}
                  </p>
                  {/*<button
                    onClick={() => setIsEditing(true)}
                    className="mt-3 inline-flex items-center gap-2 text-[#ff6200] hover:text-[#e45a00] font-medium text-sm transition cursor-pointer"
                  >
                    <LuPencil className="w-4 h-4" />
                    Edit profile
                  </button>*/}
                </>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* Donations section */}
      <div className="mt-8">
        <h2 className="text-xl font-bold text-[#3a0b2e] mb-4 flex items-center gap-2">
          <LuHeart className="text-[#ff6200]" />
          My Donations
        </h2>

        {loadingDonations ? (
          <div className="space-y-3">
            <div className="h-20 rounded-lg bg-gray-100 animate-pulse" />
            <div className="h-20 rounded-lg bg-gray-100 animate-pulse" />
            <div className="h-20 rounded-lg bg-gray-100 animate-pulse" />
          </div>
        ) : donations.length === 0 ? (
          <div className="bg-white rounded-xl border border-gray-200 p-8 text-center">
            <p className="text-gray-600">You haven&apos;t made any donations yet.</p>
            <Link
              to="/makeContribution"
              className="inline-block mt-4 text-[#ff6200] hover:text-[#e45a00] font-semibold"
            >
              Explore causes to donate →
            </Link>
          </div>
        ) : (
          <>
            <p className="text-sm text-gray-600 mb-4">
              You've donated to {donations.length} {donations.length === 1 ? "cause" : "causes"}. Total of ₹
              {donations
                .reduce((sum, d) => sum + parseFloat(d.amount || 0), 0)
                .toLocaleString()}.
              Thank You for Your Contributions!
            </p>
            <div className="space-y-3">
              {donations.map((d) => {
                const date = d.created_at ? new Date(d.created_at) : null;
                const dateStr = date
                  ? date.toLocaleDateString("en-IN", {
                    day: "numeric",
                    month: "short",
                    year: "numeric",
                  })
                  : "—";
                const amount = parseFloat(d.amount || 0);
                const cause = d.cause_id ? causesMap[d.cause_id] : null;
                return (
                  <Link
                    key={d.id}
                    to={`/campaign/${d.cause_id}`}
                    className="block bg-white rounded-xl border border-gray-200 overflow-hidden shadow-sm hover:shadow-md transition"
                  >
                    <div className="flex flex-col sm:flex-row p-4 gap-4">
                      <div className="w-35 h-20 sm:w-35 sm:h-24 flex-shrink-0 rounded-md overflow-hidden">
                        <img
                          src={getCauseImage(cause?.cover_image_url, causePlaceholder)}
                          alt={cause?.title || "Campaign"}
                          className="w-full h-full object-cover"
                        />
                      </div>
                      <div className="flex-1 min-w-0 flex flex-col sm:flex-row sm:items-start sm:justify-between gap-3">
                        {/* Left: basic donation info */}
                        <div className="space-y-1">
                          <h3 className="font-semibold text-[#3a0b2e] line-clamp-2">
                            {cause?.title || "Campaign"}
                          </h3>
                          {cause?.organization?.name && (
                            <p className="text-sm text-gray-500 mt-0.5">
                              by {cause.organization.name}
                            </p>
                          )}
                          <p className="text-xs text-gray-500 mt-2">
                            Donated on {dateStr}
                            {d.status && (
                              <span className="ml-2">
                                • <span className="capitalize">{d.status}</span>
                              </span>
                            )}
                          </p>
                        </div>
                        {/* Right: financial + blockchain details */}
                        <div className="flex items-end justify-between sm:flex-col sm:items-end gap-3 text-xs">
                          <div className="flex flex-col items-end gap-1">
                            <p className="text-[11px] uppercase tracking-wide text-gray-500 font-semibold">
                              Amount
                            </p>
                            <p className="font-bold text-[#ff6200] text-lg">
                              ₹{amount.toLocaleString("en-IN")}
                            </p>
                          </div>

                          {(d.payment_id || d.tx_hash) && (
                            <div className="flex flex-col items-end gap-1 text-[11px] text-gray-500 max-w-[260px]">
                              {d.payment_id && (
                                <p className="text-right">
                                  <span className="text-gray-500">Payment Ref</span>{" "}
                                  <span className="text-gray-400">•</span>{" "}
                                  <span className="font-mono break-all text-gray-700">
                                    {d.payment_id}
                                  </span>
                                </p>
                              )}
                            </div>
                          )}

                          <span className="text-[#ff6200] hover:text-[#e45a00] font-medium text-sm whitespace-nowrap">
                            View campaign →
                          </span>
                        </div>
                      </div>
                    </div>
                  </Link>
                );
              })}
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default ProfilePage;
