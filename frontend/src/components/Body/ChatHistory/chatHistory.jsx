import './chatHistory.scss'
import Message from '../../Message/message'

function ChatHistory({ history }) {
    console.log(history);
    const messages = history.map(message => <Message key={message.timeStamp} messageJSON={message.data} />);
    return (
        <div className='chatHistory'>
            {messages}
        </div>
    );
}

export default ChatHistory;