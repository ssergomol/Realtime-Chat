import React, { useEffect, useState } from 'react';
import './body.scss'
import ChatHistory from './ChatHistory/chatHistory'
import InputSection from './InputSection/inputSection'
import { connect, sendMessage } from '../../apis/websocket/index'

function Body() {

    const [chatHistory, updateChatHistory] = useState([])

    const processSocketMessage = function (message) {
        console.log("A new message has arrived!");
        updateChatHistory(previousHistory => [...previousHistory, message]);
        console.log(message);
    }

    useEffect(() => { connect(processSocketMessage); }, []);
    
    const sendForm = (event) => {
        if (event.key === 13 && !event.shiftKey) {
            sendMessage(event.target.value);
            event.target.value = "";
        }
    }

    return (
        <div className='body'>
            <ChatHistory messages={chatHistory} />
            <InputSection send={sendForm} />
        </div>
    )
}

export default Body;