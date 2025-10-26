import React from 'react'
import { useAuth } from '../contexts/AuthContext';
import { Link } from "react-router-dom";
import heroimg from "/img1.png"
import StatsCounter from '../components/StatsCounter';
import Card from '../components/Card';

const HomePage = () => {
    const { user } = useAuth();

    if (!user) {
        return null;
    }
    console.log(user)

    return (
        <div className="w-full">
            {/* Hero Section */}
            <section className="w-11/12 flex flex-col md:flex-row items-center justify-between bg-[#e3e5cf] mx-auto my-8 rounded-lg p-12">
                {/* Left Text */}
                <div className="md:w-1/2 w-full space-y-6">
                    <h1 className="text-5xl md:text-6xl font-bold text-[#3a0b2e] leading-tight">
                        Your <span className="text-[#ff6200] italic">Support</span>
                        <br />
                        Can Transform
                        <br />
                        Communities
                    </h1>

                    <div>
                        <h2 className="text-2xl font-semibold text-[#3a0b2e] mb-1">5000+</h2>
                        <p className="text-base text-[#3a0b2e]">People donated</p>
                    </div>

                    <p className="text-base text-[#3a0b2e] leading-relaxed md:w-4/5">
                        Empowering communities, transforming lives. Your support can bring
                        hope to those in need.
                    </p>

                    <div className="flex gap-4 pt-2">
                        <Link to="/makeContribution" className="no-underline" onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
                            <button className="bg-[#ff6200] text-white text-sm px-6 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer">
                                Start Donating
                            </button>
                        </Link>
                        <Link to="/makecontribution/volunteering" className="no-underline" onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
                            <button className="bg-transparent border border-[#3a0b2e] text-[#3a0b2e] text-sm px-6 py-2 rounded-md transition hover:bg-[#3a0b2e] hover:text-white cursor-pointer">
                                Join as Volunteer
                            </button>
                        </Link>

                    </div>
                </div>

                {/* Right Image */}
                <div className="md:w-1/3 w-full mt-8 md:mt-0">
                    <img
                        src={heroimg}
                        alt="volunteer"
                        className="w-full h-[400px] object-cover rounded-lg shadow-md"
                    />
                </div>
            </section>

            {/* Stats Section  */}
            <section>
                <div className="w-11/12 mx-auto my-16 items-center text-center space-y-6">
                    <h1 className="text-6xl font-bold">
                        Our <span className="text-[#ff6200] italic">Impact</span> in Action
                    </h1>
                    <p className="text-gray-500 max-w-lg mx-auto my-0 leading-relaxed">CharityLight brings transparency to giving — every donation is tracked, building trust and impact.</p>
                    <StatsCounter />
                </div>
            </section>

            {/* Objectives of CharityLight & Safety using Blockchain (TODO) */}
            <section className="w-11/12 mx-auto mt-20 rounded-lg bg-[#3a0b2e] py-16 px-6 md:px-12">
                <div className="max-w-7xl mx-auto flex flex-col md:flex-row items-center gap-12">

                    {/* Left Image */}
                    <div className="md:w-1/2 w-full">
                        <img
                            src="/img1.png"
                            alt="Our Mission"
                            className="w-full h-[380px] object-cover rounded-2xl shadow-lg"
                        />
                    </div>

                    {/* Right Content */}
                    <div className="md:w-1/2 w-full text-white">
                        <h2 className="text-4xl font-bold mb-6 leading-snug">
                            Building <span className="text-[#ff6200]">Trust</span> Through Transparency
                        </h2>

                        <p className="text-gray-200 mb-6 text-base leading-relaxed">
                            We’re redefining how donations work — by making every contribution
                            <span className="text-[#ff6200] font-semibold"> transparent, verifiable, </span>
                            and truly impactful. Our platform connects compassionate donors with trusted NGOs
                            using the power of <span className="font-semibold">blockchain and smart contracts.</span>
                        </p>

                        <ul className="space-y-4 text-gray-100 text-sm md:text-base">
                            <li className="flex items-start gap-3">
                                <span className="text-[#ff6200] text-xl">✓</span>
                                <span>Empowering NGOs with a decentralized, transparent donation system.</span>
                            </li>
                            <li className="flex items-start gap-3">
                                <span className="text-[#ff6200] text-xl">✓</span>
                                <span>Supporting diverse contributions — from funds to time, blood, and organ pledges.</span>
                            </li>
                            <li className="flex items-start gap-3">
                                <span className="text-[#ff6200] text-xl">✓</span>
                                <span>Ensuring authenticity with image forensics and metadata-based verification.</span>
                            </li>
                            <li className="flex items-start gap-3">
                                <span className="text-[#ff6200] text-xl">✓</span>
                                <span>Rewarding trust through verifiable receipts and community-driven validation.</span>
                            </li>
                        </ul>

                        <Link to="/about" className="no-underline">
                            <button className="mt-8 bg-[#ff6200] text-white font-semibold px-6 py-3 rounded-md hover:bg-[#e45a00] transition cursor-pointer">
                                Learn More
                            </button>
                        </Link>
                    </div>
                </div>
            </section>

            {/* Causes Section  */}
            <section className="w-11/12 mx-auto my-16 mt-20 items-center text-center space-y-6">
                <h1 className="text-6xl font-bold">
                    Featured <span className="text-[#ff6200] italic">Causes</span>
                </h1>
                <p className="text-gray-500 max-w-lg mx-auto my-0 leading-relaxed">Discover impactful charity campaigns verified on our transparent blockchain network.</p>
                <div className="grid grid-cols-1 md:grid-cols-3 gap-12 w-full my-8">
                    <Card />
                    <Card />
                    <Card />
                </div>
            </section>

        </div>
    )
}

export default HomePage
