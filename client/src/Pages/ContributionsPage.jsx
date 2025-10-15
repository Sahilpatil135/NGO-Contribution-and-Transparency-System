import React from "react";
import "./ContributionsPage.css";
import ngoImage1 from '../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu (2).png';
import ngoImage2 from '../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu (1).png';
import ngoImage3 from '../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu.png';
import DonateButton from "../components/DonateButton";

const ContributionsPage = () => {
    const ngos = [
        {
            name: "Helping Hands Foundation",
            cause: "Providing education and healthcare to underprivileged children.",
            goal: 5000,
            raised: 3200,
            image: ngoImage3,
        },
        {
            name: "Green Earth Initiative",
            cause: "Supporting tree plantation drives and clean energy projects.",
            goal: 3500,
            raised: 1500,
            image: ngoImage2,
        },
        {
            name: "Food for All",
            cause: "Feeding homeless and needy families across cities.",
            goal: 2000,
            raised: 800,
            image: ngoImage1,
        },
    ];

    return (
        <div className="contributions">
            <div className="contributions-header">
                <h1>Make Contributions</h1>
                <p>Support NGOs and bring change to the community</p>
            </div>

            <div className="contributions-list">
                {ngos.map((ngo, index) => {
                    const progress = (ngo.raised / ngo.goal) * 100;
                    return (
                        <div className="ngo-card" key={index}>
                            <div className="ngo-image">
                                <img src={ngo.image} alt={ngo.name} style={{ width: '100%', height: '100%', borderRadius: '8px' }} />
                            </div>
                            <div className="ngo-details">
                                <h3>{ngo.name}</h3>
                                <p className="ngo-cause">{ngo.cause}</p>
                                {/* Progress Bar Section */}
                                <div className="progress-container">
                                    <div className="progress-bar" style={{ width: `${progress}%` }}></div>
                                </div>
                                <div className="progress-text">
                                    <span>₹{ngo.raised} raised</span>
                                    <span>Goal: ₹{ngo.goal}</span>
                                </div>                                                                
                                <DonateButton amount={10} />
                            </div>
                        </div>
                    );
                })}
            </div>
        </div>
    );
};

export default ContributionsPage;
