import React, { useEffect, useState } from 'react'
import './header.scss'
import { Link } from "react-router-dom"

function Header( {signState, setSignState} ) {
    const [userData, setUserData] = useState({});

    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://' + document.location.hostname + ':9000/user', {
                    headers: { 'Content-Type': 'application/json' },
                    credentials: 'include',
                });


                const content = await response.json();
                if (response.ok && !("message" in content) && content["ID"] !== 0) {
                    setUserData(content)
                    setSignState(true)
                    console.log(content)
                }

            }
        )();
    }, [signState]);

    if (signState) {
        return (
            <div className='header'>
                <a href="/" className="logo">Realtime Chat</a>
                <nav className="navbar">
                    <Link to="/" className="bar">Home</Link>
                    <Link to="/sign-in" className="bar">{userData["username"]}</Link>
                    <Link to="/sign-out" className="bar">Sign out</Link>
                </nav>
            </div>
        )

    }

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