import Header from '../../components/Header/header'
import '../../components/Body/body.scss'
import './signIn.scss'
import '../../App.css'
import './signIn.scss'
import React, { useState } from 'react';
import  { useNavigate } from 'react-router-dom'


export default function SignIn() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    let navigate = useNavigate();

    const submit = async (event) => {
        event.preventDefault();
        const response = await fetch('http://' + document.location.hostname + ':9000/sign-in', {
            method: 'POST',
            credentials: 'include',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                "username": username,
                "password": password,
            })
        });

        if (response.ok) {
            const content = await response.json();
            console.log(content); 
            navigate("/");
        }
    }
    return (
        <div className='app'>
            <Header/>
            <div className='background sign-in'>
                <form className='sign-in-form' method='POST' action='/sign-in' onSubmit={submit}>
                    <label htmlFor='username'>Username: </label>
                    <input id='username' name='username' type='text'
                        onChange={event => setUsername(event.target.value)}
                    />
                    <label htmlFor='password'>Password: </label>
                    <input id='password' name='password' type='password'
                        onChange={event => setPassword(event.target.value)}
                    />
                    <button id='send-form' type='submit'>Send</button>
                </form>
            </div>
        </div>
    );
}
  