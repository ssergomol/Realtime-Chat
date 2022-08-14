import React from 'react';
import './message.scss';

function Message(props) {
    let message = JSON.parse(props.messageJSON);
    return (
        <div className='message-box'>
            <p className='message-text'>
                {message.body}
            </p>
        </div>
    )
}

export default Message;