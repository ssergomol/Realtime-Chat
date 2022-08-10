import React from 'react'
import './header.scss'

function Header() {
    return (
        <div className='header'>
            <a href="/" class="logo">Realtime Chat</a>
            <nav className="navbar">
                <a className="bar" href="#home">Home</a>
                <a className="bar" href="#contact">Contact</a>
                <a className="bar" href="#about">About</a>
            </nav>
        </div>
    )
}

export default Header;