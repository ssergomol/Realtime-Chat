import React, { useCallback, useState } from 'react';
import '../../index.css';
import App from '../../App';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import SignIn from '../../routes/SignIn/signIn'
import SignUp from '../../routes/SignUp/signUp'
import SignOut from '../../routes/SignOut/signOut'


export default function AppContainer() {
const [IsSignedIn, setSignedIn] = useState(false);
const setSignState = useCallback(val => {
  setSignedIn(val);
}, [setSignedIn]);

return (
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<App signState={IsSignedIn} setSignState ={setSignState}/>}/>
      <Route path="/sign-in" element={<SignIn/>}/>
      <Route path="/sign-up" element={<SignUp/>}/>
      <Route path="/sign-out" element={<SignOut changeSignState={setSignState}/>}/>
    </Routes>
    
  </BrowserRouter>
);
}
