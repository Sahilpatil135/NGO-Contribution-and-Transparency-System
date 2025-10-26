import React from "react";
import { useAuth } from "../contexts/AuthContext";
import { NavLink, useNavigate } from "react-router-dom";
import heroimg from "/default_user_avatar.jpg";

const Navbar = () => {
    const { user, logout } = useAuth();
    const navigate = useNavigate();

    const handleLogin = () => navigate("/login");
    const handleLogout = () => logout();

    const toTitleCase = (s = "") =>
        s.charAt(0).toUpperCase() + s.slice(1).toLowerCase();

    const navLinks = [
        { name: "Home", path: "/" },
        { name: "About Us", path: "/about" },
        { name: "Campaigns", path: "/makeContribution" },
        { name: "Blog", path: "/blog" },
        { name: "Contact", path: "/contact" },
    ];

    return (
        <header className="w-full bg-white shadow-sm"
        >
            <nav className="w-11/12 mx-auto h-20 flex items-center justify-between px-4 md:px-8">
                {/* Left: Logo */}
                <div className="flex items-center gap-10">
                    {/* <div className="flex items-center gap-2 text-2xl font-semibold text-[#3a0b2e]">
            <span className="text-[#f75c03] text-xl">●</span>
            CharityLight
          </div> */}
                    <img src="/logo1.png" alt="logo" className="w-50 h-15" />

                    {/* Nav Links */}
                    <ul className="hidden md:flex items-center gap-8">
                        {navLinks.map((link) => (
                            <NavLink
                                key={link.name}
                                to={link.path}
                                className={({ isActive }) =>
                                    `text-base transition ${isActive
                                        ? "text-[#f75c03] font-semibold"
                                        : "text-gray-800 hover:text-[#f75c03]"
                                    }`
                                }
                            >
                                {link.name}
                            </NavLink>
                        ))}
                    </ul>
                </div>

                {/* Right: User Info & Auth Buttons */}
                <div className="flex items-center gap-5">
                    {user ? (
                        <div className="flex items-center gap-5">
                            <div className="text-right leading-tight hidden sm:block">
                                <p className="text-sm text-gray-800 font-medium">
                                    Welcome, {toTitleCase(user.name)}
                                </p>
                                <p className="text-xs text-gray-500">
                                    Role: {toTitleCase(user.role)}
                                </p>
                            </div>
                            <img
                                src={user.avatar_url || heroimg}
                                alt="User Avatar"
                                className="w-10 h-10 rounded-full border border-gray-300 object-cover"
                            />
                            <button
                                onClick={handleLogout}
                                className="bg-[#ff6200] text-white text-sm px-4 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer"
                            >
                                Logout
                            </button>
                        </div>
                    ) : (
                        <button
                            onClick={handleLogin}
                            className="bg-[#ff6200] text-white text-sm px-5 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer"
                        >
                            Login
                        </button>
                    )}
                </div>
            </nav>
        </header>
    );
};

export default Navbar;
