import './chatHistory.scss'
import Message from '../../Message/message'
import React, { useEffect } from 'react';
import { scrollToBottom } from '../../utils'

function ChatHistory({ history }) {
    console.log(history);
    const messages = history.map(message => <Message key={message.timeStamp} messageJSON={message.data} />);
    useEffect(() => {
        scrollToBottom();
    });
    return (
        <div className='chatHistory'>
            {messages}
        </div>
    );
}

export default ChatHistory;