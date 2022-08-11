import './chatHistory.scss'
// import Message from '../../Message/message'

function ChatHistory({ chatHistory }) {
    console.log(chatHistory);
    //const messages = chatHistory.map(message => <Message key={message.timeStamp} messageJSON={message.data} />);
    return (
        <div className='chatHistory'>
            {/* {messages} */}
        </div>
    );
}

export default ChatHistory;