import { useEffect, useMemo, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { apiRequest, API_ENDPOINTS } from "../../config/api";
import { useAuth } from "../../contexts/AuthContext";
import { getCauseImage } from "../../utils/imageHelper";
import causePlaceholder from "../../../public/domains/domain_example.png";

const OrganizationAccountsPage = () => {
  const { user, organization, isLoading, fetchCurrentOrganization } = useAuth();
  const { organizationId } = useParams();

  const [causes, setCauses] = useState([]);
  const [loadingCauses, setLoadingCauses] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // If we're on the "my organization" route, ensure we have org details.
    if (!organizationId && user?.role === "organization" && !organization) {
      fetchCurrentOrganization();
    }
  }, [organizationId, user?.role, organization, fetchCurrentOrganization]);

  const organizationIdToFetch = organizationId || organization?.id;

  const computeFundingStatus = (cause) => {
    const collected = Number(cause?.collected_amount) || 0;
    const rawGoal = cause?.goal_amount;
    const goal = rawGoal != null ? parseFloat(rawGoal) : 0;
    const hasValidGoal = rawGoal != null && !Number.isNaN(goal) && goal > 0;

    const deadline = cause?.deadline ? new Date(cause.deadline) : null;
    const now = new Date();

    // Matches the backend logic (computeFundingStatus in models/cause.go).
    if (!hasValidGoal) {
      if (collected <= 0) return "Not Started";
      return "Active";
    }

    if (collected <= 0) {
      if (deadline && deadline < now) return "Closed";
      return "Not Started";
    }

    if (collected >= goal) return "Fully Funded";

    if (deadline && deadline < now) return "Closed";
    return "Active";
  };

  const getDaysLeft = (deadline) => {
    if (!deadline) return null;
    const targetTime = new Date(deadline).getTime();
    if (Number.isNaN(targetTime)) return null;
    const diffMs = targetTime - Date.now();
    const msPerDay = 1000 * 60 * 60 * 24;
    return Math.max(0, Math.ceil(diffMs / msPerDay));
  };

  useEffect(() => {
    const fetchCauses = async () => {
      if (!organizationIdToFetch) return;

      setLoadingCauses(true);
      setError(null);
      try {
        const res = await apiRequest(
          `${API_ENDPOINTS.GET_CAUSES_BY_ORGANIZATION}/${organizationIdToFetch}`
        );

        if (res.success && Array.isArray(res.data)) {
          setCauses(res.data);
        } else {
          setCauses([]);
          setError(res.error || "Failed to load causes");
        }
      } catch (e) {
        setCauses([]);
        setError(e?.message || "Failed to load causes");
      } finally {
        setLoadingCauses(false);
      }
    };

    fetchCauses();
  }, [organizationIdToFetch]);

  const sortedCauses = useMemo(() => {
    return [...causes].sort((a, b) => {
      const at = a?.created_at ? new Date(a.created_at).getTime() : 0;
      const bt = b?.created_at ? new Date(b.created_at).getTime() : 0;
      return bt - at;
    });
  }, [causes]);

  const fundingChipClass = (status) => {
    if (status === "Fully Funded") return "bg-green-100 text-green-800";
    if (status === "Closed") return "bg-gray-200 text-gray-700";
    if (status === "Active") return "bg-amber-100 text-amber-800";
    return "bg-slate-100 text-slate-700";
  };

  const canViewAsMyOrg =
    !!organizationId || (user?.role === "organization" && organization?.id);

  const organizationName =
    causes?.[0]?.organization?.name || organization?.organization_name || "Organization";

  if (isLoading && !organizationIdToFetch) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-12 text-center">
        <p className="text-gray-600">Loading...</p>
      </div>
    );
  }

  if (!canViewAsMyOrg) {
    return (
      <div className="max-w-7xl mx-auto px-4 py-12 text-center">
        <p className="text-gray-600">
          Organization causes are only available to organization accounts.
        </p>
        <Link
          to="/profile"
          className="inline-block mt-4 text-[#ff6200] hover:text-[#e45a00] font-semibold"
        >
          Back to profile →
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-6">
        <h1 className="text-3xl font-bold text-[#3a0b2e]">Organization accounts</h1>
        <p className="text-gray-600 mt-2">
          Showing causes for{" "}
          <span className="font-semibold text-gray-800">{organizationName}</span>
        </p>
      </div>

      {/* Organization details */}
      <div className="mb-6 rounded-xl border border-gray-200 bg-white p-5 shadow-sm">
        <div className="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-4">
          <div>
            <h2 className="text-xl font-semibold text-[#3a0b2e]">
              {organization?.organization_name || organizationName}
            </h2>
            {organization?.is_approved != null && (
              <p className="text-sm text-gray-500 mt-1">
                Status:{" "}
                <span className="font-semibold">
                  {organization.is_approved ? "Approved" : "Pending"}
                </span>
              </p>
            )}
          </div>

          <div className="text-sm text-gray-600">
            {organization?.registration_number && (
              <p>
                Registration:{" "}
                <span className="font-semibold text-gray-800">
                  {organization.registration_number}
                </span>
              </p>
            )}
            {organization?.organization_type && (
              <p>
                Type:{" "}
                <span className="font-semibold text-gray-800">
                  {organization.organization_type}
                </span>
              </p>
            )}
          </div>
        </div>

        {organization?.about && (
          <div className="mt-4">
            <p className="text-sm font-semibold text-gray-800 mb-1">
              About
            </p>
            <p className="text-gray-700 whitespace-pre-line">
              {organization.about}
            </p>
          </div>
        )}

        {(organization?.website_url || organization?.address) && (
          <div className="mt-4 grid grid-cols-1 md:grid-cols-2 gap-4">
            {organization?.website_url && (
              <div>
                <p className="text-sm font-semibold text-gray-800 mb-1">
                  Website
                </p>
                <a
                  href={organization.website_url}
                  target="_blank"
                  rel="noreferrer"
                  className="text-[#ff6200] hover:text-[#e45a00] font-semibold break-all"
                >
                  {organization.website_url}
                </a>
              </div>
            )}
            {organization?.address && (
              <div>
                <p className="text-sm font-semibold text-gray-800 mb-1">
                  Address
                </p>
                <p className="text-gray-700 whitespace-pre-line">
                  {organization.address}
                </p>
              </div>
            )}
          </div>
        )}
      </div>

      {loadingCauses ? (
        <div className="space-y-4">
          <div className="rounded-xl border border-gray-200 bg-white p-4 animate-pulse">
            <div className="h-20 bg-gray-100 rounded-lg" />
          </div>
          <div className="rounded-xl border border-gray-200 bg-white p-4 animate-pulse">
            <div className="h-20 bg-gray-100 rounded-lg" />
          </div>
          <div className="rounded-xl border border-gray-200 bg-white p-4 animate-pulse">
            <div className="h-20 bg-gray-100 rounded-lg" />
          </div>
        </div>
      ) : error ? (
        <div className="rounded-xl border border-red-200 bg-red-50 p-4 text-red-700">
          {error}
        </div>
      ) : sortedCauses.length === 0 ? (
        <div className="rounded-xl border border-gray-200 bg-white p-8 text-center">
          <p className="text-gray-600">No causes found for this organization.</p>
          {!organizationId && user?.role === "organization" && (
            <Link
              to="/createCampaign"
              className="inline-block mt-4 text-[#ff6200] hover:text-[#e45a00] font-semibold"
            >
              Create your first campaign →
            </Link>
          )}
        </div>
      ) : (
        <div className="space-y-4">
          {sortedCauses.map((cause) => {
            const goal = parseFloat(cause?.goal_amount) || 0;
            const collected = parseFloat(cause?.collected_amount) || 0;
            const percentage = goal > 0 ? Math.min((collected / goal) * 100, 100) : 0;
            const fundingStatus = computeFundingStatus(cause);
            const daysLeft = getDaysLeft(cause?.deadline);

            return (
              <div
                key={cause.id}
                className="rounded-xl border border-gray-200 bg-white overflow-hidden shadow-sm hover:shadow-md transition"
              >
                <div className="p-5 flex flex-col md:flex-row gap-4 md:items-center">
                  <div className="w-full md:w-48 flex-shrink-0">
                    <img
                      src={getCauseImage(cause?.cover_image_url, causePlaceholder)}
                      alt={cause?.title || "Campaign"}
                      className="w-full h-32 md:h-28 rounded-lg object-cover"
                    />
                  </div>

                  <div className="flex-1 min-w-0">
                    <div className="flex flex-wrap items-center gap-3">
                      <h2 className="text-xl font-semibold text-[#3a0b2e] line-clamp-2">
                        {cause?.title || "Untitled campaign"}
                      </h2>
                      <span
                        className={`inline-flex items-center px-3 py-1 rounded-full text-sm font-medium ${fundingChipClass(
                          fundingStatus
                        )}`}
                      >
                        {fundingStatus}
                      </span>
                    </div>

                    <div className="text-sm text-gray-500 mt-1">
                      {cause?.domain?.name ? `${cause.domain.name} • ` : ""}
                      {cause?.aid_type?.name ? `${cause.aid_type.name}` : ""}
                    </div>

                    <div className="mt-3">
                      <div className="text-xs uppercase tracking-wide text-gray-500 font-semibold mb-2">
                        Progress
                      </div>
                      <div className="w-full bg-gray-200 rounded-full h-3 overflow-hidden">
                        <div
                          className="h-full bg-[#ff6200] transition-all duration-500"
                          style={{ width: `${percentage}%` }}
                        />
                      </div>
                      <div className="flex flex-wrap items-center gap-x-4 gap-y-1 text-sm text-gray-600 mt-2">
                        <span>
                          ₹{collected.toLocaleString("en-IN")} raised of{" "}
                          ₹{goal.toLocaleString("en-IN")} goal
                        </span>
                        {daysLeft != null && cause?.deadline && (
                          <span>• {daysLeft} days left</span>
                        )}
                      </div>
                    </div>
                  </div>

                  <div className="flex gap-3 md:flex-col md:items-stretch">
                    <Link
                      to={`/campaign/${cause.id}`}
                      className="inline-flex items-center justify-center bg-[#3a0b2e] hover:bg-[#6d1f57] text-white font-semibold px-4 py-2 rounded-lg transition cursor-pointer text-sm"
                      onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
                    >
                      View
                    </Link>
                    {user?.role === "organization" &&
                      organization?.id &&
                      cause?.organization?.id &&
                      String(organization.id) === String(cause.organization.id) && (
                        <Link
                          to={`/campaign/${cause.id}/update`}
                          className="inline-flex items-center justify-center bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-4 py-2 rounded-lg transition cursor-pointer text-sm"
                          onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
                        >
                          Post update
                        </Link>
                      )}
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      )}
    </div>
  );
};

export default OrganizationAccountsPage;

