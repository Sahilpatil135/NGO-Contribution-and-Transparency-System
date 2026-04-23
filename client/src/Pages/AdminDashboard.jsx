import { useEffect, useMemo, useState } from "react";
import { Link } from "react-router-dom";
import { API_ENDPOINTS, apiRequest } from "../config/api";

const governanceCards = [
  {
    title: "NGO Verification Management",
    description:
      "Review organization identity, documents, and registration details. Approve or reject NGO applications.",
    path: "/admin/ngo-verifications",
    cta: "Review Verifications",
  },
  {
    title: "Dispute Resolution Panel",
    description:
      "Inspect reported cases, gather issue context, and resolve NGO-raised disputes with clear outcomes.",
    path: "/admin/disputes",
    cta: "Open Disputes",
  },
  {
    title: "Cause & Activity Monitoring",
    description:
      "Track active causes, funding progress, and milestone completion across all running initiatives.",
    path: "/admin/causes-monitoring",
    cta: "Monitor Causes",
  },
  {
    title: "Trust Score Overview",
    description:
      "View NGO trust scores and understand weighted breakdowns from verification, ratings, and completion.",
    path: "/admin/trust-scores",
    cta: "View Trust Scores",
  },
  {
    title: "User & NGO Management",
    description:
      "Search user or NGO profiles and apply governance actions like inspect, suspend, or restore access.",
    path: "/admin/user-ngo-management",
    cta: "Manage Accounts",
  },
];

const PREVIEW_LIMIT = 5;

