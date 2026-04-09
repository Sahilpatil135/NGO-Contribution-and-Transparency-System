import { useState, useEffect, useRef } from "react";
import img1 from "../../public/domains/domain_example.png";
import { Link, useParams } from "react-router-dom";
import { apiRequest, API_ENDPOINTS, API_BASE_URL } from "../config/api";
import { getCauseImage } from "../utils/imageHelper";
import { formatGoal, formatCollected, getGoalLabel, getCollectedLabel } from "../utils/goalHelper";
import { LuLink2, LuLock, LuShieldCheck } from "react-icons/lu";
import { BiUpvote, BiDownvote } from "react-icons/bi";
import { useAuth } from "../contexts/AuthContext";

const PRODUCT_AID_TYPE_NAMES = [
  "Goods & Resources",
  "Educational Support",
  "Medical Assistance",
  "Environmental Support",
  "Disaster Relief Assistance",
];

const BLOOD_DONATION_AID_TYPE_NAME = "Blood Donations";
const VOLUNTEERING_AID_TYPE_NAME = "Volunteering";

const CampaignPage = () => {
  const { user, organization } = useAuth();
  const { causeID } = useParams();
  const [loading, setLoading] = useState(true);
  const [cause, setCause] = useState({});
  const [activeTab, setActiveTab] = useState("project");
  const [donations, setDonations] = useState([]);
  const [donationsLoading, setDonationsLoading] = useState(true);
  const [donationsError, setDonationsError] = useState(null);
  const [selectedImage, setSelectedImage] = useState(null);
  const [votes, setVotes] = useState({ upvotes: 0, downvotes: 0, my_vote: null });
  const [votesLoading, setVotesLoading] = useState(false);
  const [reviews, setReviews] = useState({ count: 0, reviews: [] });
  const [reviewsLoading, setReviewsLoading] = useState(false);
  const [reviewsSubmitting, setReviewsSubmitting] = useState(false);
  const [canReview, setCanReview] = useState(false);
  const [reviewText, setReviewText] = useState("");
  const [proofsBySession, setProofsBySession] = useState({});
  const [proofsLoadingBySession, setProofsLoadingBySession] = useState({});
  const proofsFetchedRef = useRef(new Set());
  const [totalDisbursed, setTotalDisbursed] = useState(0);
  const [disbursementsLoading, setDisbursementsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      const causeResult = await apiRequest(
        `${API_ENDPOINTS.GET_CAUSES}/${causeID}`
      );

      if (causeResult.success && causeResult.data) {
        setCause(causeResult.data);
        setLoading(false);
      }

      const donationsResult = await apiRequest(
        API_ENDPOINTS.GET_CAUSE_CHAIN_DONATIONS(causeID)
      );

      if (donationsResult.success && donationsResult.data) {
        setDonations(donationsResult.data || []);
      } else if (!donationsResult.success && donationsResult.error) {
        setDonationsError(donationsResult.error);
      }

      setDonationsLoading(false);

      // Fetch upvote/downvote counters for this cause.
      setVotesLoading(true);
      const votesResult = await apiRequest(
        API_ENDPOINTS.GET_CAUSE_VOTES(causeID)
      );
      if (votesResult.success && votesResult.data) {
        setVotes(votesResult.data);
      }
      setVotesLoading(false);

      // Fetch whether the current user can review (must have donated)
      // and load existing textual reviews.
      setReviewsLoading(true);
      const [myDonationsResult, reviewsResult] = await Promise.all([
        apiRequest(API_ENDPOINTS.GET_MY_DONATIONS),
        apiRequest(API_ENDPOINTS.GET_CAUSE_REVIEWS(causeID)),
      ]);

      if (myDonationsResult.success && Array.isArray(myDonationsResult.data)) {
        const hasDonated = myDonationsResult.data.some(
          (d) => String(d.cause_id) === String(causeID)
        );

        const hasReviewed = reviewsResult.data["reviews"].some(
          (d) => String(d.user_id) === String(user.id)
        )

        setCanReview(hasDonated && !hasReviewed);
      } else {
        setCanReview(false);
      }

      if (reviewsResult.success && reviewsResult.data) {
        setReviews(reviewsResult.data);
      } else {
        setReviews({ count: 0, reviews: [] });
      }
      setReviewsLoading(false);

      // Fetch disbursements for this cause
      setDisbursementsLoading(true);
      const disbursementsResult = await apiRequest(
        API_ENDPOINTS.GET_CAUSE_DISBURSEMENTS(causeID)
      );
      if (disbursementsResult.success && disbursementsResult.data) {
        setTotalDisbursed(disbursementsResult.data.total_disbursed || 0);
      }
      setDisbursementsLoading(false);
    };

    fetchData();
  }, [causeID]);

  const handleUpvote = async () => {
    if (!causeID) return;
    setVotesLoading(true);
    const res = await apiRequest(API_ENDPOINTS.UPVOTE_CAUSE(causeID), {
      method: "POST",
    });
    if (res.success && res.data) setVotes(res.data);
    setVotesLoading(false);
  };

  const handleDownvote = async () => {
    if (!causeID) return;
    setVotesLoading(true);
    const res = await apiRequest(API_ENDPOINTS.DOWNVOTE_CAUSE(causeID), {
      method: "POST",
    });
    if (res.success && res.data) setVotes(res.data);
    setVotesLoading(false);
  };

  const getDaysLeft = (targetTimestamp) => {
    if (!targetTimestamp) return 0;
    const now = Date.now();
    const targetTime = new Date(targetTimestamp).getTime();
    const differenceInMs = targetTime - now;
    const millisecondsPerDay = 1000 * 60 * 60 * 24;
    const daysLeft = Math.ceil(differenceInMs / millisecondsPerDay);
    return Math.max(0, daysLeft);
  };

  // const getCoverImageSrc = () => {
  //   if (cause.cover_image_url) {
  //     const url = cause.cover_image_url.startsWith("http")
  //       ? cause.cover_image_url
  //       : `${API_BASE_URL}${cause.cover_image_url}`;
  //     return url;
  //   }
  //   return img1;
  // };

  const showProductsTab =
    cause.aid_type &&
    PRODUCT_AID_TYPE_NAMES.includes(cause.aid_type.name) &&
    cause.products &&
    cause.products.length > 0;

  const aidTypeName = cause?.aid_type?.name || "";
  const fundingStatus = cause.funding_status || "Not Started";
  const goal = parseFloat(cause.goal_amount) || 0;
  const collected = parseFloat(cause.collected_amount) || 0;
  const disbursed = parseFloat(totalDisbursed) || 0;
  const percentage = goal > 0 ? Math.min((collected / goal) * 100, 100) : 0;
  const disbursedPercentage = goal > 0 ? Math.min((disbursed / goal) * 100, 100) : 0;
  const daysLeft = getDaysLeft(cause.deadline);

  const tabs = [
    ...(showProductsTab ? [{ id: "products", label: "Products" }] : []),
    { id: "project", label: "Project" },
    { id: "updates", label: "Updates" },
    { id: "reviews", label: "Reviews" },
    { id: "donations", label: "Donations" },
  ];

  const isOwner =
    user?.role === "organization" &&
    organization?.id &&
    cause?.organization?.id &&
    String(organization.id) === String(cause.organization.id);

  const getScoreStyle = (score) => {
    if (score >= 70)
      return "bg-blue-100 text-blue-700";
    if (score >= 40)
      return "bg-amber-100 text-amber-700";
    return "bg-red-100 text-red-700";
  };

  const getProofStatusStyle = (status) => {
    const s = String(status ?? "").toLowerCase();
    if (s === "verified") return "bg-green-100 text-green-700";
    if (s === "review") return "bg-amber-100 text-amber-700";
    if (s === "rejected") return "bg-red-100 text-red-700";
    return "bg-gray-100 text-gray-700";
  };

  // Fetch proof-of-work images for execution updates (using stored proof_session_id).
  useEffect(() => {
    const updates = Array.isArray(cause?.updates) ? cause.updates : [];
    const sessionIds = Array.from(
      new Set(
        updates
          .map((u) => u?.proof_session_id)
          .filter((id) => id)
          .map((id) => String(id))
      )
    );

    const toFetch = sessionIds.filter(
      (id) => !proofsFetchedRef.current.has(id)
    );
    if (toFetch.length === 0) return;

    toFetch.forEach((id) => proofsFetchedRef.current.add(id));

    const fetchProofs = async () => {
      // Mark loading up-front so the UI can show placeholders.
      setProofsLoadingBySession((prev) => {
        const next = { ...prev };
        toFetch.forEach((id) => (next[id] = true));
        return next;
      });

      await Promise.all(
        toFetch.map(async (sessionId) => {
          const res = await apiRequest(
            API_ENDPOINTS.GET_PROOF_IMAGES_BY_SESSION(sessionId)
          );
          if (res.success) {
            setProofsBySession((prev) => ({
              ...prev,
              [sessionId]: Array.isArray(res.data) ? res.data : [],
            }));
          }
          setProofsLoadingBySession((prev) => ({
            ...prev,
            [sessionId]: false,
          }));
        })
      );
    };

    fetchProofs();
  }, [cause?.updates]);

  const canDonate = fundingStatus !== "Fully Funded" && fundingStatus !== "Closed";
  const isBloodDonationAid =
    cause.aid_type?.name === BLOOD_DONATION_AID_TYPE_NAME;
  const isVolunteeringAid =
    cause.aid_type?.name === VOLUNTEERING_AID_TYPE_NAME;

  // Set default active tab when products tab is not available
  useEffect(() => {
    if (!showProductsTab && activeTab === "products") {
      setActiveTab("project");
    }
  }, [showProductsTab, activeTab]);

  const formatDonationAmount = (amount) => {
    if (!amount) return "-";
    const numeric = Number(amount);

    if (!Number.isFinite(numeric) || Number.isNaN(numeric)) {
      return amount.toString();
    }

    return `₹${numeric.toFixed(2)}`;
  };

  const formatDonationTimestamp = (timestamp) => {
    if (!timestamp) return "";
    const numeric = Number(timestamp);

    if (!Number.isFinite(numeric) || Number.isNaN(numeric)) {
      return "";
    }

    // On-chain timestamps are usually in seconds
    const date = new Date(numeric * 1000);
    if (Number.isNaN(date.getTime())) return "";
    return date.toLocaleString();
  };

  if (loading) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-12 text-center">
        <p className="text-gray-600">Loading campaign...</p>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-4 lg:px-4 py-8">
      {/* SECTION 1: HEADER */}
      <div className="mb-8">
        <div className="flex flex-wrap items-center gap-2 mb-2">
          <span
            className={`inline-flex px-3 py-1 rounded-full text-sm font-medium ${fundingStatus === "Fully Funded"
              ? "bg-green-100 text-green-800"
              : fundingStatus === "Closed"
                ? "bg-gray-200 text-gray-700"
                : fundingStatus === "Active"
                  ? "bg-amber-100 text-amber-800"
                  : "bg-slate-100 text-slate-700"
              }`}
          >
            {fundingStatus}
          </span>
          {cause.domain && (
            <span className="text-gray-500 text-sm">• {cause.domain.name}</span>
          )}
          {cause.aid_type && (
            <span className="text-gray-500 text-sm">• {cause.aid_type.name}</span>
          )}
        </div>
        <h1 className="text-4xl md:text-5xl font-bold text-[#3a0b2e]">
          {cause.title}
        </h1>
        <h3 className="text-xl text-gray-500 mt-3">
          Campaign by{" "}
          {cause?.organization?.id ? (
            <Link
              to={`/organization/${cause.organization.id}/accounts`}
              className="font-semibold text-gray-800 hover:text-[#ff6200]"
            >
              {cause.organization?.name || "NGO"}
            </Link>
          ) : (
            <span className="font-semibold text-gray-800">
              {cause.organization?.name || "NGO"}
            </span>
          )}
        </h3>
        <div className="flex items-center gap-4 mt-4">
          <button
            type="button"
            className={`inline-flex items-center gap-2 px-3 py-1.5 rounded-full text-sm font-semibold border transition cursor-pointer ${votes.my_vote === "up"
              ? "bg-green-50 border-green-200 text-green-700"
              : "bg-white border-gray-200 text-gray-600 hover:border-green-200"
              }`}
            disabled={votesLoading}
            onClick={handleUpvote}
          >
            <BiUpvote
              className={
                votes.my_vote === "up" ? "text-green-600" : "text-gray-500"
              }
            />
            <span>{votes.upvotes}</span>
          </button>
          <button
            type="button"
            className={`inline-flex items-center gap-2 px-3 py-1.5 rounded-full text-sm font-semibold border transition cursor-pointer ${votes.my_vote === "down"
              ? "bg-red-50 border-red-200 text-red-700"
              : "bg-white border-gray-200 text-gray-600 hover:border-red-200"
              }`}
            disabled={votesLoading}
            onClick={handleDownvote}
          >
            <BiDownvote
              className={
                votes.my_vote === "down" ? "text-red-600" : "text-gray-500"
              }
            />
            <span>{votes.downvotes}</span>
          </button>
        </div>
        <div className="flex flex-wrap gap-4 mt-3 text-gray-600 text-sm">
          {cause.execution_location && (
            <span>📍 {cause.execution_location}</span>
          )}
          {cause.beneficiaries_count != null && cause.beneficiaries_count > 0 && (
            <span>👥 {cause.beneficiaries_count} beneficiaries</span>
          )}
          {cause.deadline && (
            <span>📅 Deadline: {new Date(cause.deadline).toLocaleDateString()}</span>
          )}
          {typeof reviews?.count === "number" && (
            <span>📝 {reviews.count} Reviews</span>
          )}
        </div>
      </div>

      {/* Grid: Left content | Right stats sidebar */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Left: Cover image + Tabs + Tab content */}
        <div className="md:col-span-2">
          <img
            // src={getCoverImageSrc()}
            src={getCauseImage(cause.cover_image_url, img1)}
            alt={cause.title}
            className="w-full object-cover rounded-lg h-64 md:h-96"
          />

          {/* SECTION 3: TABS */}
          <div className="flex gap-2 mt-6 border-b border-gray-200">
            {tabs.map((tab) => (
              <button
                key={tab.id}
                onClick={() => setActiveTab(tab.id)}
                className={`px-4 py-3 text-sm font-medium border-b-2 -mb-px transition-colors cursor-pointer ${activeTab === tab.id
                  ? "border-[#ff6200] text-[#ff6200]"
                  : "border-transparent text-gray-500 hover:text-gray-700"
                  }`}
              >
                {tab.label}
              </button>
            ))}
          </div>

          {/* Tab content */}
          {activeTab === "products" && showProductsTab && (
            <div className="mt-6 space-y-6">
              {cause.products.map((p) => {
                const remaining =
                  (p.quantity_needed || 0) - (p.quantity_funded || 0);
                return (
                  <div
                    key={p.id}
                    className="border rounded-lg p-5 bg-white shadow-sm flex flex-col md:flex-row gap-4"
                  >
                    <div className="md:w-32 flex-shrink-0">
                      {p.image_url ? (
                        <img
                          src={
                            p.image_url.startsWith("http")
                              ? p.image_url
                              : `${API_BASE_URL}${p.image_url}`
                          }
                          alt={p.name}
                          className="w-full h-28 object-cover rounded"
                        />
                      ) : (
                        <div className="w-full h-28 bg-gray-200 rounded flex items-center justify-center text-gray-400 text-sm">
                          No image
                        </div>
                      )}
                    </div>
                    <div className="flex-1">
                      <h4 className="font-semibold text-[#3a0b2e]">{p.name}</h4>
                      <p className="text-gray-600 text-sm mt-1">
                        {p.description}
                      </p>
                      <div className="mt-3 flex flex-wrap gap-4 text-sm">
                        <span>₹{parseFloat(p.price_per_unit || 0).toLocaleString()} per unit</span>
                        <span>Needed: {p.quantity_needed || 0}</span>
                        <span>Funded: {p.quantity_funded || 0}</span>
                        <span>Remaining: {remaining}</span>
                      </div>
                      <Link
                        to="/checkout"
                        state={{ causeID: cause.id }}
                        className="inline-block mt-3"
                      >
                        <button className="bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-4 py-2 rounded-lg text-sm transition cursor-pointer">
                          Sponsor This Item
                        </button>
                      </Link>
                    </div>
                  </div>
                );
              })}
            </div>
          )}

          {activeTab === "project" && (
            <div className="mt-6 space-y-6">
              <section>
                <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                  About the Campaign
                </h3>
                <p className="text-gray-700 leading-relaxed whitespace-pre-line">
                  {cause.description || "No description provided."}
                </p>
              </section>

              {cause.problem_statement && (
                <section>
                  <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                    Problem Statement
                  </h3>
                  <p className="text-gray-700 leading-relaxed whitespace-pre-line">
                    {cause.problem_statement}
                  </p>
                </section>
              )}

              {cause.execution_plan && (
                <section>
                  <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                    Execution Plan
                  </h3>
                  <p className="text-gray-700 leading-relaxed whitespace-pre-line">
                    {cause.execution_plan}
                  </p>
                </section>
              )}

              {(cause.beneficiaries_count != null ||
                cause.execution_location) && (
                  <section>
                    <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                      Beneficiaries & Location
                    </h3>
                    <p className="text-gray-700">
                      {cause.beneficiaries_count != null &&
                        cause.beneficiaries_count > 0 && (
                          <span>
                            {cause.beneficiaries_count} beneficiaries
                            {cause.execution_location && " • "}
                          </span>
                        )}
                      {cause.execution_location}
                    </p>
                  </section>
                )}

              {(cause.execution_start_time || cause.execution_end_time) && (
                <section>
                  <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                    Timeline
                  </h3>
                  <p className="text-gray-700">
                    {cause.execution_start_time &&
                      new Date(cause.execution_start_time).toLocaleDateString()}
                    {cause.execution_end_time && (
                      <>
                        {" → "}
                        {new Date(cause.execution_end_time).toLocaleDateString()}
                      </>
                    )}
                  </p>
                </section>
              )}

              {cause.impact_goal && (
                <section>
                  <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                    Impact Goal
                  </h3>
                  <p className="text-gray-700 leading-relaxed whitespace-pre-line">
                    {cause.impact_goal}
                  </p>
                </section>
              )}

              {(cause.execution_lat != null || cause.execution_lng != null) && (
                <section>
                  <h3 className="text-lg font-semibold text-[#3a0b2e] mb-2">
                    Execution Location (Coordinates)
                  </h3>
                  <p className="text-gray-700 text-sm">
                    Lat: {cause.execution_lat}, Lng: {cause.execution_lng}
                    {cause.execution_radius_meters != null && (
                      <span> • Radius: {cause.execution_radius_meters}m</span>
                    )}
                  </p>
                </section>
              )}
            </div>
          )}

          {activeTab === "updates" && (
            <div className="mt-6">
              {cause.updates && cause.updates.length > 0 ? (
                <div className="space-y-4">
                  {cause.updates.map((u) => {
                    const receipts =
                      Array.isArray(u.media) && u.media.length > 0
                        ? u.media.filter((m) => m.media_type === "receipt")
                        : [];

                    const proofSessionId = u?.proof_session_id
                      ? String(u.proof_session_id)
                      : null;
                    const proofs =
                      proofSessionId && proofsBySession[proofSessionId]
                        ? proofsBySession[proofSessionId]
                        : [];

                    console.log("score ", u)

                    return (
                      <div
                        key={u.id}
                        className="border rounded-lg p-4 bg-white shadow-sm"
                      >
                        <div className="flex items-center gap-2 mb-2">
                          <h4 className="font-semibold text-[#3a0b2e]">
                            {u.title}
                          </h4>
                          {u.is_verified && (
                            <span className="text-xs bg-green-100 text-green-700 px-2 py-0.5 rounded">
                              Verified
                            </span>
                          )}
                        </div>
                        <p className="text-gray-600 text-sm mb-2">
                          {u.description}
                        </p>
                        <p className="text-xs text-gray-400 mb-2">
                          {new Date(u.created_at).toLocaleDateString()}
                          {u.update_type && ` • ${u.update_type}`}
                          {u.funding_percentage != null &&
                            ` • ${u.funding_percentage}% funded`}
                        </p>
                        {u.update_type === "Execution" && receipts.length > 0 && (
                          <div className="mt-2">
                            <p className="text-sm font-semibold text-gray-800 mb-2">
                              Receipts
                            </p>
                            <div className="flex flex-wrap gap-3 mb-2">
                              <div className="text-sm text-gray-600">
                                <span className="font-medium text-gray-800"> Claimed Amount:</span>{" "}
                                ₹{u.claimed_amount}
                              </div>

                              <div className="text-sm text-gray-600">
                                <span className="font-medium text-gray-800"> Verification Score:</span>{" "}
                                {/* <span className="px-2 py-0.5 rounded bg-blue-100 text-blue-700 font-semibold"> */}
                                <span
                                  className={`px-2 py-0.5 rounded font-semibold ${getScoreStyle(
                                    u.verification_score
                                  )}`}
                                >
                                  {u.verification_score}
                                </span>
                              </div>

                              <div className="text-sm text-gray-600">
                                {receipts.some((m) =>
                                  String(m.media_url || "").startsWith("/api/ipfs/")
                                ) && (
                                    <p className="text-[11px] text-amber-700 bg-amber-50 border border-amber-100 px-2 py-1 rounded-full inline-block mb-3">
                                      Stored securely on IPFS
                                    </p>
                                  )}
                              </div>
                            </div>
                            <div className="grid grid-cols-2 md:grid-cols-3 gap-3">
                              {receipts.map((m) => {

                                const src = m.media_url.startsWith("http")
                                  ? m.media_url
                                  : `${API_BASE_URL}${m.media_url}`;

                                return (
                                  <div
                                    key={m.id}
                                    onClick={() => setSelectedImage(src)}
                                    className="relative group border rounded-lg overflow-hidden cursor-pointer"
                                  >
                                    <img
                                      src={src}
                                      alt="Receipt"
                                      className="w-full h-28 object-cover transition-transform duration-200 group-hover:scale-105"
                                    />

                                    {/* Hover overlay */}
                                    <div className="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 flex items-center justify-center transition">
                                      <span className="text-white text-xs">View</span>
                                    </div>

                                    {selectedImage && (
                                      <div
                                        className="fixed inset-0 bg-black/80 flex items-center justify-center z-50"
                                        onClick={() => setSelectedImage(null)}
                                      >
                                        <img
                                          src={selectedImage}
                                          alt="Full Receipt"
                                          className="max-h-[90%] max-w-[90%] rounded-lg shadow-lg"

                                        />
                                        <button
                                          className="absolute top-6 right-10 text-white text-xl cursor-pointer"
                                          onClick={(e) => {
                                            e.stopPropagation();
                                            setSelectedImage(null);
                                          }}
                                        >
                                          ✕
                                        </button>
                                      </div>
                                    )}

                                  </div>
                                );
                              })}
                            </div>
                          </div>
                        )}

                        {u.update_type === "Execution" && proofSessionId && (
                          <div className="mt-3">
                            <p className="text-sm font-semibold text-gray-800 mb-2">
                              Proof of Work
                            </p>

                            {proofsLoadingBySession[proofSessionId] ? (
                              <p className="text-xs text-gray-500">
                                Loading proofs...
                              </p>
                            ) : proofs.length > 0 ? (
                              <div className="grid grid-cols-2 md:grid-cols-3 gap-3">
                                {proofs.map((p, idx) => {
                                  const src = p?.image?.startsWith("http")
                                    ? p.image
                                    : p?.image
                                      ? `${API_BASE_URL}/uploads/${p.image}`
                                      : null;
                                  console.log("proofs data", p);
                                  return (
                                    <div
                                      key={p?.image ? p.image : idx}
                                      onClick={() => src && setSelectedImage(src)}
                                      className="relative group border rounded-lg overflow-hidden cursor-pointer bg-white"
                                    >
                                      {src ? (
                                        <img
                                          src={src}
                                          alt="Proof"
                                          className="w-full h-28 object-cover transition-transform duration-200 group-hover:scale-105"
                                        />
                                      ) : (
                                        <div className="w-full h-28 bg-gray-100 flex items-center justify-center text-xs text-gray-400">
                                          Image not found
                                        </div>
                                      )}


                                      <div className="flex mb-1">
                                        <div className="p-2">
                                          <p className="text-sm text-gray-500">
                                            AI Status
                                          </p>
                                          <span
                                            className={`inline-block mt-1 px-2 py-0.5 rounded text-sm font-semibold ${getProofStatusStyle(
                                              p.aiStatus
                                            )}`}
                                          >
                                            {p.aiStatus || "unknown"}
                                          </span>
                                        </div>
                                        <div className="p-2">
                                          <p className="text-sm text-gray-500">
                                            Score
                                          </p>
                                          <span
                                            className={`inline-block mt-1 px-2 py-0.5 rounded text-sm font-semibold ${getScoreStyle(
                                              p.score
                                            )}`}
                                          >
                                            {p.score || "unknown"}
                                          </span>
                                        </div>
                                      </div>
                                    </div>
                                  );
                                })}
                              </div>
                            ) : (
                              <p className="text-xs text-gray-500">
                                Proofs not found for this session.
                              </p>
                            )}
                          </div>
                        )}
                      </div>
                    );
                  })}
                </div>
              ) : (
                <p className="text-gray-600">
                  No updates yet. Be the first to support the campaign!
                </p>
              )}
            </div>
          )}

          {activeTab === "reviews" && (
            <div className="mt-6">
              {reviewsLoading ? (
                <div className="space-y-4">
                  <div className="h-16 rounded-xl bg-gray-100 animate-pulse" />
                  <div className="h-16 rounded-xl bg-gray-100 animate-pulse" />
                </div>
              ) : (
                <div className="space-y-6">
                  <div className="rounded-xl border border-gray-200 bg-white p-5 shadow-sm">
                    <h3 className="text-lg font-semibold text-[#3a0b2e]">
                      Write a review
                    </h3>
                    <p className="text-sm text-gray-600 mt-1">
                      Review can be left only after you donate.
                    </p>

                    {canReview ? (
                      <div className="mt-4 space-y-3">
                        <textarea
                          className="w-full border border-gray-300 rounded-lg p-3 focus:outline-none focus:border-[#ff6200] placeholder-gray-500"
                          rows={4}
                          placeholder="Share your experience in a few sentences..."
                          value={reviewText}
                          onChange={(e) => setReviewText(e.target.value)}
                        />
                        <button
                          type="button"
                          className="bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-6 py-2 rounded-lg transition cursor-pointer disabled:opacity-60"
                          disabled={reviewsSubmitting}
                          onClick={async () => {
                            if (!causeID) return;
                            const text = reviewText.trim();
                            if (text.length < 5) return;
                            setReviewsSubmitting(true);
                            const res = await apiRequest(
                              API_ENDPOINTS.CREATE_CAUSE_REVIEW(causeID),
                              {
                                method: "POST",
                                body: JSON.stringify({ review_text: text }),
                              }
                            );
                            if (res.success) {
                              setReviewText("");
                              const refresh = await apiRequest(
                                API_ENDPOINTS.GET_CAUSE_REVIEWS(causeID)
                              );
                              if (refresh.success && refresh.data) {
                                setReviews(refresh.data);
                              }
                            }
                            setReviewsSubmitting(false);
                          }}
                        >
                          {reviewsSubmitting ? "Submitting..." : "Submit review"}
                        </button>
                      </div>
                    ) : (
                      <p className="text-sm text-gray-600 mt-4">
                        You can only leave one review after donating to this cause.
                      </p>
                    )}
                  </div>

                  <div>
                    <h3 className="text-lg font-semibold text-[#3a0b2e] mb-3">
                      Reviews
                    </h3>
                    {reviews?.reviews && reviews.reviews.length > 0 ? (
                      <div className="space-y-4">
                        {reviews.reviews.map((r) => (
                          <div
                            key={r.id}
                            className="rounded-xl border border-gray-200 bg-white p-5 shadow-sm"
                          >
                            <div className="flex items-center justify-between gap-3 flex-wrap">
                              <p className="font-semibold text-[#3a0b2e]">
                                {r.user_name || "Anonymous"}
                              </p>
                              <p className="text-xs text-gray-400">
                                {r.created_at
                                  ? new Date(r.created_at).toLocaleDateString()
                                  : ""}
                              </p>
                            </div>
                            <p className="text-gray-700 mt-3 whitespace-pre-line">
                              {r.review_text}
                            </p>
                          </div>
                        ))}
                      </div>
                    ) : (
                      <p className="text-gray-600">
                        No reviews yet. Be the first to share after donating!
                      </p>
                    )}
                  </div>
                </div>
              )}
            </div>
          )}

          {activeTab === "donations" && (
            <div className="mt-6">
              {donationsLoading ? (
                <div className="space-y-4">
                  <div className="rounded-xl border border-gray-200 bg-white shadow-sm overflow-hidden">
                    <div className="h-1 bg-gradient-to-r from-[#ff6200] via-amber-300 to-[#3a0b2e] opacity-80" />
                    <div className="p-5">
                      <div className="flex items-start justify-between gap-4 flex-wrap">
                        <div className="space-y-2">
                          <div className="flex items-center gap-2">
                            <div className="h-9 w-9 rounded-lg bg-amber-50 border border-amber-100 flex items-center justify-center">
                              <LuShieldCheck className="text-[#ff6200]" />
                            </div>
                            <div>
                              <p className="text-sm font-semibold text-[#3a0b2e]">
                                Blockchain-secured donations
                              </p>
                              <p className="text-xs text-gray-500">
                                Verified on-chain • Tamper-resistant • Transparent
                              </p>
                            </div>
                          </div>
                          <p className="text-sm text-gray-600">
                            Loading the latest on-chain donations…
                          </p>
                        </div>
                        <div className="flex items-center gap-2 text-xs font-semibold text-green-700 bg-green-50 border border-green-100 px-3 py-1.5 rounded-full">
                          <span className="relative flex h-2 w-2">
                            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
                            <span className="relative inline-flex rounded-full h-2 w-2 bg-green-500"></span>
                          </span>
                          Live ledger
                        </div>
                      </div>

                      <div className="mt-5 space-y-3 animate-pulse">
                        <div className="h-16 rounded-lg bg-gray-100 border border-gray-200" />
                        <div className="h-16 rounded-lg bg-gray-100 border border-gray-200" />
                        <div className="h-16 rounded-lg bg-gray-100 border border-gray-200" />
                      </div>
                    </div>
                  </div>
                </div>
              ) : donationsError ? (
                <p className="text-red-500 text-sm">
                  Could not load donations: {donationsError}
                </p>
              ) : donations && donations.length > 0 ? (
                <div className="space-y-4">
                  <div className="rounded-xl border border-gray-200 bg-white shadow-sm overflow-hidden">
                    <div className="h-1 bg-gradient-to-r from-[#ff6200] via-amber-300 to-[#3a0b2e] opacity-80" />
                    <div className="p-5">
                      <div className="flex items-start justify-between gap-4 flex-wrap">
                        <div className="flex items-start gap-3">
                          <div className="h-10 w-10 rounded-xl bg-amber-50 border border-amber-100 flex items-center justify-center">
                            <LuShieldCheck className="text-[#ff6200] text-lg" />
                          </div>
                          <div>
                            <h3 className="text-lg font-semibold text-[#3a0b2e]">
                              On-chain Donations
                            </h3>
                            <p className="text-sm text-gray-600">
                              These records are written to the blockchain, making them secure and immutable.
                            </p>
                            <div className="mt-2 flex flex-wrap gap-2">
                              <span className="inline-flex items-center gap-1.5 text-xs font-semibold text-green-700 bg-green-50 border border-green-100 px-2.5 py-1 rounded-full">
                                <LuLock className="text-green-700" />
                                Secure
                              </span>
                              <span className="inline-flex items-center gap-1.5 text-xs font-semibold text-amber-700 bg-amber-50 border border-amber-100 px-2.5 py-1 rounded-full">
                                <LuLink2 className="text-amber-700" />
                                Verified on-chain
                              </span>
                            </div>
                          </div>
                        </div>

                        <div className="text-xs text-gray-500">
                          Showing latest{" "}
                          <span className="font-semibold text-gray-700">
                            {Math.min(10, donations.length)}
                          </span>{" "}
                          donations
                        </div>
                      </div>
                    </div>
                  </div>

                  {donations
                    .slice(-10)
                    .reverse()
                    .map((d, index) => (
                      <div
                        key={d.PaymentRef || `${d.DonorId}-${index}`}
                        className="border rounded-xl p-4 bg-white shadow-sm transition-all duration-200 hover:-translate-y-0.5 hover:shadow-md hover:border-amber-200"
                      >
                        <div className="flex items-start justify-between gap-4">
                          <div className="flex items-start gap-3">
                            <div className="mt-0.5">
                              <div className="relative h-9 w-9 rounded-xl bg-amber-50 border border-amber-100 flex items-center justify-center">
                                <LuLink2 className="text-[#ff6200]" />
                              </div>
                            </div>

                            <div>
                              <div className="flex items-center gap-2 flex-wrap">
                                <h4 className="font-semibold text-[#3a0b2e]">
                                  Donation recorded on-chain
                                </h4>
                                <span className="inline-flex items-center gap-1.5 text-[11px] font-semibold text-green-700 bg-green-50 border border-green-100 px-2 py-0.5 rounded-full">
                                  <span className="relative flex h-2 w-2">
                                    <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
                                    <span className="relative inline-flex rounded-full h-2 w-2 bg-green-500"></span>
                                  </span>
                                  Verified
                                </span>
                              </div>

                              <p className="text-sm text-gray-500 mt-1">
                                {formatDonationTimestamp(d.Timestamp) ||
                                  "Timestamp not available"}
                              </p>

                              <div className="mt-3 text-xs text-gray-500 space-y-1">
                                {d.PaymentRef && (
                                  <p>
                                    Payment Ref{" "}
                                    <span className="text-gray-400">•</span>{" "}
                                    <span className="font-mono break-all text-gray-700">
                                      {d.PaymentRef}
                                    </span>
                                  </p>
                                )}
                                {d.DonorId && (
                                  <p>
                                    Donor ID{" "}
                                    <span className="text-gray-400">•</span>{" "}
                                    <span className="font-mono break-all text-gray-700">
                                      {d.DonorId}
                                    </span>
                                  </p>
                                )}
                              </div>
                            </div>
                          </div>

                          <div className="text-right">
                            <p className="text-lg font-extrabold text-[#ff6200]">
                              {formatDonationAmount(d.Amount)}
                            </p>
                            <p className="text-xs text-gray-500 mt-1 flex items-center justify-end gap-1.5">
                              <LuLock className="text-gray-400" />
                              Immutable record
                            </p>
                          </div>
                        </div>
                      </div>
                    ))}
                </div>
              ) : (
                <div className="rounded-xl border border-gray-200 bg-white shadow-sm p-6">
                  <div className="flex items-start gap-3">
                    <div className="h-10 w-10 rounded-xl bg-amber-50 border border-amber-100 flex items-center justify-center">
                      <LuShieldCheck className="text-[#ff6200] text-lg" />
                    </div>
                    <div>
                      <h4 className="font-semibold text-[#3a0b2e]">
                        No on-chain donations yet
                      </h4>
                      <p className="text-gray-600 text-sm mt-1">
                        Once a donation is made and verified, it will appear here as a secure blockchain record.
                      </p>
                    </div>
                  </div>
                </div>
              )}
            </div>
          )}
        </div>

        {/* SECTION 2: STATS SIDEBAR */}
        <div className="bg-white rounded-xl shadow-lg p-6 border border-gray-200 h-fit sticky top-24">
          <div className="space-y-4 mb-6">
            <div className="grid grid-cols-3 text-center mb-4">
              <div>
                <p className="text-2xl font-extrabold text-[#ff6200]">
                  {formatCollected(collected, aidTypeName)}
                </p>
                <p className="text-gray-600 text-sm">{getCollectedLabel(aidTypeName)}</p>
              </div>
              <div>
                <p className="text-2xl font-extrabold text-[#3a0b2e]">
                  {aidTypeName.toLowerCase() != "volunteering" && (
                    donations ? donations.length : 0
                  )}
                </p>
                {aidTypeName.toLowerCase() != "volunteering" &&
                  <p className="text-gray-600 text-sm">
                    Donors</p>}
              </div>
              <div>
                <p className="text-2xl font-extrabold text-[#3a0b2e]">
                  {daysLeft}
                </p>
                <p className="text-gray-600 text-sm">Days Left</p>
              </div>
            </div>
            <div className="grid grid-cols-2 text-center mb-6">
              <div>
                <p className="text-2xl font-extrabold text-[#3a0b2e]">
                  {formatGoal(goal, aidTypeName)}
                </p>
                <p className="text-gray-600 text-sm">{getGoalLabel(aidTypeName)}</p>
              </div>
              <div>
                <p className="text-2xl font-extrabold text-[#3a0b2e]">
                  {percentage.toFixed(0)}%
                </p>
                <p className="text-gray-600 text-sm">Completed</p>
              </div>
            </div>
          </div>

          <div className="w-full bg-gray-200 rounded-full h-3 mb-2 overflow-hidden relative">
            {/* Collected amount bar (orange, back layer) */}
            <div
              className="h-full bg-[#ff6200] transition-all duration-500 absolute top-0 left-0"
              style={{ width: `${percentage}%` }}
            />
            {/* Disbursed amount bar (green, front layer - on top) */}
            <div
              className="h-full bg-green-500 transition-all duration-500 absolute top-0 left-0"
              style={{ width: `${disbursedPercentage}%` }}
            />
          </div>

          {aidTypeName.toLowerCase() != "volunteering" &&
            <div className="flex justify-between text-xs text-gray-600 mb-4">
              <div className="flex items-center gap-2">
                <span className="inline-block w-3 h-3 bg-green-500 rounded"></span>
                <span>{formatCollected(disbursed, aidTypeName)} disbursed</span>
              </div>
              <div className="flex items-center gap-2">
                <span className="inline-block w-3 h-3 bg-[#ff6200] rounded"></span>
                <span>{formatCollected(collected, aidTypeName)} collected</span>
              </div>
            </div>
          }

          <p className="text-gray-600 mb-6 text-sm text-center">
            {getCollectedLabel(aidTypeName)} of {formatGoal(goal, aidTypeName)} goal
          </p>

          {/* {fundingStatus !== "Fully Funded" &&
            fundingStatus !== "Closed" && (
              <Link
                to="/checkout"
                state={{ causeID: cause.id }}
                onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
              >
                <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition cursor-pointer">
                  Donate Now
                </button>
              </Link>
            )} */}

          {isOwner ? (
            <Link to={`/campaign/${cause.id}/update`} onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
              <button className="w-full bg-[#3a0b2e] hover:bg-[#6d1f57] text-white font-semibold py-3 rounded-lg transition cursor-pointer">
                Post Update
              </button>
            </Link>
          ) : (
            canDonate && (
              <Link
                to={
                  isBloodDonationAid
                    ? "/bloodDonation"
                    : isVolunteeringAid
                      ? "/volunteer"
                      : "/checkout"
                }
                state={{ causeID: cause.id }}
                onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
              >
                <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition cursor-pointer">
                  {isBloodDonationAid
                    ? "Provide details"
                    : isVolunteeringAid
                      ? "Apply Now"
                      : "Donate Now"}
                </button>
              </Link>
            )
          )}

        </div>
      </div>
    </div>
  );
};

export default CampaignPage;
