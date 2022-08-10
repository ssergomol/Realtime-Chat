import React from 'react';
import './body.scss'
import ChatHistory from './ChatHistory/chatHistory'
import InputSection from './InputSection/inputSection'

function Body(){
    return(
        <div className='body'>
            <ChatHistory/>
            <InputSection/>
        </div>
    )
}

export default Body;