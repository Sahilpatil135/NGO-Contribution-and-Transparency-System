import React from 'react'
import { useAuth } from '../contexts/AuthContext';
import { NavLink, useNavigate } from 'react-router-dom';
import heroimg from "/default_user_avatar.jpg"

const Navbar = () => {
    const { user, logout } = useAuth();

    const handleLogout = () => {
        logout();
    };

    let nav = useNavigate()

    const handleLogin = () => {
        nav('/login')
    };

    const toTitleCase = s => { return s[0].toUpperCase() + s.substring(1).toLowerCase() }

    return (
        <div className="w-full">
            <nav className="w-11/12 h-20 flex items-center justify-between px-8 bg-white mx-auto shadow-sm">
                <div className="flex items-center gap-10">
                    <div className="text-2xl font-semibold text-[#3a0b2e] flex items-center gap-2">
                        <span className="text-[#f75c03] text-xl">●</span> CharityLight
                    </div>


                    <ul className="hidden md:flex list-none gap-8">
                        <NavLink to="/" className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Home</NavLink>
                        <NavLink to="" className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">About Us</NavLink>
                        <NavLink to="/makeContribution" className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Campaigns</NavLink>
                        <NavLink to="" className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Blog</NavLink>
                        <NavLink to="" className="text-base cursor-pointer text-black hover:text-[#f75c03] transition">Contact</NavLink>
                    </ul>
                </div>

                {user && (
                    <div>
                        <span>Welcome, {toTitleCase(user.name)}</span><br />
                        <span>Your Role is {toTitleCase(user.role)}</span>
                        {user.avatar_url ?
                            (<img className="rounded-full w-9 h-9" src={user.avatar_url} />) :
                            (<img className="rounded-full w-9 h-9" src={heroimg} />)
                        }

                    </div>
                )}

                <div className="flex items-center gap-4">
                    <button className="bg-[#ff6200] text-white text-sm px-5 py-2 rounded-md transition hover:bg-[#e65500] cursor-pointer" onClick={handleLogin}>
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
