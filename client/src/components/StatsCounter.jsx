import React from "react";
import CountUp from "react-countup";
import { useInView } from "react-intersection-observer";
import img1 from "../assets/Gemini_Generated_Image_tjfuxotjfuxotjfu.png";

const statsData = [
    {
        id: 1,
        title: "Causes Accomplished",
        image: img1, // update paths based on your assets
        end: 5000,
        desc: "Empowering real change through successfully funded charity campaigns worldwide.",
    },
    {
        id: 2,
        title: "NGO's Listed",
        image: img1,
        end: 1200,
        desc: "Trusted organizations verified and connected through our transparent blockchain network.",
    },
    {
        id: 3,
        title: "Donations Tracked",
        image: img1,
        end: 50,
        desc: "Every contribution recorded securely — bringing complete visibility to where help goes.",
    },
];

const StatsCounter = () => {
    const { ref, inView } = useInView({ triggerOnce: true });

    return (
        <section
            ref={ref}
            className="flex flex-col items-center justify-center py-8"
        >
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8 w-full ">
                {statsData.map((stat) => (
                    <div
                        key={stat.id}
                        className="bg-[#f1f1f1] shadow-md rounded-2xl p-6 flex flex-col items-center text-center hover:shadow-xl transition-shadow duration-300"
                    >
                        <h3 className="text-xl font-semibold text-gray-600 mb-4">
                            {stat.title}
                        </h3>
                        <img
                            src={stat.image}
                            alt={stat.title}
                            className="w-100 h-50 mb-4 object-contain rounded-lg"
                        />

                        <h2 className="text-5xl font-bold text-gray-800 mb-2 flex">
                            {inView && <CountUp start={0} end={stat.end} duration={2.5} separator="," />} <span className="text-5xl text-[#ff6200]">+</span>
                        </h2>
                        <p className="text-gray-500">{stat.desc}</p>
                    </div>
                ))}
            </div>
        </section>
    );
};

export default StatsCounter;
