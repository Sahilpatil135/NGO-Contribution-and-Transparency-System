import { Link } from "react-router-dom";

const BackLink = () => (
  <Link to="/admin" className="inline-flex w-fit rounded-md border border-slate-300 px-3 py-1.5 text-sm text-slate-700 hover:bg-slate-100">
    Back to Dashboard
  </Link>
);

const PageShell = ({ title, subtitle, children }) => (
  <main className="min-h-screen bg-slate-50 py-10">
    <div className="mx-auto w-11/12 max-w-7xl space-y-6">
      <BackLink />
      <header className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
        <h1 className="text-2xl font-bold text-slate-900">{title}</h1>
        <p className="mt-2 text-sm text-slate-600">{subtitle}</p>
      </header>
      {children}
    </div>
  </main>
);

export const NgoVerificationManagementPage = () => (
  <PageShell
    title="NGO Verification Management"
    subtitle="Review submitted organization data and attached compliance documents."
  >
    <section className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <h2 className="text-lg font-semibold text-slate-900">Pending Verifications</h2>
      <div className="mt-4 overflow-x-auto">
        <table className="min-w-full text-left text-sm">
          <thead className="border-b border-slate-200 text-slate-500">
            <tr>
              <th className="px-3 py-2">NGO</th>
              <th className="px-3 py-2">Registration</th>
              <th className="px-3 py-2">Documents</th>
              <th className="px-3 py-2">Action</th>
            </tr>
          </thead>
          <tbody className="text-slate-700">
            <tr className="border-b border-slate-100">
              <td className="px-3 py-3">Green Earth Initiative</td>
              <td className="px-3 py-3">NGO-98422</td>
              <td className="px-3 py-3">4 uploaded</td>
              <td className="px-3 py-3 space-x-2">
                <button className="rounded bg-emerald-600 px-3 py-1 text-white hover:bg-emerald-700">Approve</button>
                <button className="rounded bg-rose-600 px-3 py-1 text-white hover:bg-rose-700">Reject</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </PageShell>
);

export const DisputeResolutionPanelPage = () => (
  <PageShell
    title="Dispute Resolution Panel"
    subtitle="Investigate governance issues and mark disputes with traceable outcomes."
  >
    <section className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <h2 className="text-lg font-semibold text-slate-900">Open Cases</h2>
      <div className="mt-4 space-y-3">
        <article className="rounded-lg border border-slate-200 p-4">
          <p className="font-medium text-slate-900">Case #DSP-1048 - Milestone payout mismatch</p>
          <p className="mt-1 text-sm text-slate-600">Raised by: Health For All Trust</p>
          <div className="mt-3 flex gap-2">
            <button className="rounded bg-indigo-600 px-3 py-1 text-white hover:bg-indigo-700">View Details</button>
            <button className="rounded bg-slate-900 px-3 py-1 text-white hover:bg-slate-700">Resolve</button>
          </div>
        </article>
      </div>
    </section>
  </PageShell>
);

export const CauseActivityMonitoringPage = () => (
  <PageShell
    title="Cause & Activity Monitoring"
    subtitle="Track campaign status, funding progression, and milestone completion health."
  >
    <section className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <h2 className="text-lg font-semibold text-slate-900">Active Cause Snapshot</h2>
      <div className="mt-4 grid grid-cols-1 gap-4 md:grid-cols-3">
        <article className="rounded-lg border border-slate-200 p-4">
          <p className="text-sm text-slate-500">Cause Name</p>
          <p className="mt-1 font-semibold text-slate-900">Rural Water Access</p>
        </article>
        <article className="rounded-lg border border-slate-200 p-4">
          <p className="text-sm text-slate-500">Funding Progress</p>
          <p className="mt-1 font-semibold text-slate-900">$41,000 / $60,000</p>
        </article>
        <article className="rounded-lg border border-slate-200 p-4">
          <p className="text-sm text-slate-500">Milestone Status</p>
          <p className="mt-1 font-semibold text-slate-900">2/4 Completed</p>
        </article>
      </div>
    </section>
  </PageShell>
);

export const TrustScoreOverviewPage = () => (
  <PageShell
    title="Trust Score Overview"
    subtitle="Use transparent score breakdowns to understand NGO reliability and delivery quality."
  >
    <section className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <h2 className="text-lg font-semibold text-slate-900">Top NGO Trust Scores</h2>
      <div className="mt-4 overflow-x-auto">
        <table className="min-w-full text-left text-sm">
          <thead className="border-b border-slate-200 text-slate-500">
            <tr>
              <th className="px-3 py-2">NGO</th>
              <th className="px-3 py-2">Verification</th>
              <th className="px-3 py-2">Ratings</th>
              <th className="px-3 py-2">Completion</th>
              <th className="px-3 py-2">Final Score</th>
            </tr>
          </thead>
          <tbody className="text-slate-700">
            <tr className="border-b border-slate-100">
              <td className="px-3 py-3">Health For All Trust</td>
              <td className="px-3 py-3">40/40</td>
              <td className="px-3 py-3">28/30</td>
              <td className="px-3 py-3">25/30</td>
              <td className="px-3 py-3 font-semibold text-slate-900">93</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </PageShell>
);

export const UserNgoManagementPage = () => (
  <PageShell
    title="User & NGO Management"
    subtitle="Inspect profiles and apply governance actions to maintain platform integrity."
  >
    <section className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <h2 className="text-lg font-semibold text-slate-900">Account Actions</h2>
      <div className="mt-4 overflow-x-auto">
        <table className="min-w-full text-left text-sm">
          <thead className="border-b border-slate-200 text-slate-500">
            <tr>
              <th className="px-3 py-2">Entity</th>
              <th className="px-3 py-2">Type</th>
              <th className="px-3 py-2">Status</th>
              <th className="px-3 py-2">Actions</th>
            </tr>
          </thead>
          <tbody className="text-slate-700">
            <tr className="border-b border-slate-100">
              <td className="px-3 py-3">impact.user@ngo.org</td>
              <td className="px-3 py-3">User</td>
              <td className="px-3 py-3">Active</td>
              <td className="px-3 py-3 space-x-2">
                <button className="rounded bg-slate-900 px-3 py-1 text-white hover:bg-slate-700">Inspect</button>
                <button className="rounded bg-amber-600 px-3 py-1 text-white hover:bg-amber-700">Suspend</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>
  </PageShell>
);
