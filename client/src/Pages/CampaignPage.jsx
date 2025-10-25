import React from "react";
import img1 from "/img1.png";
import { Link } from "react-router-dom";

const campaign = {
  title: "Help Dr Amruta provide life-saving cancer treatment to poor patients who can’t afford care",
  CampaignBy: "Helping Hands Foundation",
  imageUrl: img1,
  raised: 420000,
  goal: 500000,
  donors: 1350,
  daysLeft: 12,
  description: `Dr Amruta has dedicated her career to treating cancer among under-privileged patients.  
                Your contribution will help fund treatment, medications and support for those who cannot afford it.
                Lakhs of cancer patients across the country are fighting against the dreaded disease with little support or help. Often treatment is costly which poor families can’t afford, leaving patients in a painful battle that they are bound to lose without medical care. 

Many families even sell their land and property to afford treatment, and still, end up losing their loved ones. “When such cancer patients come to us we stand by them. We hold their hands and assure them that they will get the treatment they need”, says Dr Amruta Tripathi, a compassionate oncologist at Charutar Arogya Mandal.`,
};

const CampaignPage = () => {
  const { title, CampaignBy, imageUrl, raised, goal, donors, daysLeft, description } = campaign;
  const progress = Math.min((raised / goal) * 100, 100);

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-4 lg:px-4 py-8">
      {/* Title */}
      <div className="mb-10">
        <h1 className="text-4xl md:text-5xl font-bold mt-6 text-[#3a0b2e]">{title}</h1>
        <h3 className="text-xl text-gray-500 mt-4">Campaign by <span className="font-semibold">{CampaignBy}</span></h3>
      </div>

      {/* Grid Layout - Left: Image + Info | Right: Stats Box */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {/* Left Section */}
        <div className="md:col-span-2">
          <img src={imageUrl} alt={title} className="w-full object-cover rounded-lg h-64 md:h-96" />

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
          <p className="text-gray-700 leading-relaxed mb-6 whitespace-pre-line">{description}</p>
        </div>

        {/* Right Section - Stats, Progress, Donate */}
        <div className="bg-white rounded-xl shadow-lg p-6 border border-gray-200 h-fit sticky top-24">
          {/* Stats */}
          <div className="grid grid-cols-3 text-center mb-6">
            <div>
              <p className="text-2xl font-extrabold text-[#ff6200]">
                ₹{raised.toLocaleString()}
              </p>
              <p className="text-gray-600 text-sm">Raised</p>
            </div>
            <div>
              <p className="text-2xl font-extrabold text-[#ff6200]">{donors}</p>
              <p className="text-gray-600 text-sm">Donations</p>
            </div>
            <div>
              <p className="text-2xl font-extrabold text-[#ff6200]">{daysLeft}</p>
              <p className="text-gray-600 text-sm">Days Left</p>
            </div>
          </div>

          {/* Progress Bar */}
          <div className="w-full bg-gray-300 rounded-full h-3 mb-6 overflow-hidden">
            <div
              className="h-full bg-[#ff6200] transition-all duration-500"
              style={{ width: `${progress}%` }}
            ></div>
          </div>

          <p className="text-gray-600 mb-6 text-sm text-center">
            Raised of ₹{goal.toLocaleString()} goal
          </p>

          {/* Donate Button */}
          <Link to="/checkout">
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
    </div>
  );
};

export default CampaignPage;
