import React from "react";
import { Link, useLocation } from "react-router-dom";

const DonationSuccess = () => {
  const location = useLocation();
  const { amount, causeId, paymentId } = location.state || {};

  return (
    <div className="bg-gray-100 py-12 min-h-[70vh]">
      <div className="max-w-3xl mx-auto px-4">
        <div className="bg-white rounded-xl shadow-lg border border-gray-200 p-8">
          <div className="flex items-start justify-between gap-6 flex-wrap">
            <div>
              <p className="text-sm font-semibold text-[#ff6200]">
                Payment confirmed
              </p>
              <h1 className="text-3xl md:text-4xl font-bold text-[#3a0b2e] mt-2">
                Thank you for your donation
              </h1>
              <p className="text-gray-600 mt-3">
                Your contribution helps make this campaign possible.
              </p>
            </div>

            <div className="rounded-lg bg-amber-50 border border-amber-100 px-5 py-4 min-w-[220px]">
              <p className="text-xs uppercase tracking-wide text-amber-700 font-semibold">
                Amount
              </p>
              <p className="text-2xl font-extrabold text-[#ff6200] mt-1">
                {amount != null && amount !== "" ? `₹${amount}` : "—"}
              </p>
            </div>
          </div>

          <div className="mt-6 border-t border-gray-200 pt-6 space-y-3">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2">
              <p className="text-sm text-gray-600">Payment ID</p>
              <p className="text-sm font-mono text-gray-800 break-all">
                {paymentId || "—"}
              </p>
            </div>
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2">
              <p className="text-sm text-gray-600">Cause</p>
              {causeId ? (
                <Link
                  to={`/campaign/${causeId}`}
                  className="text-sm font-semibold text-[#3a0b2e] hover:text-[#ff6200] transition break-all"
                >
                  View campaign
                </Link>
              ) : (
                <p className="text-sm text-gray-800">—</p>
              )}
            </div>
          </div>

          <div className="mt-8 flex flex-col sm:flex-row gap-3">
            {causeId && (
              <Link to={`/campaign/${causeId}`} className="w-full sm:w-auto">
                <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold px-5 py-3 rounded-lg transition cursor-pointer">
                  Back to campaign
                </button>
              </Link>
            )}
            <Link to="/makeContribution" className="w-full sm:w-auto">
              <button className="w-full bg-white border border-gray-300 hover:border-[#ff6200] text-[#3a0b2e] font-semibold px-5 py-3 rounded-lg transition cursor-pointer">
                Explore more causes
              </button>
            </Link>
          </div>
        </div>

        <p className="text-xs text-gray-500 mt-4 text-center">
          If you have any issues, please contact support with your Payment ID.
        </p>
      </div>
    </div>
  );
};

export default DonationSuccess;

