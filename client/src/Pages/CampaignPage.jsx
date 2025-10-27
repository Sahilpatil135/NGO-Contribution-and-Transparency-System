import { useState, useEffect } from "react";
import img1 from "../../public/domains/domain_example.png";
import { Link, useParams } from "react-router-dom";
import { apiRequest, API_ENDPOINTS } from "../config/api";

const CampaignPage = () => {
  const { causeID } = useParams()

  const [loading, setLoading] = useState(true)
  const [cause, setCause] = useState({})

  useEffect(() => {
    const fetchData = async () => {
      const causeResult = await apiRequest(`${API_ENDPOINTS.GET_CAUSES}/${causeID}`)

      if (causeResult.success && causeResult.data) {
        setCause(causeResult.data)
        console.log(causeResult.data)
        setLoading(false)
      }
    }

    fetchData()
  }, [causeID])

  const getDaysLeft = (targetTimestamp) => {
    const now = Date.now(); // Current time in milliseconds
    const targetTime = new Date(targetTimestamp).getTime(); // Target time in milliseconds

    const differenceInMs = targetTime - now;

    // Convert milliseconds to days
    const millisecondsPerDay = 1000 * 60 * 60 * 24;
    const daysLeft = Math.ceil(differenceInMs / millisecondsPerDay);

    return daysLeft;
  }

  return (
    <>
      {!loading ? (
        <div className="max-w-7xl mx-auto px-4 sm:px-4 lg:px-4 py-8">
          {/* Title */}
          <div className="mb-10">
            <h1 className="text-4xl md:text-5xl font-bold mt-6 text-[#3a0b2e]">{cause.title}</h1>
            <h3 className="text-xl text-gray-500 mt-4">Campaign by <span className="font-semibold">{cause.organization.name}</span></h3>
            <span className="text-md text-gray-500 mt-1">{cause.domain.name} | {cause.aid_type.name}</span>
          </div>

          {/* Grid Layout - Left: Image + Info | Right: Stats Box */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {/* Left Section */}
            <div className="md:col-span-2">
              <img src={img1} alt={cause.title} className="w-full object-cover rounded-lg h-64 md:h-96" />

              {/* Tabs (Products, Projects, Updates) */}
              <div className="grid grid-cols-3 gap-4 my-6">
                {["Products", "Project", "Updates"].map((tab) => (
                  <div
                    key={tab}
                    className="bg-gray-100 p-4 shadow-md rounded-lg text-center border-b-2 border-black cursor-pointer hover:bg-gray-200 transition-all"
                  >
                    <h4>{tab}</h4>
                  </div>
                ))}
              </div>

              {/* About Section */}
              <h2 className="text-2xl font-semibold text-[#3a0b2e] mb-4">About the Campaign</h2>
              <p className="text-gray-700 leading-relaxed mb-6 whitespace-pre-line">{cause.description}</p>
            </div>

            {/* Right Section - Stats, Progress, Donate */}
            <div className="bg-white rounded-xl shadow-lg p-6 border border-gray-200 h-fit sticky top-24">
              {/* Stats */}
              <div className="grid grid-cols-3 text-center mb-6">
                <div>
                  <p className="text-2xl font-extrabold text-[#ff6200]">
                    ₹{parseFloat(cause.collected_amount).toLocaleString()}
                  </p>
                  <p className="text-gray-600 text-sm">Raised</p>
                </div>
                <div>
                  <p className="text-2xl font-extrabold text-[#ff6200]">{1000}+</p>
                  <p className="text-gray-600 text-sm">Donations</p>
                </div>
                <div>
                  <p className="text-2xl font-extrabold text-[#ff6200]">{getDaysLeft(cause.deadline) + 50}</p>
                  <p className="text-gray-600 text-sm">Days Left</p>
                </div>
              </div>

              {/* Progress Bar */}
              <div className="w-full bg-gray-300 rounded-full h-3 mb-6 overflow-hidden">
                <div
                  className="h-full bg-[#ff6200] transition-all duration-500"
                  style={{ width: `${Math.min(parseFloat(cause.collected_amount) / parseFloat(cause.goal_amount) * 100, 100)}%` }}
                ></div>
              </div>

              <p className="text-gray-600 mb-6 text-sm text-center">
                Raised of ₹{parseFloat(cause.goal_amount).toLocaleString()} goal
              </p>

              {/* Donate Button */}
              <Link to="/checkout" onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
                <button className="w-full bg-[#ff6200] hover:bg-[#e45a00] text-white font-semibold py-3 rounded-lg transition cursor-pointer">
                  Donate Now
                </button>
              </Link>

            </div>
          </div>

          {/* Updates Section */}
          <div className="my-12">
            <h2 className="text-2xl font-semibold text-[#3a0b2e] mb-4">Updates</h2>
            <p className="text-gray-600">No updates yet. Be the first to support the campaign!</p>
          </div>
        </div >
      ) : (
        <h1>LOADING</h1>
      )
      }
    </>
  );
};

export default CampaignPage;
