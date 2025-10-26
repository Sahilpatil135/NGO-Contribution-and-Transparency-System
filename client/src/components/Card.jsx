import React, { useState } from "react";
import heroimg from "/img1.png";
import { FaRegHeart } from "react-icons/fa";
import { BiUpvote, BiDownvote, BiGroup } from "react-icons/bi";
import { IoMdTime } from "react-icons/io";
import { Link } from "react-router-dom";

const Card = () => {
  const raised = 2000;
  const goal = 5000;
  const progress = Math.min((raised / goal) * 100, 100);

  const [hover, setHover] = useState(false);

  return (
    <div
      className="bg-white rounded-lg shadow-md overflow-hidden transition-transform hover:scale-[1.02] cursor-pointer group"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
    >
      <Link to="/campaign/1" onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
        {/* Image Section */}
        <div className="relative">
          <img
            src={heroimg}
            alt="Campaign"
            className="w-full h-56 object-cover"
          />

          {/* Heart Button */}
          <button className="absolute top-3 right-3 bg-white p-2 rounded-full shadow hover:bg-gray-100 transition cursor-pointer">
            <FaRegHeart className="text-red-500 text-lg" />
          </button>

          {/* Upvote/Downvote */}
          <div className="absolute bottom-3 right-3 flex items-center bg-white/90 px-2 py-1 rounded-full text-gray-700 text-sm shadow">
            <button className="hover:text-green-600 cursor-pointer">
              <BiUpvote className="text-lg" />
            </button>
            <span className="mx-1">55 |</span>
            <button className="hover:text-red-600 cursor-pointer">
              <BiDownvote className="text-lg" />
            </button>
          </div>
        </div>

        {/* Text and Progress Section */}
        <div className="px-6 py-4 text-left">
          <h1 className="text-xl font-bold text-gray-800">Helping Hands Foundation</h1>
          <p className="text-md text-gray-500 mb-3">by Campaign Name</p>

          {/* Donations & Time Left */}
          <div className="flex justify-between text-sm text-gray-600 mb-2">
            <p className="flex"><BiGroup className="h-5 mr-1 text-lg" />100 Donations</p>
            <p className="flex"><IoMdTime className="h-5 text-lg mr-1" />3 Days Left</p>
          </div>

          {/* Hover Logic: hide details, show donate button */}
          {!hover ? (
            <div>
              {/* Progress Bar */}
              <div className="w-full bg-gray-200 h-2 rounded-full overflow-hidden mb-2">
                <div
                  className="h-full bg-green-500 transition-all duration-500"
                  style={{ width: `${progress}%` }}
                ></div>
              </div>

              {/* Raised vs Goal */}
              <div className="flex justify-between text-sm text-gray-600">
                <span> <span className="text-lg font-bold text-black">₹{raised.toLocaleString()}</span> Raised</span>
                <span>Goal: <span className="text-red-500 font-bold">₹{goal.toLocaleString()}</span></span>
              </div>
            </div>
          ) : (
            <div className="flex justify-center items-center h-12">
              <Link to="/checkout" className="w-full">
                <button className="bg-green-600 hover:bg-green-700 w-full text-white font-semibold py-2 px-6 mx-2 rounded-full transition cursor-pointer">
                  Donate
                </button>
              </Link>
            </div>
          )}
        </div>
      </Link>
    </div>
  );
};

export default Card;
