import React, { useEffect, useState, useCallback } from 'react';
import './body.scss'
import ChatHistory from './ChatHistory/chatHistory'
import InputSection from './InputSection/inputSection'
import { defineWsListeners, sendMessage } from '../../apis/websocket/index'
import { setDefaultHeight } from '../utils'

function Body() {

    const [chatHistory, updateChatHistory] = useState([])

    const processSocketMessage = useCallback((message) => {
        console.log("A new message has arrived!");
        updateChatHistory(prevHistory => [...prevHistory, message]);
        console.log(message);
    }, []);

    useEffect(() => {
        defineWsListeners(processSocketMessage);
    }, [processSocketMessage]);

    const sendForm = (event) => {
        if (event.keyCode === 13 && !event.shiftKey) {
            let text = event.target.value.trim();
            if (text.length !== 0) {
                sendMessage(text);
            }
            event.target.value = "";
            setDefaultHeight();
        }
    }

    return (
        <div className='background body'>
            <ChatHistory history={chatHistory} />
            <InputSection send={sendForm} />
        </div>
    )
}

export default Body;