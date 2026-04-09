import React, { useEffect, useState } from "react";
import heroimg from "../../public/domains/domain_example.png";
import { FaRegHeart } from "react-icons/fa";
import { BiUpvote, BiDownvote, BiGroup } from "react-icons/bi";
import { IoMdTime } from "react-icons/io";
import { Link, useNavigate } from "react-router-dom";
import { apiRequest, API_ENDPOINTS } from "../config/api";
import { getCauseImage } from "../utils/imageHelper";
import { formatGoal, formatCollected, getCollectedLabel } from "../utils/goalHelper";

const Card = ({ cause }) => {
  const navigate = useNavigate();
  const aidTypeName = cause?.aid_type?.name || "";
  const raised = parseFloat(cause.collected_amount);

  let goal

  if (cause.goal_amount) {
    goal = parseFloat(cause.goal_amount)
  } else {
    goal = 2 * raised
  }

  const progress = Math.min((raised / goal) * 100, 100);

  const [hover, setHover] = useState(false);
  const [votes, setVotes] = useState({ upvotes: 0, downvotes: 0, my_vote: null });
  const [votesLoading, setVotesLoading] = useState(false);
  const [reviewCount, setReviewCount] = useState(0);
  const [reviewCountLoading, setReviewCountLoading] = useState(false);

  const getDaysLeft = (targetTimestamp) => {
    const now = Date.now();
    const targetTime = new Date(targetTimestamp).getTime();

    const differenceInMs = targetTime - now;

    // Convert milliseconds to days
    const millisecondsPerDay = 1000 * 60 * 60 * 24;
    const daysLeft = Math.ceil(differenceInMs / millisecondsPerDay);

    return daysLeft;
  }
  
  const imageSrc = getCauseImage(cause.cover_image_url, heroimg);

  const handleOrganizationClick = (e) => {
    // Prevent outer card navigation to `/campaign/:id`
    e.preventDefault();
    e.stopPropagation();

    const orgId = cause?.organization?.id;
    if (!orgId) return;

    navigate(`/organization/${orgId}/accounts`);
  };

  useEffect(() => {
    const fetchVotes = async () => {
      if (!cause?.id) return;
      setVotesLoading(true);
      const res = await apiRequest(API_ENDPOINTS.GET_CAUSE_VOTES(cause.id));
      if (res.success && res.data) {
        setVotes(res.data);
      }
      setVotesLoading(false);
    };

    fetchVotes();
  }, [cause?.id]);

  useEffect(() => {
    const fetchReviewCount = async () => {
      if (!cause?.id) return;
      setReviewCountLoading(true);
      const res = await apiRequest(
        API_ENDPOINTS.GET_CAUSE_REVIEW_COUNT(cause.id)
      );
      if (res.success && typeof res.data?.count === "number") {
        setReviewCount(res.data.count);
      }
      setReviewCountLoading(false);
    };

    fetchReviewCount();
  }, [cause?.id]);

  const handleUpvote = async (e) => {
    e.preventDefault();
    e.stopPropagation();
    if (!cause?.id) return;

    const res = await apiRequest(API_ENDPOINTS.UPVOTE_CAUSE(cause.id), {
      method: "POST",
    });
    if (res.success && res.data) setVotes(res.data);
  };

  const handleDownvote = async (e) => {
    e.preventDefault();
    e.stopPropagation();
    if (!cause?.id) return;

    const res = await apiRequest(API_ENDPOINTS.DOWNVOTE_CAUSE(cause.id), {
      method: "POST",
    });
    if (res.success && res.data) setVotes(res.data);
  };

  const myVote = votes?.my_vote;
  const isUpActive = myVote === "up";
  const isDownActive = myVote === "down";

  return (
    <div
      className="bg-white rounded-lg shadow-md overflow-hidden transition-transform hover:scale-[1.02] cursor-pointer group"
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
    >
      <Link to={`/campaign/${cause.id}`} onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
        {/* Image Section */}
        <div className="relative">
          <img
            src={imageSrc}
            alt={cause.title}
            className="w-full h-56 object-cover"
          />

          {/* Heart Button */}
          <button className="absolute top-3 right-3 bg-white p-2 rounded-full shadow hover:bg-gray-100 transition cursor-pointer">
            <FaRegHeart className="text-red-500 text-lg" />
          </button>

          {/* Upvote/Downvote */}
          <div className="absolute bottom-3 right-3 flex items-center bg-white/90 px-2 py-1 rounded-full text-gray-700 text-sm shadow">
            <button
              type="button"
              className={isUpActive ? "text-green-600 cursor-pointer" : "hover:text-green-600 cursor-pointer"}
              onClick={handleUpvote}
              disabled={votesLoading}
              aria-label="Upvote cause"
            >
              <BiUpvote className="text-lg" />
            </button>
            <span className="mx-1">
              {votes?.upvotes} | {votes?.downvotes}
            </span>
            <button
              type="button"
              className={isDownActive ? "text-red-600 cursor-pointer" : "hover:text-red-600 cursor-pointer"}
              onClick={handleDownvote}
              disabled={votesLoading}
              aria-label="Downvote cause"
            >
              <BiDownvote className="text-lg" />
            </button>
          </div>
        </div>

        {/* Text and Progress Section */}
        <div className="px-6 py-4 text-left">
          <h1 className="text-xl font-bold text-gray-800">{cause.title}</h1>
          <p className="text-sm text-gray-500 mb-3">{cause.domain.name} | {cause.aid_type.name}</p>
          <p className="text-md text-gray-500 mb-3">
            By{" "}
            <span
              role="link"
              tabIndex={0}
              className="font-semibold text-gray-800 hover:text-[#ff6200] underline cursor-pointer"
              onClick={handleOrganizationClick}
              onKeyDown={(e) => {
                if (e.key === "Enter" || e.key === " ") handleOrganizationClick(e);
              }}
            >
              {cause.organization.name}
            </span>
          </p>
          {/*<p className="text-md text-gray-500 mb-3">by Campaign Name</p>*/}

          {/* Donations & Time Left */}
          <div className="flex justify-between text-sm text-gray-600 mb-2">
            <p className="flex items-center gap-2">
              <BiGroup className="h-5 text-lg" />
              1000+ Donations
              <span className="text-gray-300">•</span>
              <span>
                {reviewCountLoading ? "..." : `${reviewCount} Reviews`}
              </span>
            </p>
            <p className="flex"><IoMdTime className="h-5 text-lg mr-1" />{getDaysLeft(cause.deadline) + 50} Days Left</p>
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
                <span> <span className="text-lg font-bold text-black">{formatCollected(raised, aidTypeName)}</span> {getCollectedLabel(aidTypeName)}</span>
                <span>Goal: <span className="text-red-500 font-bold">{formatGoal(goal, aidTypeName)}</span></span>
              </div>
            </div>
          ) : (
            <div className="flex justify-center items-center h-12">
              <Link to="/checkout" className="w-full" onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
                <button className="bg-green-600 hover:bg-green-700 w-full text-white font-semibold py-2 px-6 mx-2 rounded-full transition cursor-pointer" onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}>
                  Donate
                </button>
              </Link>
            </div>
          )}
        </div>
      </Link>
    </div >
  );
};

export default Card;
