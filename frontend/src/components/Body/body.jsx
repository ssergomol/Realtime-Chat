import React, { useEffect, useState, useCallback } from 'react';
import './body.scss'
import ChatHistory from './ChatHistory/chatHistory'
import InputSection from './InputSection/inputSection'
import { connect, sendMessage } from '../../apis/websocket/index'

function Body() {

    const [chatHistory, updateChatHistory] = useState([])

    const processSocketMessage = useCallback((message) => {
        console.log("A new message has arrived!");
        updateChatHistory(prevHistory => [...prevHistory, message]);
        console.log(message);
    }, []);

    useEffect(() => {
        connect(processSocketMessage);
    }, [processSocketMessage]);

    const sendForm = (event) => {

        if (event.keyCode === 13 && !event.shiftKey) {
            sendMessage(event.target.value);
            event.target.value = "";
        }
    }

    return (
        <div className='body'>
            <ChatHistory history={chatHistory} />
            <InputSection send={sendForm} />
        </div>
    )
}

export default Body;