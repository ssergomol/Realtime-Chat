import React from 'react';
import './App.css';
import Header from './components/Header/header.jsx';
import Body from './components/Body/body';


function App( {signState, setSignState} ) {
  return (
    <div className="app">
      <Header signState={signState} setSignState={setSignState}/>
      <Body/>
    </div>
  );
}

export default App;
