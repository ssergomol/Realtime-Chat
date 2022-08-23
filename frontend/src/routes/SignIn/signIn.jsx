import Header from '../../components/Header/header'
import '../../components/Body/body.scss'
import './signIn.scss'
import '../../App.css'
import './signIn.scss'

export default function SignIn() {
    return (
        <div className='app'>
            <Header/>
            <div className='background sign-in'>
                <form className='sign-in-form' method='POST' action='/sign-in'>
                    <label for='username'>Username: </label>
                    <input id='username' name='username' type='text'/>
                    <label for='password'>Password: </label>
                    <input id='password' name='password' type='password'/>
                    <button id='send-form' type='submit'>Send</button>
                </form>
            </div>
        </div>
    );
}
  