import Header from '../../components/Header/header'
import './signUp.scss'

export default function SignUp() {
    return (
        <div className='app'>
            <Header />
            <div className='background sign-up'>
                <form className='sign-up-form' method='POST' action='/sign-up'>
                    <label for='username'>Username: </label>
                    <input id='username' name='username' type='text'/>
                    <label for='password'>Password: </label>
                    <input id='password' name='password' type='password'/>
                    <label for='confirm-password'>Password: </label>
                    <input id='confirm-password' name='confirm-password' type='password'/>
                    <label for='email'>Email: </label>
                    <input id='email' name='email' type='email'/>
                    <button id='send-form' type='submit'>Send</button>
                </form>
            </div>
        </div>
    );
}
