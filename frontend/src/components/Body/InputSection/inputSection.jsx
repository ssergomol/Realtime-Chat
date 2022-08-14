import React, { useEffect } from 'react'
import './inputSection.scss'
import { inputAutoGrow } from '../../utils'

function InputSection({ send }) {
    
    useEffect(inputAutoGrow, []);
    return (
        <div className="input-box">
            <textarea className="textarea resize" onKeyUp={send} autoFocus/>
        </div>

    );
}

export default InputSection;