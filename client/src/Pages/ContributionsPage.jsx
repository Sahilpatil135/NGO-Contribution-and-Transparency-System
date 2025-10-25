// import React, { useState } from "react";
// import "./ContributionsPage.css";
// import ngoImage1 from '../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu (2).png';
// import ngoImage2 from '../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu (1).png';
// import ngoImage3 from '../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu.png';
// import DonateButton from "../components/DonateButton";

// const ContributionsPage = () => {
//     const ngos = [
//         {
//             name: "Helping Hands Foundation",
//             cause: "Providing education and healthcare to underprivileged children.",
//             goal: 5000,
//             raised: 3200,
//             image: ngoImage3,
//         },
//         {
//             name: "Green Earth Initiative",
//             cause: "Supporting tree plantation drives and clean energy projects.",
//             goal: 3500,
//             raised: 1500,
//             image: ngoImage2,
//         },
//         {
//             name: "Food for All",
//             cause: "Feeding homeless and needy families across cities.",
//             goal: 2000,
//             raised: 800,
//             image: ngoImage1,
//         },
//     ];

//     const [donationAmounts, setDonationAmounts] = useState(
//         ngos.map(() => "")
//     );

//     // 🪄 Function to update specific NGO donation amount
//     const handleAmountChange = (index, value) => {
//         const newAmounts = [...donationAmounts];
//         newAmounts[index] = value;
//         setDonationAmounts(newAmounts);
//     };


//     return (
//         <div className="contributions">
//             <div className="contributions-header">
//                 <h1>Make Contributions</h1>
//                 <p>Support NGOs and bring change to the community</p>
//             </div>

//             <div className="contributions-list">
//                 {ngos.map((ngo, index) => {
//                     const progress = (ngo.raised / ngo.goal) * 100;
//                     const amount = donationAmounts[index];
//                     return (
//                         <div className="ngo-card" key={index}>
//                             <div className="ngo-image">
//                                 <img src={ngo.image} alt={ngo.name} style={{ width: '100%', height: '100%', borderRadius: '8px' }} />
//                             </div>
//                             <div className="ngo-details">
//                                 <h3>{ngo.name}</h3>
//                                 <p className="ngo-cause">{ngo.cause}</p>
//                                 {/* Progress Bar Section */}
//                                 <div className="progress-container">
//                                     <div className="progress-bar" style={{ width: `${progress}%` }}></div>
//                                 </div>
//                                 <div className="progress-text">
//                                     <span>₹{ngo.raised} raised</span>
//                                     <span>Goal: ₹{ngo.goal}</span>
//                                 </div>
//                                 <div className="donation-input">
//                                     <input
//                                         type="number"
//                                         placeholder="Enter amount (₹)"
//                                         value={amount}
//                                         min="1"
//                                         onChange={(e) =>
//                                             handleAmountChange(index, e.target.value)
//                                         }
//                                     />
//                                 </div>
//                                 <DonateButton amount={Number(amount) || 0} />
//                             </div>
//                         </div>
//                     );
//                 })}
//             </div>
//         </div>
//     );
// };

// export default ContributionsPage;


import React from 'react'
import img1 from "/makeContributionbanner.jpg"
import img2 from "/ngo_helping_hand.png"
import CausesCarousel from '../components/CausesCarousel';
import Card from '../components/Card';
import { Link } from 'react-router-dom';

const ContributionsPage = () => {
    return (
        <div className="w-full">
            {/* Hero section  */}
            <section>
                <img src={img1} alt="..." />
            </section>

            {/* Types of Donations section  */}
            <section className="mt-30">
                <div className="w-11/12 flex flex-col items-center justify-between mx-auto my-6">
                    <h1 className="text-5xl font-bold mb-16">Where Do <span className="text-[#ff6200] italic">YOU</span> Want To <span className="text-[#ff6200] italic">Contribute</span> ?</h1>
                    <div className="grid grid-cols-1 md: grid-cols-5 gap-8 w-full my-8">
                        {/* will replace by clickable div's with links to respective pages */}
                        {/* <div className="flex flex-col"><img src={img2} alt=".." /> Monetary Donations</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Voluntering</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Blood & Organ Donations</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Goods & Resources Donation</div>
                        <div className="flex flex-col"><img src={img2} alt=".." /> Environmental Support</div> */}

                        <Link to="/makecontribution/monetary" className="flex flex-col items-center cursor-pointer">
                            <img src="/monetary_donations.png" alt="Monetary" className="rounded-lg" />
                            <p className="mt-2 font-semibold text-[#3a0b2e]">Monetary Donations</p>
                        </Link>

                        <Link to="/makecontribution/volunteering" className="flex flex-col items-center cursor-pointer">
                            <img src="/volunteering.png" alt="Volunteering" className="rounded-lg" />
                            <p className="mt-2 font-semibold text-[#3a0b2e]">Volunteering</p>
                        </Link>

                        <Link to="/makecontribution/blood" className="flex flex-col items-center cursor-pointer">
                            <img src="blood_organ_donations.png" alt="Blood Donation" className="rounded-lg" />
                            <p className="mt-2 font-semibold text-[#3a0b2e]">Blood & Organ Donations</p>
                        </Link>

                        <Link to="/makecontribution/goods" className="flex flex-col items-center cursor-pointer">
                            <img src="goods_resources.png" alt="Goods" className="rounded-lg" />
                            <p className="mt-2 font-semibold text-[#3a0b2e]">Goods & Resources</p>
                        </Link>

                        <Link to="/makecontribution/environment" className="flex flex-col items-center cursor-pointer">
                            <img src="environmental_support.png" alt="Environment" className="rounded-lg" />
                            <p className="mt-2 font-semibold text-[#3a0b2e]">Environmental Support</p>
                        </Link>

                    </div>
                </div>
            </section>

            {/* Explore NGO's  */}
            {/* <section className="my-16 bg-gray-200">
            <div className="w-11/12 flex flex-col justify-between mx-auto py-16">
                <h1 className="text-4xl pl-16">Explore Causes</h1>
                
            </div>
        </section> */}

            {/* Types of Causes */}
            <section>
                <CausesCarousel />
            </section>

            {/* Dynamically causes card will get displayed on selecting type of cause  */}
            <section className="w-11/12 mx-auto my-20 items-center text-center space-y-6">
                <div className="grid grid-cols-1 md:grid-cols-3 gap-12 w-full my-8">
                    <Card />
                    <Card />
                    <Card />
                </div>
            </section>
        </div>
    )
}

export default ContributionsPage
