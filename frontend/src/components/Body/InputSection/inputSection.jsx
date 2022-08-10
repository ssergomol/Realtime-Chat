import React, { useEffect } from 'react'
import './inputSection.scss'
import inputAutoGrow from './utils'
function InputSection() {
    
    useEffect(inputAutoGrow, []);
    return (
        <div className="input-box">
            <textarea class="textarea resize"></textarea>
        </div>

    );
}

export default InputSection;