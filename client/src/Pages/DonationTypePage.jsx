import { useParams } from "react-router-dom";
import monetaryBanner from "/makeContributionbanner.jpg";
import volunteerBanner from "/makeContributionbanner.jpg";
import bloodBanner from "/makeContributionbanner.jpg";
import goodsBanner from "/makeContributionbanner.jpg";
import environmentBanner from "/makeContributionbanner.jpg";
import Card from "../components/Card";

const banners = {
  monetary: monetaryBanner,
  volunteering: volunteerBanner,
  blood: bloodBanner,
  goods: goodsBanner,
  environment: environmentBanner,
};

const titles = {
  monetary: "Monetary Donations",
  volunteering: "Volunteering Opportunities",
  blood: "Blood & Organ Donations",
  goods: "Goods & Resources Support",
  environment: "Environmental Support Initiatives",
};

const DonationTypePage = () => {
  const { id } = useParams();
  const banner = banners[id] || monetaryBanner;
  const title = titles[id] || "Donation Type";

  return (
    <div className="max-w-7xl mx-auto py-12 px-4">
      <img src={banner} alt={title} className="w-full h-64 object-cover rounded-lg mb-12" />
      <h1 className="text-center text-4xl font-bold text-[#3a0b2e] mb-6">{title}</h1>
      <p className="text-center text-gray-700 mb-12">
        Explore campaigns, projects, and opportunities related to {title.toLowerCase()}.
      </p>

      {/* Replace below with dynamic cards or projects */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <Card />
        <Card />
        <Card />
      </div>
    </div>
  );
};

export default DonationTypePage;
