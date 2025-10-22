import React from 'react'
import { useAuth } from '../contexts/AuthContext';

const Navbar = () => {
    const { logout } = useAuth();

    const handleLogout = () => {
        logout();
    };

    return (
        <div className="w-full">
            <nav className="w-11/12 h-20 flex items-center justify-between px-8 bg-white mx-auto shadow-sm">
                <div className="flex items-center gap-10">
                    <div className="text-2xl font-semibold text-[#3a0b2e] flex items-center gap-2">
                        <span className="text-[#f75c03] text-xl">●</span> CharityLight
                    </div>

                    <ul className="hidden md:flex list-none gap-8">
                        <li className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Home</li>
                        <li className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">About Us</li>
                        <li className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Campaigns</li>
                        <li className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Blog</li>
                        <li className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Contact</li>
                    </ul>
                </div>

                <div className="flex items-center gap-4">
                    <button className="bg-[#ff6200] text-white text-sm px-5 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer">
                        Login
                    </button>
                    <button className="bg-[#ff6200] text-white text-sm px-5 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer" onClick={handleLogout}>
                        Logout
                    </button>
                </div>
            </nav>
        </div>
    )
}

export default Navbar