const AdminDashboard = () => {
  const [stats, setStats] = useState({
    total_organizations: 0,
    total_donors: 0,
    total_causes: 0,
    total_donations: 0,
  });
  const [organizationNames, setOrganizationNames] = useState([]);
  const [donorNames, setDonorNames] = useState([]);
  const [showAllOrganizations, setShowAllOrganizations] = useState(false);
  const [showAllDonors, setShowAllDonors] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const loadDashboard = async () => {
      setLoading(true);
      setError("");

      const result = await apiRequest(API_ENDPOINTS.GET_ADMIN_DASHBOARD);
      if (!result.success) {
        setError(result.error || "Unable to load dashboard data.");
        setLoading(false);
        return;
      }

      const data = result.data || {};
      const fetchedStats = data.stats || {};

      setStats({
        total_organizations: fetchedStats.total_organizations || 0,
        total_donors: fetchedStats.total_donors || 0,
        total_causes: fetchedStats.total_causes || 0,
        total_donations: fetchedStats.total_donations || 0,
      });
      setOrganizationNames(data.organization_names || []);
      setDonorNames(data.donor_names || []);
      setLoading(false);
    };

    loadDashboard();
  }, []);

  const metrics = useMemo(
    () => [
      { label: "Total Organizations", value: stats.total_organizations, hint: "Registered organizations" },
      { label: "Total Donors", value: stats.total_donors, hint: "Registered donor accounts" },
      { label: "Total Causes", value: stats.total_causes, hint: "Created causes on platform" },
      { label: "Number of Donations", value: stats.total_donations, hint: "Total donation records" },
    ],
    [stats]
  );

  const visibleOrganizations = showAllOrganizations
    ? organizationNames
    : organizationNames.slice(0, PREVIEW_LIMIT);
  const visibleDonors = showAllDonors ? donorNames : donorNames.slice(0, PREVIEW_LIMIT);

  return (
    <main className="min-h-screen bg-slate-50 py-10">
      <div className="mx-auto w-11/12 max-w-7xl">
        <header className="mb-8 rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
          <p className="text-sm font-medium uppercase tracking-wide text-indigo-600">Admin Governance</p>
          <h1 className="mt-2 text-3xl font-bold text-slate-900">Transparency Control Center</h1>
          <p className="mt-2 max-w-3xl text-sm text-slate-600">
            Lightweight oversight dashboard for monitoring platform health, resolving governance issues, and
            taking action quickly.
          </p>
        </header>

        <section>
          <h2 className="mb-4 text-lg font-semibold text-slate-900">Platform Metrics</h2>
          {loading ? (
            <p className="rounded-xl border border-slate-200 bg-white p-5 text-sm text-slate-600 shadow-sm">
              Loading dashboard metrics...
            </p>
          ) : (
            <div className="grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-4">
              {metrics.map((metric) => (
                <article key={metric.label} className="rounded-xl border border-slate-200 bg-white p-5 shadow-sm">
                  <p className="text-sm font-medium text-slate-500">{metric.label}</p>
                  <p className="mt-2 text-3xl font-bold text-slate-900">{metric.value}</p>
                  <p className="mt-2 text-sm text-slate-600">{metric.hint}</p>
                </article>
              ))}
            </div>
          )}
          {error ? <p className="mt-3 text-sm text-red-600">{error}</p> : null}
        </section>

        <section className="mt-10 grid grid-cols-1 gap-4 lg:grid-cols-2">
          <article className="rounded-xl border border-slate-200 bg-white p-5 shadow-sm">
            <div className="flex items-center justify-between gap-2">
              <h2 className="text-base font-semibold text-slate-900">Organizations</h2>
              {organizationNames.length > PREVIEW_LIMIT ? (
                <button
                  type="button"
                  onClick={() => setShowAllOrganizations((prev) => !prev)}
                  className="text-sm font-medium text-indigo-600 hover:text-indigo-800"
                >
                  {showAllOrganizations ? "View Less" : "View More"}
                </button>
              ) : null}
            </div>
            <ul className="mt-3 space-y-2 text-sm text-slate-700">
              {visibleOrganizations.map((name, index) => (
                <li key={`${name}-${index}`} className="rounded-md bg-slate-50 px-3 py-2">
                  {name}
                </li>
              ))}
              {!loading && visibleOrganizations.length === 0 ? <li>No organizations found.</li> : null}
            </ul>
          </article>

          <article className="rounded-xl border border-slate-200 bg-white p-5 shadow-sm">
            <div className="flex items-center justify-between gap-2">
              <h2 className="text-base font-semibold text-slate-900">Donors</h2>
              {donorNames.length > PREVIEW_LIMIT ? (
                <button
                  type="button"
                  onClick={() => setShowAllDonors((prev) => !prev)}
                  className="text-sm font-medium text-indigo-600 hover:text-indigo-800"
                >
                  {showAllDonors ? "View Less" : "View More"}
                </button>
              ) : null}
            </div>
            <ul className="mt-3 space-y-2 text-sm text-slate-700">
              {visibleDonors.map((name, index) => (
                <li key={`${name}-${index}`} className="rounded-md bg-slate-50 px-3 py-2">
                  {name}
                </li>
              ))}
              {!loading && visibleDonors.length === 0 ? <li>No donors found.</li> : null}
            </ul>
          </article>
        </section>

        <section className="mt-10">
          <div className="mb-4 flex items-center justify-between gap-2">
            <h2 className="text-lg font-semibold text-slate-900">Governance Panels</h2>
            <span className="rounded-full bg-indigo-100 px-3 py-1 text-xs font-medium text-indigo-700">
              Quick Actions
            </span>
          </div>
          <div className="grid grid-cols-1 gap-4 lg:grid-cols-2">
            {governanceCards.map((card) => (
              <article
                key={card.title}
                className="flex h-full flex-col justify-between rounded-xl border border-slate-200 bg-white p-5 shadow-sm"
              >
                <div>
                  <h3 className="text-base font-semibold text-slate-900">{card.title}</h3>
                  <p className="mt-2 text-sm leading-relaxed text-slate-600">{card.description}</p>
                </div>
                <Link
                  to={card.path}
                  className="mt-5 inline-flex w-fit items-center rounded-md bg-slate-900 px-4 py-2 text-sm font-medium text-white transition hover:bg-slate-700"
                >
                  {card.cta}
                </Link>
              </article>
            ))}
          </div>
        </section>
      </div>
    </main>
  );
};

export default AdminDashboard;
