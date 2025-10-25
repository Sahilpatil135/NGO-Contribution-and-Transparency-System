import React from 'react'
import { useAuth } from '../contexts/AuthContext';
import { Link } from "react-router-dom";
import heroimg from "../assets/img1.png"
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
                        <Link to="/makeContribution" className="no-underline">
                            <button className="bg-[#ff6200] text-white text-sm px-6 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer">
                                Donate Now
                            </button>
                        </Link>
                        <button className="bg-transparent border border-[#3a0b2e] text-[#3a0b2e] text-sm px-6 py-2 rounded-md transition hover:bg-[#3a0b2e] hover:text-white">
                            Get Involved
                        </button>
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

            {/* Causes Section  */}
            <section className="w-11/12 mx-auto my-16 items-center text-center space-y-6">
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
