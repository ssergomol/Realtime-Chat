import Header from '../../components/Header/header';
import './signUp.scss';
import React, { useState } from 'react';
import  { useNavigate } from 'react-router-dom'

export default function SignUp( {signState, setSignState} ) {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [secondPassword, setSecondPassword] = useState("")
    let navigate = useNavigate();

    const submitHandler = async (event) => {
        event.preventDefault();
        if (password !== secondPassword) {
            console.log("Passwords are different")
        }
        const response = await fetch('http://' + document.location.hostname + ':9000/sign-up', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                "username": username,
                "password": password,
                "email": email,
            })
        });

        if (response.ok) {
            const content = await response.json();
            console.log(content); 
            navigate("/sign-in");
        }
          
    }


    return (
        <div className='app'>
            <Header signState={signState} setSignState={setSignState}/>
            <div className='background sign-up'>
                <form className='sign-up-form' method='POST' action='/sign-up' onSubmit={submitHandler}>
                    <label htmffor='username'>Username: </label>
                    <input id='username' name='username' type='text' 
                    onChange={event => setUsername(event.target.value)}/>

                    <label htmlFor='password'>Password: </label>
                    <input id='password' name='password' type='password' 
                    onChange={event => setPassword(event.target.value)}/>

                    <label htmlFor='confirm-password'>Password: </label>
                    <input id='confirm-password' name='confirm-password' type='password' 
                    onChange={event => setSecondPassword(event.target.value)}/>

                    <label htmlFor='email'>Email: </label>
                    <input id='email' name='email' type='email'
                    onChange={event => setEmail(event.target.value)}/>
                    <input id='send-form' type='submit' value="Submit"/>
                </form>
            </div>
        </div>
    );
}
