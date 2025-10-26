import React from "react";
import { Link } from "react-router-dom";
import { FaFacebookF, FaInstagram, FaLinkedinIn, FaTwitter } from "react-icons/fa";

const Footer = () => {
  return (
    <footer className="bg-[#3a0b2e] text-gray-300 pt-12 pb-6">
      <div className="max-w-7xl mx-auto px-6 md:px-12 grid grid-cols-1 md:grid-cols-4 gap-10 border-b border-gray-600 pb-10">

        {/* Brand Description */}
        <div>
          <h2 className="text-2xl font-bold text-white mb-4">
            Charity<span className="text-[#ff6200]">Light</span>
          </h2>
          <p className="text-sm leading-relaxed mb-4">
            Empowering change through transparency. CharityLight connects donors
            and NGOs via a secure, blockchain-powered platform — ensuring every
            contribution creates measurable impact.
          </p>

          {/* Social Icons */}
          <div className="flex space-x-4 mt-4">
            <a href="#" className="hover:text-[#ff6200] transition">
              <FaFacebookF />
            </a>
            <a href="#" className="hover:text-[#ff6200] transition">
              <FaTwitter />
            </a>
            <a href="#" className="hover:text-[#ff6200] transition">
              <FaInstagram />
            </a>
            <a href="#" className="hover:text-[#ff6200] transition">
              <FaLinkedinIn />
            </a>
          </div>
        </div>

        {/* Quick Links */}
        <div>
          <h3 className="text-xl font-semibold text-white mb-4">Quick Links</h3>
          <ul className="space-y-2">
            <li><Link to="/" className="hover:text-[#ff6200] transition">Home</Link></li>
            <li><Link to="/about" className="hover:text-[#ff6200] transition">About</Link></li>
            <li><Link to="/campaigns" className="hover:text-[#ff6200] transition">Campaigns</Link></li>
            <li><Link to="/blogs" className="hover:text-[#ff6200] transition">Blogs</Link></li>
            <li><Link to="/contact" className="hover:text-[#ff6200] transition">Contact</Link></li>
          </ul>
        </div>

        {/* Support */}
        <div>
          <h3 className="text-xl font-semibold text-white mb-4">Get Involved</h3>
          <ul className="space-y-2">
            <li><Link to="/makeContribution" className="hover:text-[#ff6200] transition">Make a Donation</Link></li>
            <li><Link to="/volunteer" className="hover:text-[#ff6200] transition">Become a Volunteer</Link></li>
            <li><Link to="/register-ngo" className="hover:text-[#ff6200] transition">Register Your NGO</Link></li>
            <li><Link to="/faq" className="hover:text-[#ff6200] transition">FAQs</Link></li>
          </ul>
        </div>

        {/* Contact Info */}
        <div>
          <h3 className="text-xl font-semibold text-white mb-4">Contact Us</h3>
          <p className="text-sm mb-2">📍 123 Hope Avenue, Mumbai, India</p>
          <p className="text-sm mb-2">📞 +91 98765 43210</p>
          <p className="text-sm mb-4">📧 contact@charitylight.org</p>
          <Link
            to="/contact"
            className="bg-[#ff6200] hover:bg-[#e45a00] text-white text-sm font-semibold px-4 py-2 rounded-md inline-block transition"
          >
            Get in Touch
          </Link>
        </div>
      </div>

      {/* Bottom Bar */}
      <div className="text-center text-gray-400 text-sm mt-6">
        © {new Date().getFullYear()} <span className="text-white font-semibold">CharityLight</span>. All Rights Reserved.
      </div>
    </footer>
  );
};

export default Footer;
