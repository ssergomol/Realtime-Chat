import React, { useEffect } from 'react';
import  { useNavigate } from 'react-router-dom'


export default function SignOut( {changeSignState}) {
    let navigate = useNavigate();

    useEffect(() => {
        (
            async () => {
                
                const response = await fetch('http://' + document.location.hostname + ':9000/sign-out', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    credentials: 'include',
                });


                const content = await response.json();
                if (response.ok) {
                    changeSignState(false);
                    console.log(content)
                    navigate("/sign-in")
                } else {
                    navigate("/")
                }
            }
        )();
    }, [])

    return (
        <></>
    );
}
  