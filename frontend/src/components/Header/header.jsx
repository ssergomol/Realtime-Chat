import React from 'react'
import './header.scss'
import { Link } from "react-router-dom"

function Header() {
    return (
        <div className='header'>
            <a href="/" className="logo">Realtime Chat</a>
            <nav className="navbar">
                <Link to="/" className="bar">Home</Link>
                <Link to="/sign-in" className="bar">Sign in</Link>
                <Link to="/sign-up" className="bar">Sign up</Link>
            </nav>
        </div>
    )
}

export default Header;