// import React from 'react'
// import heroimg from "../assets/img1.png"
// import { FaRegHeart } from "react-icons/fa";
// import { BiUpvote } from "react-icons/bi";
// import { BiDownvote } from "react-icons/bi";
// import "./Card.css"

// const Card = () => {
//     return (
//         <div>
//             <div>
//                 <img src={heroimg} alt="..." />
//                 <button><FaRegHeart /></button>
//                 <div className="flex ">
//                     <button><BiUpvote /></button> 55 | <button><BiDownvote /></button>
//                 </div>
//             </div>
//             <div>
//                 <h1>Title</h1>
//                 <p>by Campaign name</p>
//                 <div className="flex flex-col">
//                     <div className="flex justify-between">
//                         <p>No. of Donations</p>
//                         <p>Time left</p>
//                     </div>
//                     <div className="progress-container">
//                         <div className="progress-bar"></div>
//                     </div>
//                     <div className="progress-text">
//                         <span>₹{1000} raised</span>
//                         <span>Goal: ₹{5000}</span>
//                     </div>
//                 </div>
//                 <button>Donate</button>
//             </div>
//         </div>
//     )
// }

// export default Card


import React, { useState } from "react";
import heroimg from "../assets/img1.png";
import { FaRegHeart } from "react-icons/fa";
import { BiUpvote, BiDownvote } from "react-icons/bi";

const Card = () => {
  const raised = 1000;
  const goal = 5000;
  const progress = Math.min((raised / goal) * 100, 100);

  const [hover, setHover] = useState(false);

  return (
    <div
      className="bg-white rounded-lg shadow-md overflow-hidden transition-transform hover:scale-[1.02] cursor-pointer group"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
    >
      {/* Image Section */}
      <div className="relative">
        <img
          src={heroimg}
          alt="Campaign"
          className="w-full h-56 object-cover"
        />

        {/* Heart Button */ }
        <button className="absolute top-3 right-3 bg-white p-2 rounded-full shadow hover:bg-gray-100 transition">
          <FaRegHeart className="text-red-500 text-lg" />
        </button>

        {/* Upvote/Downvote */}
        <div className="absolute bottom-3 right-3 flex items-center bg-white/90 px-2 py-1 rounded-full text-gray-700 text-sm shadow">
          <button className="hover:text-green-600">
            <BiUpvote className="text-lg" />
          </button>
          <span className="mx-1">55</span>
          <button className="hover:text-red-600">
            <BiDownvote className="text-lg" />
          </button>
        </div>
      </div>

      {/* Text and Progress Section */}
      <div className="px-6 py-4 text-left">
        <h1 className="text-xl font-bold text-gray-800">Helping Hands Foundation</h1>
        <p className="text-md text-gray-500 mb-3">by Campaign Name</p>

        {/* Hover Logic: hide details, show donate button */}
        {!hover ? (
          <div>
            {/* Donations & Time Left */}
            <div className="flex justify-between text-sm text-gray-600 mb-2">
              <p>No. of Donations</p>
              <p>Time left</p>
            </div>

            {/* Progress Bar */}
            <div className="w-full bg-gray-200 h-2 rounded-full overflow-hidden mb-2">
              <div
                className="h-full bg-green-500 transition-all duration-500"
                style={{ width: `${progress}%` }}
              ></div>
            </div>

            {/* Raised vs Goal */}
            <div className="flex justify-between text-xs text-gray-600">
              <span>₹{raised.toLocaleString()} raised</span>
              <span>Goal: ₹{goal.toLocaleString()}</span>
            </div>
          </div>
        ) : (
          <div className="flex justify-center items-center h-12">
            <button className="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-6 rounded-full transition">
              Donate
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default Card;
